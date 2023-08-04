package common

import "log"

func (c debugEnable) LogDebug(message string) {
	if c.Enable {
		log.Println(message)
	}
}

func (c debugEnable) LogDebugError(message string, err error) {
	if c.Enable {
		log.Println(message + err.Error())
	}
}

func (c debugEnable) LogFatalDebug(message string) {
	if c.Enable {
		log.Fatal(message)
	}
}

func (c debugEnable) LogFatalDebugError(message string, err error) {
	if c.Enable {
		log.Fatal(message, err)
	}
}
