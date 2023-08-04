package common

import (
	"os"
	"strings"

	"github.com/google/uuid"
)

func GetHostID() string {
	var hostID string = ""
	var fullUUID uuid.UUID
	var err error
	var hostIDBool bool = false
	AllC2Configs.Debug.LogDebug("Getting HostID")
	fullUUID, err = uuid.NewRandom()
	if err != nil {
		AllC2Configs.Debug.LogDebugError("Error generating UUID: ", err)
		return "qwer"
	}
	AllC2Configs.Debug.LogDebug("Full UUID: " + fullUUID.String())
	hostID, _, hostIDBool = strings.Cut(fullUUID.String(), "-")
	if !hostIDBool {
		AllC2Configs.Debug.LogDebug("Error Creating Host ID")
		return "asdf"
	} else {
		AllC2Configs.Debug.LogDebug("Host ID: " + hostID)
		return hostID
	}
}

func GetHostName() string {
	//check OS and get hostname for WIndows and Linux
	hostname, err := os.Hostname()
	if err != nil {
		AllC2Configs.Debug.LogDebugError("Error getting hostname: ", err)
		return "NotAHostName"
	}
	return hostname
}
