package mlog

type LogFormatWriter interface {
	Emit(logger *Logger, level int, message string, extra Map)
}
