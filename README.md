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


## Donation

If you like what I've done and would like to donate, click the link below

[PayPal Donation](<https://www.paypal.com/donate/?hosted_button_id=VN9G973VD8HRA>)
