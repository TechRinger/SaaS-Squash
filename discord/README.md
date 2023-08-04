# Discord C2 Setup

## Requirements

- Discord account
- Discord Sever

## Set up

1. go to the Discord Developer Portal <https://discord.com/developers/applications>
1. Click on New Application
![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1691010901276_0.png)
1. Fill out the Name and Team field, select the check box and click Create
![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1691010990594_0.png)
1. Click on Settings > Bot then configure the following then select Save Changes:
    1. (Deselect) PUBLIC BOT
    1. (Select) MESSAGE CONTENT INTENT
1. Click on Settings > Bot > Reset Token
![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1691013071090_0.png)
1. Click Yes, do it! on the pop up
[image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1691013100282_0.png)
1. Copy the TOKEN, this will be used as the Token for the C2 app
![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1691013152089_0.png)
1. Click on Settings > OAuth2 > URL Generator and select permissions.  When you are done copy the URL at the bottom.  This is what you'll use to add the bot to your Discord Server.  (Example: <https://discord.com/api/oauth2/authorize?client_id=1136407212537942036&permissions=56456500&scope=bot>)

   1. Read Messages/View Channels
   2. Send Messages
   3. Create Public Threads
   4. Create Private Threads
   5. Send Messages in Threads
   6. Manage Messages
   7. Manage Threads
   8. Attach Files
   9. Read Message History
[image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1691011651650_0.png)

1. Navigate to the URL you generated in the previous step and select the server you want to add it to and then click Continue
![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1691012283021_0.png)
1. Verify the options are what you selected for the Bot and click Authorize
![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1691012328460_0.png)
1. Log into your Discord Server and select the channel you want to add the Bot to and click on Settings
![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1691012495786_0.png)
1. Go to Permissions and click on Add members or roles > Select your Bot by role or by it's name and click Done
![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1691012634207_0.png)

1. If you've added your Bot correctly you should see it subdued under the member list in the channel.
![image.png](https://raw.githubusercontent.com/TechRinger/media/main/images/image_1691012747559_0.png)
