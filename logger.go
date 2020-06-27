// package logger_v1

// import (
// 	"log"
// 	"os"
// )

// // WLogger exported
// type WLogger struct {
// 	info    *log.Logger
// 	error   *log.Logger
// 	warning *log.Logger
// 	fatal   *log.Logger
// }

// // exported ResetColor ...
// const (
// 	flags        = log.Ldate | log.Ltime | log.Lshortfile
// 	infoTag      = "INFO:"
// 	errorTag     = "ERROR:"
// 	warningTag   = "WARNING:"
// 	fatalTag     = "FATAL:"
// 	ResetColor   = "\033[0m"  // Reset the default color
// 	InfoColor    = "\033[34m" //Blue
// 	FatalColor   = "\033[35m" //Purple
// 	WarningColor = "\033[33m" // Yellow
// 	ErrorColor   = "\033[31m" // Red color obivously
// 	DebugColor   = "\033[32m" // Green
// )

// func getLogger() *WLogger {

// 	logger := &WLogger{

// 		error:   log.New(os.Stderr, ErrorColor+errorTag+ResetColor, flags),
// 		info:    log.New(os.Stderr, InfoColor+infoTag+ResetColor, flags),
// 		warning: log.New(os.Stderr, WarningColor+warningTag+ResetColor, flags),
// 		fatal:   log.New(os.Stderr, FatalColor+fatalTag+ResetColor, flags),
// 	}
// 	return logger
// }
