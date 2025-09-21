package logger

import (
	"log"
)

//// Logger wraps logging functionality behind a boolean toggle
//
type Logger struct {
	enabled bool
}
//
////

//// Create a new Logger instance
// 
func NewLogger(enabled bool) *Logger {
	return &Logger { enabled: enabled }
}
// 
////

//// Wrapping log.Print
//
func (logger *Logger) Print(args ...interface{}) {
	if logger.enabled == true {
		log.Print(args)
	}
}	
//
////

//// Wrapping log.Printf
//
func (logger *Logger) Printf(format string, args ...interface{}) {
	if logger.enabled == true {
		log.Printf(format, args)
	}
}	
//
////

//// Wrapping log.Println
//
func (logger *Logger) Println(args ...interface{}) {
	if logger.enabled == true {
		log.Println(args)
	}
}
//
////

//// Wrapping log.Fatal
//
func (logger *Logger) Fatal(args ...interface{}) {
	if logger.enabled == true {
		log.Fatal(args)
	}
}
//
////

