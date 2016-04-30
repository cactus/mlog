// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

// A Logger represents a logging object, that embeds log.Logger, and
// provides support for a toggle-able debug flag.
type Logger struct {
	mu      sync.Mutex // ensures atomic writes are synchronized
	out     io.Writer
	flags   uint64
	fwriter LogFormatWriter
}

func (l *Logger) Write(b []byte) (int, error) {
	// lock writing to serialize log output (no scrambled log lines)
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.out.Write(b)
}

// Emit invokes the LogFormatWriter and logs the event.
func (l *Logger) Emit(level int, message string, extra Map) {
	l.fwriter.Emit(l, level, message, extra)
}

// SetFormatter sets the LogFormatWriter
func (l *Logger) SetFormatter(w LogFormatWriter) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.fwriter = w
}

// Flags retuns the current FlagSet
func (l *Logger) Flags() FlagSet {
	return FlagSet(atomic.LoadUint64(&l.flags))
}

// SetFlags sets the current FlagSet
func (l *Logger) SetFlags(flags FlagSet) {
	atomic.StoreUint64(&l.flags, uint64(flags))
}

// HasDebug returns true if the debug logging FlagSet is enabled, false
// otherwise.
func (l *Logger) HasDebug() bool {
	flags := FlagSet(atomic.LoadUint64(&l.flags))
	return flags&Ldebug != 0
}

// Debugm conditionally logs message and any Map elements at level="debug".
// If the Logger does not have the Ldebug flag, nothing is logged.
func (l *Logger) Debugm(message string, v Map) {
	if l.HasDebug() {
		l.Emit(-1, message, v)
	}
}

// Infom logs message and any Map elements at level="info".
func (l *Logger) Infom(message string, v Map) {
	l.Emit(0, message, v)
}

// Printm logs message and any Map elements at level="info".
func (l *Logger) Printm(message string, v Map) {
	l.Emit(0, message, v)
}

// Fatalm logs message and any Map elements at level="fatal", then calls
// os.Exit(1)
func (l *Logger) Fatalm(message string, v Map) {
	l.Emit(1, message, v)
	os.Exit(1)
}

// Debugf formats and conditionally logs message at level="debug".
// If the Logger does not have the Ldebug flag, nothing is logged.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.HasDebug() {
		l.Emit(-1, fmt.Sprintf(format, v...), nil)
	}
}

// Infof formats and logs message at level="info".
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Emit(0, fmt.Sprintf(format, v...), nil)
}

// Printf formats and logs message at level="info".
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Emit(0, fmt.Sprintf(format, v...), nil)
}

// Fatalf formats and logs message at level="fatal", then calls
// os.Exit(1)
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Emit(1, fmt.Sprintf(format, v...), nil)
	os.Exit(1)
}

// Debug conditionally logs message at level="debug".
// If the Logger does not have the Ldebug flag, nothing is logged.
func (l *Logger) Debug(v ...interface{}) {
	if l.HasDebug() {
		l.Emit(-1, fmt.Sprint(v...), nil)
	}
}

// Info logs message at level="info".
func (l *Logger) Info(v ...interface{}) {
	l.Emit(0, fmt.Sprint(v...), nil)
}

// Print logs message at level="info".
func (l *Logger) Print(v ...interface{}) {
	l.Emit(0, fmt.Sprint(v...), nil)
}

// Fatal logs message at level="fatal", then calls
// os.Exit(1)
func (l *Logger) Fatal(v ...interface{}) {
	l.Emit(1, fmt.Sprint(v...), nil)
	os.Exit(1)
}

// New creates a new Logger.
// The debug argument specifies whether debug should be set or not.
func New(out io.Writer, flags FlagSet) *Logger {
	return &Logger{
		out:     out,
		flags:   uint64(flags),
		fwriter: &PlainLogWriter{},
	}
}
