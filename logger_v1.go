package main

import (
	"log"
	"os"
)

// exported ResetColor ...
const (
	flags        = log.Ldate | log.Ltime | log.Lshortfile
	infoTag      = "INFO:"
	errorTag     = "ERROR:"
	warningTag   = "WARNING:"
	fatalTag     = "FATAL:"
	ResetColor   = "\033[0m"  // Reset the default color
	InfoColor    = "\033[34m" //Blue
	FatalColor   = "\033[35m" //Purple
	WarningColor = "\033[33m" // Yellow
	ErrorColor   = "\033[31m" // Red color obivously
	DebugColor   = "\033[32m" // Green
)

var UserIpAddress = "127.0.0.1"
var UserRsessionId = "UUID"
var UserRequestId = "UUID if not coming."

type WLogger struct {
}

func (l WLogger) info(text string) {
	log.New(os.Stderr, InfoColor+infoTag+ResetColor, flags).Printf(InfoColor + UserIpAddress + " " + UserRsessionId + " " + text + ResetColor)
}
func (l WLogger) error(text string) {
	log.New(os.Stderr, ErrorColor+errorTag+ResetColor, flags).Printf(ErrorColor + UserIpAddress + " " + UserRsessionId + " " + text + ResetColor)
}
func (l WLogger) warning(text string) {
	log.New(os.Stderr, WarningColor+warningTag+ResetColor, flags).Printf(WarningColor + UserIpAddress + " " + UserRsessionId + " " + text + ResetColor)
}
func (l WLogger) fatal(text string) {
	log.New(os.Stderr, FatalColor+fatalTag+ResetColor, flags).Printf(FatalColor + UserIpAddress + " " + UserRsessionId + " " + text + ResetColor)
}

// GetLogger exported
func GetLogger() *WLogger {

	logger := &WLogger{}
	return logger
}
