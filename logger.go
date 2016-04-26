// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	bufPool = newSliceBufferPool()
)

// A Logger represents a logging object, that embeds log.Logger, and
// provides support for a toggle-able debug flag.
type Logger struct {
	mu    sync.Mutex // ensures atomic writes are synchronized
	out   io.Writer
	flags uint64
}

func (l *Logger) Output(depth int, level string, message string, data ...Map) {
	sb := bufPool.Get()
	defer bufPool.Put(sb)

	flags := FlagSet(atomic.LoadUint64(&l.flags))

	// if time is being logged, handle time as soon as possible
	if flags&Ltimestamp != 0 {
		t := time.Now()
		sb.WriteString(`time="`)
		writeTime(sb, &t, flags)
		sb.WriteString(`" `)
	}

	if flags&Llevel != 0 {
		sb.WriteString(`level="`)
		sb.WriteString(level)
		sb.WriteString(`" `)
	}

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

		sb.WriteString(`caller="`)
		sb.WriteString(file)
		sb.WriteByte(':')
		sb.AppendIntWidth(line, 0)
		sb.WriteString(`" `)
	}

	if flags != 0 {
		sb.WriteString(`msg="`)
	}
	// as a kindness, strip any newlines off the end of the string
	for i := len(message) - 1; i > 0; i-- {
		if message[i] == '\n' {
			message = message[:i]
		} else {
			break
		}
	}
	sb.WriteString(message)
	if flags != 0 {
		sb.WriteByte('"')
	}

	if len(data) > 0 {
		for _, e := range data {
			sb.WriteByte(' ')
			if flags&Lsort != 0 {
				e.SortedWriteTo(sb)
			} else {
				e.WriteTo(sb)
			}
		}
	}

	sb.WriteByte('\n')

	// lock writing to serialize log output (no scrambled log lines)
	l.mu.Lock()
	defer l.mu.Unlock()
	sb.WriteTo(l.out)
}

func (l *Logger) Flags() FlagSet {
	return FlagSet(atomic.LoadUint64(&l.flags))
}

func (l *Logger) SetFlags(flags FlagSet) {
	l.mu.Lock()
	defer l.mu.Unlock()
	atomic.StoreUint64(&l.flags, uint64(flags))
}

func (l *Logger) HasDebug() bool {
	flags := FlagSet(atomic.LoadUint64(&l.flags))
	return flags&Ldebug != 0
}

// Debugm conditionally logs message and any Map elements at level="debug".
// If the Logger does not have the Ldebug flag, nothing is logged.
func (l *Logger) Debugm(message string, v ...Map) {
	if l.HasDebug() {
		l.Output(2, "D", message, v...)
	}
}

// Infom logs message and any Map elements at level="info".
func (l *Logger) Infom(message string, v ...Map) {
	l.Output(2, "I", message, v...)
}

// Printm logs message and any Map elements at level="info".
func (l *Logger) Printm(message string, v ...Map) {
	l.Output(2, "I", message, v...)
}

// Fatalm logs message and any Map elements at level="fatal", then calls
// os.Exit(1)
func (l *Logger) Fatalm(message string, v ...Map) {
	l.Output(2, "F", message, v...)
	os.Exit(1)
}

// Debugf formats and conditionally logs message at level="debug".
// If the Logger does not have the Ldebug flag, nothing is logged.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.HasDebug() {
		l.Output(2, "D", fmt.Sprintf(format, v...))
	}
}

// Infof formats and logs message at level="info".
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(2, "I", fmt.Sprintf(format, v...))
}

// Printf formats and logs message at level="info".
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Output(2, "I", fmt.Sprintf(format, v...))
}

// Fatalf formats and logs message at level="fatal", then calls
// os.Exit(1)
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(2, "F", fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Debug conditionally logs message at level="debug".
// If the Logger does not have the Ldebug flag, nothing is logged.
func (l *Logger) Debug(v ...interface{}) {
	if l.HasDebug() {
		l.Output(2, "D", fmt.Sprint(v...))
	}
}

// Info logs message at level="info".
func (l *Logger) Info(v ...interface{}) {
	l.Output(2, "I", fmt.Sprint(v...))
}

// Print logs message at level="info".
func (l *Logger) Print(v ...interface{}) {
	l.Output(2, "I", fmt.Sprint(v...))
}

// Fatal logs message at level="fatal", then calls
// os.Exit(1)
func (l *Logger) Fatal(v ...interface{}) {
	l.Output(2, "F", fmt.Sprint(v...))
	os.Exit(1)
}

// New creates a new Logger.
// The debug argument specifies whether debug should be set or not.
func New(out io.Writer, flags FlagSet) *Logger {
	return &Logger{
		out:   out,
		flags: uint64(flags),
	}
}
