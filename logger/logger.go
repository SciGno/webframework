package logger

import (
	"fmt"
	"log"
	"os"
)

var info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds)
var warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lmicroseconds)
var debug = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
var errors = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds)
var fatal = log.New(os.Stdout, "ERROR: ", log.Lshortfile|log.LstdFlags|log.Ldate|log.Ltime|log.Lmicroseconds)

// Info function
func Info(format string, v ...interface{}) {
	info.Printf(fmt.Sprintf(format, v...))
}

// Warning function
func Warning(format string, v ...interface{}) {
	warning.Printf(fmt.Sprintf(format, v...))
}

// Debug function
func Debug(format string, v ...interface{}) {
	debug.Printf(fmt.Sprintf(format, v...))
}

// Error function
func Error(format string, v ...interface{}) {
	errors.Printf(fmt.Sprintf(format, v...))
}

// Fatal function
func Fatal(format string, v ...interface{}) {
	fatal.Printf(fmt.Sprintf(format, v...))
}
