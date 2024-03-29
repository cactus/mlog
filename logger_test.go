// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/assert/opt"
	"gotest.tools/v3/golden"
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestLoggerMsgs(t *testing.T) {
	var infoTests = map[string]struct {
		flags   FlagSet
		method  string
		message string
		extra   interface{}
	}{
		"infom1":  {Llevel | Lsort, "infom", "test", Map{"x": "y"}},
		"infom2":  {Llevel | Lsort, "infom", "test", Map{"x": "y", "y": "z", "t": "u", "u": "v"}},
		"infom3":  {Llevel | Lsort, "infom", "test", Map{"y": "z", "x": "y", "u": "v", "t": "u"}},
		"infom4":  {Llevel | Lsort, "infom", "test", Map{"x": 1, "y": 2, "z": 3, "haz_string": "such tests"}},
		"debug1":  {Llevel | Lsort | Ldebug, "debugm", "test", nil},
		"debug2":  {Llevel | Lsort | Ldebug, "debugm", "test", nil},
		"infof1":  {Llevel, "infof", "test: %d", 5},
		"infof2":  {Llevel, "infof", "test: %s", "test"},
		"infof3":  {Llevel, "infof", "test: %s %s", []interface{}{"test", "pickles"}},
		"infox1":  {Llevel, "infox", "test", []*Attr{{"x", "y"}}},
		"infox2":  {Llevel, "infox", "test", []*Attr{{"x", "y"}, {"y", "z"}}},
		"infox3":  {Llevel, "infox", "test", nil},
		"debugx1": {Llevel | Ldebug, "debugx", "test", []*Attr{{"x", "y"}}},
		"debugx2": {Llevel | Ldebug, "debugx", "test", []*Attr{{"x", "y"}, {"y", "z"}}},
		"debugx3": {Llevel | Ldebug, "debugx", "test", nil},
	}

	buf := &bytes.Buffer{}
	logger := New(io.Discard, Llevel|Lsort)
	logger.out = buf

	for name, tt := range infoTests {
		buf.Truncate(0)
		logger.flags = uint64(tt.flags)

		switch tt.method {
		case "debugx":
			m, ok := tt.extra.([]*Attr)
			if !ok && tt.extra != nil {
				t.Errorf("%s: failed type assertion", name)
				continue
			}
			logger.Debugx(tt.message, m...)
		case "infox":
			m, ok := tt.extra.([]*Attr)
			if !ok && tt.extra != nil {
				t.Errorf("%s: failed type assertion", name)
				continue
			}
			logger.Infox(tt.message, m...)
		case "debugm":
			m, ok := tt.extra.(Map)
			if !ok && tt.extra != nil {
				t.Errorf("%s: failed type assertion", name)
				continue
			}
			logger.Debugm(tt.message, m)
		case "infom":
			m, ok := tt.extra.(Map)
			if !ok && tt.extra != nil {
				t.Errorf("%s: failed type assertion", name)
				continue
			}
			logger.Infom(tt.message, m)
		case "debug":
			logger.Debug(tt.message)
		case "info":
			logger.Info(tt.message)
		case "debugf":
			if i, ok := tt.extra.([]interface{}); ok {
				logger.Debugf(tt.message, i...)
			} else {
				logger.Debugf(tt.message, tt.extra)
			}
		case "infof":
			if i, ok := tt.extra.([]interface{}); ok {
				logger.Infof(tt.message, i...)
			} else {
				logger.Infof(tt.message, tt.extra)
			}
		default:
			t.Errorf("%s: not sure what to do", name)
			continue
		}

		goldenFixture := fmt.Sprintf("test_logger_msgs.%s.golden", name)
		golden.AssertBytes(t, buf.Bytes(), goldenFixture, "%s: did not match expectation", name)
	}

}

func TestLoggerTimestamp(t *testing.T) {
	buf := &bytes.Buffer{}

	// test nanoseconds
	logger := New(buf, Lstd)
	tnow := time.Now()
	logger.Info("test this")
	ts := bytes.Split(buf.Bytes()[6:], []byte{'"'})[0]
	tlog, err := time.Parse(time.RFC3339Nano, string(ts))
	assert.Check(t, err, "Failed to parse time from log")
	assert.Assert(t, is.DeepEqual(tnow, tlog, opt.TimeWithThreshold(2*time.Second)), "Time not even close")

	buf.Truncate(0)

	// test microeconds
	logger.SetFlags(Lstd)
	tnow = time.Now()
	logger.Info("test this")
	ts = bytes.Split(buf.Bytes()[6:], []byte{'"'})[0]
	tlog, err = time.Parse(time.RFC3339Nano, string(ts))
	assert.Check(t, err, "Failed to parse time from log")
	assert.Assert(t, is.DeepEqual(tnow, tlog, opt.TimeWithThreshold(2*time.Second)), "Time not even close")

	buf.Truncate(0)

	// test standard (seconds)
	logger.SetFlags(Lstd)
	tnow = time.Now()
	logger.Info("test this")
	ts = bytes.Split(buf.Bytes()[6:], []byte{'"'})[0]
	tlog, err = time.Parse(time.RFC3339Nano, string(ts))
	assert.Check(t, err, "Failed to parse time from log")
	assert.Assert(t, is.DeepEqual(tnow, tlog, opt.TimeWithThreshold(2*time.Second)), "Time not even close")
}

func TestPanics(t *testing.T) {
	var infoTests = map[string]struct {
		flags   FlagSet
		method  string
		message string
		extra   interface{}
	}{
		"panic":  {Llevel | Lsort, "panic", "test", nil},
		"panicf": {Llevel | Lsort, "panicf", "test: %d", 5},
		"panicm": {Llevel | Lsort, "panicm", "test", Map{"x": "y"}},
	}

	buf := &bytes.Buffer{}
	logger := New(io.Discard, Llevel|Lsort)
	logger.out = buf

	for name, tt := range infoTests {
		buf.Truncate(0)
		logger.flags = uint64(tt.flags)

		switch tt.method {
		case "panicm":
			m, ok := tt.extra.(Map)
			if !ok && tt.extra != nil {
				t.Errorf("%s: failed type assertion", name)
				continue
			}
			assertPanic(t, func() {
				logger.Panicm(tt.message, m)
			})
		case "panicf":
			assertPanic(t, func() {
				logger.Panicf(tt.message, tt.extra)
			})
		case "panic":
			assertPanic(t, func() {
				logger.Panic(tt.message)
			})
		default:
			t.Errorf("%s: not sure what to do", name)
			continue
		}

		goldenFixture := fmt.Sprintf("test_logger_msgs.%s.golden", name)
		golden.AssertBytes(t, buf.Bytes(), goldenFixture, "%s: did not match expectation", name)
	}
}
