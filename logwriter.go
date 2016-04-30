package mlog

// LogFormatWriter is the interface implemented by mlog logging format writers.
type LogFormatWriter interface {
	Emit(logger *Logger, level int, message string, extra Map)
}
