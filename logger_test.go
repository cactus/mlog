// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type tester struct {
	level   string
	pattern string
	message string
	extra   Map
}

var infoTests = []tester{
	{
		"info",
		`level="I" msg="test one %d" x="y"`,
		"test one %d",
		Map{"x": "y"},
	},
	{
		"info",
		`level="I" msg="test one %d" x="y"`,
		"test one %d",
		Map{"x": "y"},
	},
	{
		"info",
		`level="I" msg="test one" t="u" u="v" x="y" y="z"`,
		"test one",
		Map{"x": "y", "y": "z", "t": "u", "u": "v"},
	},
	{
		"info",
		`level="I" msg="test one" t="u" u="v" x="y" y="z"`,
		"test one",
		Map{"y": "z", "x": "y", "u": "v", "t": "u"},
	},
	{
		"info",
		`level="I" msg="test one" haz_string="such tests" x="1" y="2" z="3"`,
		"test one",
		Map{
			"x":          1,
			"y":          2,
			"z":          3,
			"haz_string": "such tests",
		},
	},
}

var debugTests = []tester{
	{
		"debug",
		`level="D" msg="test: %s %d"`,
		"test: %s %d",
		Map{},
	},
}

func testInfom(t *testing.T, logger *Logger, level, message, pattern string, extra Map) {
	buf := &bytes.Buffer{}
	logger.out = buf

	switch level {
	case "debug":
		logger.Debugm(message, extra)
	default:
		logger.Infom(message, extra)
	}
	line := buf.String()

	if len(line) == 0 && len(pattern) != 0 {
		t.Errorf("log output should match\n%12s %q\n%12s %q",
			"expected:", pattern[1:len(pattern)-1],
			"actual:", line)
		return
	}

	line = line[0 : len(line)-1]
	pattern = "^" + pattern + "$"
	matched, err := regexp.MatchString(pattern, line)
	if err != nil {
		t.Fatal("pattern did not compile:", err)
	}
	if !matched {
		t.Errorf("log output should match\n%12s %q\n%12s %q",
			"expected:", pattern[1:len(pattern)-1],
			"actual:", line)
	}
}

func TestAllInfom(t *testing.T) {
	logger := New(ioutil.Discard, Llevel|Lsort)
	for _, tt := range infoTests {
		testInfom(t, logger, tt.level, tt.message, tt.pattern, tt.extra)
		testInfom(t, logger, tt.level, tt.message, tt.pattern, tt.extra)
	}
}

func TestAllDebug(t *testing.T) {
	logger := New(ioutil.Discard, Llevel|Lsort|Ldebug)
	for _, tt := range debugTests {
		testInfom(t, logger, tt.level, tt.message, tt.pattern, tt.extra)
		testInfom(t, logger, tt.level, tt.message, tt.pattern, tt.extra)
	}
}

func TestOnce(t *testing.T) {
	logger := New(ioutil.Discard, Lstd)
	logger.Infom("test this", Map{"test": 1})
}

func TestTimestampLog(t *testing.T) {
	buf := &bytes.Buffer{}

	// test nanoseconds
	logger := New(buf, Lstd|Lnanoseconds)
	tnow := time.Now()
	logger.Info("test this")
	ts := bytes.Split(buf.Bytes()[6:], []byte{'"'})[0]
	tlog, err := time.Parse(time.RFC3339Nano, string(ts))
	assert.Nil(t, err, "Failed to parse time from log")
	assert.WithinDuration(t, tnow, tlog, 2*time.Second, "Time not even close")

	buf.Truncate(0)

	// test microeconds
	logger.SetFlags(Lstd | Lmicroseconds)
	tnow = time.Now()
	logger.Info("test this")
	ts = bytes.Split(buf.Bytes()[6:], []byte{'"'})[0]
	tlog, err = time.Parse(time.RFC3339Nano, string(ts))
	assert.Nil(t, err, "Failed to parse time from log")
	assert.WithinDuration(t, tnow, tlog, 2*time.Second, "Time not even close")

	buf.Truncate(0)

	// test standard (seconds)
	logger.SetFlags(Lstd)
	tnow = time.Now()
	logger.Info("test this")
	ts = bytes.Split(buf.Bytes()[6:], []byte{'"'})[0]
	tlog, err = time.Parse(time.RFC3339Nano, string(ts))
	assert.Nil(t, err, "Failed to parse time from log")
	assert.WithinDuration(t, tnow, tlog, 2*time.Second, "Time not even close")
}
