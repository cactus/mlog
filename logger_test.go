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
	level     string
	pattern   string
	message   string
	arguments []Map
}

var infoTests = []tester{
	{
		"info",
		`level="I" msg="test one %d" x="y"`,
		"test one %d",
		[]Map{
			Map{"x": "y"},
		},
	},
	{
		"info",
		`level="I" msg="test one %d" x="y"`,
		"test one %d",
		[]Map{
			Map{"x": "y"},
		},
	},
	{
		"info",
		`level="I" msg="test one" x="y" y="z" t="u" u="v"`,
		"test one",
		[]Map{
			Map{"x": "y", "y": "z"},
			Map{"t": "u", "u": "v"},
		},
	},
	{
		"info",
		`level="I" msg="test one" x="y" y="z" t="u" u="v"`,
		"test one",
		[]Map{
			Map{"y": "z", "x": "y"},
			Map{"u": "v", "t": "u"},
		},
	},
	{
		"info",
		`level="I" msg="test one" haz_string="such tests" x="1" y="2" z="3"`,
		"test one",
		[]Map{
			Map{
				"x":          1,
				"y":          2,
				"z":          3,
				"haz_string": "such tests",
			},
		},
	},
}

var debugTests = []tester{
	{
		"debug",
		`level="D" msg="test: %s %d"`,
		"test: %s %d",
		[]Map{},
	},
}

func testInfom(t *testing.T, logger *Logger, level, message, pattern string, arguments []Map) {
	buf := &bytes.Buffer{}
	logger.out = buf

	switch level {
	case "debug":
		logger.Debugm(message, arguments...)
	default:
		logger.Infom(message, arguments...)
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
		testInfom(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
		testInfom(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
	}
}

func TestAllDebug(t *testing.T) {
	logger := New(ioutil.Discard, Llevel|Lsort|Ldebug)
	for _, tt := range debugTests {
		testInfom(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
		testInfom(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
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
