package logger

import (
	"log"
	"io"
)

//// Logger wraps logging functionality behind a boolean toggle
//
type Logger struct {
	enabledPrint bool // Enables Print/Printf/Println functionality
	enabledPanic bool // Enables Panic/Panicf/Panicln functionality
	enabledFatal bool // Enables Fatal/Fatalf/Fatalln functionality
}
//
////

//// These flags define which text to prefix to each log entry (https://pkg.go.dev/log#pkg-constants)
//
const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LUTC          = log.LUTC
	Lmsgprefix    = log.Lmsgprefix
	LstdFlags     = log.LstdFlags
)
//
////

//// Create a new Logger instance
// 
func NewLogger(enabledPrint bool, enabledPanic bool, enabledFatal bool) *Logger {
	return &Logger {
		enabledPrint: enabledPrint,
		enabledPanic: enabledPanic,
		enabledFatal: enabledFatal,
	}
}
// 
////

//// Wrapping log.Print
//
func (logger *Logger) Print(args ...interface{}) {
	if logger.enabledPrint == true {
		log.Print(args...)
	}
}	
//
////

//// Wrapping log.Printf
//
func (logger *Logger) Printf(format string, args ...interface{}) {
	if logger.enabledPrint == true {
		log.Printf(format, args...)
	}
}	
//
////

//// Wrapping log.Println
//
func (logger *Logger) Println(args ...interface{}) {
	if logger.enabledPrint == true {
		log.Println(args...)
	}
}
//
////

//// Wrapping log.Fatal
//
func (logger *Logger) Fatal(args ...interface{}) {
	if logger.enabledFatal == true {
		log.Fatal(args...)
	}
}
//
////

//// Wrapping log.Fatalf
//
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	if logger.enabledFatal == true {
		log.Fatalf(format, args...)
	}
}
//
////

//// Wrapping log.Fatalln
//
func (logger *Logger) Fatalln(args ...interface{}) {
	if logger.enabledFatal == true {
		log.Fatalln(args...)
	}
}
//
////

//// Wrapping log.Panic
//
func (logger *Logger) Panic(args ...interface{}) {
	if logger.enabledPanic == true {
		log.Panic(args...)
	}
}
//
////

//// Wrapping log.Panicf
//
func (logger *Logger) Panicf(format string, args ...interface{}) {
	if logger.enabledPanic == true {
		log.Panicf(format, args...)
	}
}
//
////

//// Wrapping log.Panicln
//
func (logger *Logger) Panicln(args ...interface{}) {
	if logger.enabledPanic == true {
		log.Panicln(args...)
	}
}
//
////

//// Wrapping log.Prefix
//
func (logger *Logger) Prefix() string {
	return log.Prefix()
}
//
////

//// Wrapping log.SetPrefix
//
func (logger *Logger) SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}
//
////

//// Wrapping log.Flags
//
func (logger *Logger) Flags() int {
	return log.Flags()
}
//
////

//// Wrapping log.SetFlags
//
func (logger *Logger) SetFlags(flag int) {
	log.SetFlags(flag)
}
//
////

//// Wrapping log.Output
//
func (logger *Logger) Output(calldepth int, s string) error {
	return log.Output(calldepth, s)
}
//
////

//// Wrapping log.Writer
//
func (logger *Logger) Writer() io.Writer {
	return log.Writer()
}
//
////

