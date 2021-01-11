package utility

import (
	log "github.com/sirupsen/logrus"
)

// Version Build Version
var Version = "0.1"

// StandardFields for logger
var StandardFields = log.Fields{
	"hostname": "ICARUS-PC",
	"appname":  "LogTracker",
}

// CheckError for checking any errors
func CheckError(message string, err error) {
	if err != nil {
		log.WithFields(StandardFields).Fatalf("%s %+v", message, err)
	}
}

// Message Return Response as map
func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// MessageList Return Response as map array
func MessageList(status int, message []string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// MessageByte Return Response as byte array
func MessageByte(status int, message []byte) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}
