package mlog

import (
	"fmt"
	"os"
)

// the default package level Logger
var DefaultLogger = New(os.Stderr, Lstd)

func Flags() FlagSet {
	return DefaultLogger.Flags()
}

func SetFlags(flags FlagSet) {
	DefaultLogger.SetFlags(flags)
}

// Debugm logs to the default Logger. See Logger.Debugm
func Debugm(message string, v ...Map) {
	if DefaultLogger.HasDebug() {
		DefaultLogger.Output(2, "D", message, v...)
	}
}

// Infom logs to the default Logger. See Logger.Infom
func Infom(message string, v ...Map) {
	DefaultLogger.Output(2, "I", message, v...)
}

// Printm logs to the default Logger. See Logger.Printm
func Printm(message string, v ...Map) {
	DefaultLogger.Output(2, "I", message, v...)
}

// Fatalm logs to the default Logger. See Logger.Fatalm
func Fatalm(message string, v ...Map) {
	DefaultLogger.Output(2, "F", message, v...)
	os.Exit(1)
}

// Debugf logs to the default Logger. See Logger.Debugf
func Debugf(format string, v ...interface{}) {
	if DefaultLogger.HasDebug() {
		DefaultLogger.Output(2, "D", fmt.Sprintf(format, v...))
	}
}

// Infof logs to the default Logger. See Logger.Infof
func Infof(format string, v ...interface{}) {
	DefaultLogger.Output(2, "I", fmt.Sprintf(format, v...))
}

// Printf logs to the default Logger. See Logger.Printf
func Printf(format string, v ...interface{}) {
	DefaultLogger.Output(2, "I", fmt.Sprintf(format, v...))
}

// Fatalf logs to the default Logger. See Logger.Fatalf
func Fatalf(format string, v ...interface{}) {
	DefaultLogger.Output(2, "F", fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Debug logs to the default Logger. See Logger.Debug
func Debug(v ...interface{}) {
	if DefaultLogger.HasDebug() {
		DefaultLogger.Output(2, "D", fmt.Sprint(v...))
	}
}

// Info logs to the default Logger. See Logger.Info
func Info(v ...interface{}) {
	DefaultLogger.Output(2, "I", fmt.Sprint(v...))
}

// Print logs to the default Logger. See Logger.Print
func Print(v ...interface{}) {
	DefaultLogger.Output(2, "I", fmt.Sprint(v...))
}

// Fatal logs to the default Logger. See Logger.Fatal
func Fatal(v ...interface{}) {
	DefaultLogger.Output(2, "F", fmt.Sprint(v...))
	os.Exit(1)
}
