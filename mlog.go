// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	Lbase      uint64 = 0
	Ldatetime  uint64 = 1 << iota // log the date+time
	Ldebug                        // enable debug level log
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	Lsort                         // sort keys in output
	Lstd       = Lbase | Ldatetime
)

var (
	bufPool     = newBufferPool()
	NEWLINE     = []byte("\n")
	SPACE       = []byte{' '}
	COLON       = []byte{':'}
	QUOTE       = []byte{'"'}
	EQUAL_QUOTE = []byte{'=', '"'}
	QUOTE_SPACE = []byte{'"', ' '}
)

// A Logger represents a logging object, that embeds log.Logger, and
// provides support for a toggle-able debug flag.
type Logger struct {
	mu    sync.Mutex // ensures atomic writes are synchronized
	out   io.Writer
	flags uint64
}

func (l *Logger) Output(depth int, level string, message string, data ...*LogMap) {
	// get this as soon as possible
	now := formattedDate.String()

	buf := bufPool.Get()
	defer bufPool.Put(buf)

	flags := atomic.LoadUint64(&l.flags)
	if flags&Ldatetime != 0 {
		buf.WriteString(`time="`)
		buf.WriteString(now)
		buf.Write(QUOTE_SPACE)
	}

	buf.WriteString(`level="`)
	buf.WriteString(level)
	buf.Write(QUOTE)

	if flags&(Lshortfile|Llongfile) != 0 {
		_, file, line, ok := runtime.Caller(depth)
		if !ok {
			file = "???"
			line = 0
		}

		if flags&Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}

		buf.WriteString(` caller="`)
		buf.WriteString(file)
		buf.Write(COLON)
		buf.WriteString(strconv.Itoa(line))
		buf.Write(QUOTE)
	}

	buf.WriteString(` msg="`)
	buf.WriteString(strings.TrimSpace(message))
	buf.Write(QUOTE)

	if len(data) > 0 {
		for _, e := range data {
			buf.Write(SPACE)
			if flags&Lsort != 0 {
				e.SortedWriteTo(buf)
			} else {
				e.WriteTo(buf)
			}
		}
	}

	buf.Write(NEWLINE)

	l.mu.Lock()
	defer l.mu.Unlock()
	buf.WriteTo(l.out)
}

func (l *Logger) SetFlags(flags uint64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	atomic.StoreUint64(&l.flags, flags)
}

func (l *Logger) HasDebug() bool {
	flags := atomic.LoadUint64(&l.flags)
	return flags&Ldebug != 0
}

// Debug conditionally logs message and any LogMaps at level="debug"
// If the Logger does not have the Ldebug flag, nothing is logged.
func (l *Logger) Debug(message string, v ...*LogMap) {
	if l.HasDebug() {
		l.Output(2, "debug", message, v...)
	}
}

// Logs message and any LogMaps at level="info"
func (l *Logger) Info(message string, v ...*LogMap) {
	l.Output(2, "info", message, v...)
}

// Logs message and any LogMaps at level="fatal", then calls os.Exit(1)
func (l *Logger) Fatal(message string, v ...*LogMap) {
	l.Output(2, "fatal", message, v...)
	os.Exit(1)
}

// New creates a new Logger.
// The debug argument specifies whether debug should be set or not.
func New(out io.Writer, flags uint64) *Logger {
	return &Logger{
		out:   out,
		flags: flags,
	}
}

// default Logger
var DefaultLogger = New(os.Stderr, Lstd)

func SetFlags(flags uint64) {
	DefaultLogger.SetFlags(flags)
}

// Logs to the default Logger. See Logger.Debug
func Debug(message string, v ...*LogMap) {
	if DefaultLogger.HasDebug() {
		DefaultLogger.Output(2, "debug", message, v...)
	}
}

// Logs to the default Logger. See Logger.Info
func Info(message string, v ...*LogMap) {
	DefaultLogger.Output(2, "info", message, v...)
}

// Logs to the default Logger. See Logger.Fatal
func Fatal(message string, v ...*LogMap) {
	DefaultLogger.Output(2, "fatal", message, v...)
	os.Exit(1)
}
