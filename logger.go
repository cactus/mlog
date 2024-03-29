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

// Emitter is the interface implemented by mlog logging format writers.
type Emitter interface {
	Emit(logger *Logger, level int, message string, extra Map)
	EmitAttrs(logger *Logger, level int, message string, extra ...*Attr)
}

// A Logger represents a logging object, that embeds log.Logger, and
// provides support for a toggle-able debug flag.
type Logger struct {
	out   io.Writer
	e     Emitter
	mu    sync.Mutex // ensures atomic writes are synchronized
	flags uint64
}

// SetOutput sets the Logger output io.Writer
func (l *Logger) SetOutput(writer io.Writer) {
	// lock writing to serialize log output (no scrambled log lines)
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = writer
}

func (l *Logger) Write(b []byte) (int, error) {
	// lock writing to serialize log output (no scrambled log lines)
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.out.Write(b)
}

// Emit invokes the FormatWriter and logs the event.
func (l *Logger) Emit(level int, message string, extra Map) {
	l.e.Emit(l, level, message, extra)
}

// Emit invokes the FormatWriter and logs the event.
func (l *Logger) EmitAttrs(level int, message string, extra ...*Attr) {
	l.e.EmitAttrs(l, level, message, extra...)
}

// SetEmitter sets the Emitter
func (l *Logger) SetEmitter(e Emitter) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.e = e
}

// Flags returns the current FlagSet
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

// Debugx conditionally logs message and any Attr elements at level="debug".
// If the Logger does not have the Ldebug flag, nothing is logged.
func (l *Logger) Debugx(message string, attrs ...*Attr) {
	if l.HasDebug() {
		l.EmitAttrs(-1, message, attrs...)
	}
}

// Infox logs message and any Map elements at level="info".
func (l *Logger) Infox(message string, attrs ...*Attr) {
	l.EmitAttrs(0, message, attrs...)
}

// Printx logs message and any Map elements at level="info".
func (l *Logger) Printx(message string, attrs ...*Attr) {
	l.EmitAttrs(0, message, attrs...)
}

// Fatalx logs message and any Map elements at level="fatal", then calls
// os.Exit(1)
func (l *Logger) Fatalx(message string, attrs ...*Attr) {
	l.EmitAttrs(1, message, attrs...)
	os.Exit(1)
}

// Panicx logs message and any Map elements at level="fatal", then calls
// panic().
func (l *Logger) Panicx(message string, attrs ...*Attr) {
	l.EmitAttrs(1, message, attrs...)
	panic(message)
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

// Panicm logs message and any Map elements at level="fatal", then calls
// panic().
func (l *Logger) Panicm(message string, v Map) {
	l.Emit(1, message, v)
	panic(message)
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

// Panicf formats and logs message at level="fatal", then calls
// panic().
func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Emit(1, s, nil)
	panic(s)
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

// Panic logs message at level="fatal", then calls
// panic().
func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Emit(1, s, nil)
	panic(s)
}

// New creates a new Logger.
func New(out io.Writer, flags FlagSet) *Logger {
	return NewFormatLogger(out, flags, &FormatWriterStructured{})
}

// NewFormatLogger creates a new Logger, using the specified Emitter.
func NewFormatLogger(out io.Writer, flags FlagSet, e Emitter) *Logger {
	return &Logger{
		out:   out,
		flags: uint64(flags),
		e:     e,
	}
}
