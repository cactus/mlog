// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"testing"
)

type tester struct {
	level     string
	pattern   string
	message   string
	arguments []*LogMap
}

var infoTests = []tester{
	{
		"info",
		`level="info" msg="test one %d" x="y"`,
		"test one %d",
		[]*LogMap{
			&LogMap{"x": "y"},
		},
	},
	{
		"info",
		`level="info" msg="test one %d" x="y"`,
		"test one %d",
		[]*LogMap{
			&LogMap{"x": "y"},
		},
	},
	{
		"info",
		`level="info" msg="test one" x="y" y="z" t="u" u="v"`,
		"test one",
		[]*LogMap{
			&LogMap{"x": "y", "y": "z"},
			&LogMap{"t": "u", "u": "v"},
		},
	},
	{
		"info",
		`level="info" msg="test one" x="y" y="z" t="u" u="v"`,
		"test one",
		[]*LogMap{
			&LogMap{"y": "z", "x": "y"},
			&LogMap{"u": "v", "t": "u"},
		},
	},
	{
		"info",
		`level="info" msg="test one" x="1" y="2" z="3"`,
		"test one",
		[]*LogMap{
			&LogMap{
				"x": 1,
				"y": 2,
				"z": 3,
			},
		},
	},
}

var debugTests = []tester{
	{
		"debug",
		`level="debug" msg="test: %s %d"`,
		"test: %s %d",
		[]*LogMap{},
	},
}

func testInfo(t *testing.T, logger *Logger, level, message, pattern string, arguments []*LogMap) {
	buf := &bytes.Buffer{}
	logger.out = buf

	switch level {
	case "debug":
		logger.Debug(message, arguments...)
	default:
		logger.Info(message, arguments...)
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

func TestAllInfo(t *testing.T) {
	logger := New(ioutil.Discard, Lbase|Lsort)
	for _, tt := range infoTests {
		testInfo(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
		testInfo(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
	}
}

func TestAllDebug(t *testing.T) {
	logger := New(ioutil.Discard, Lbase|Lsort|Ldebug)
	for _, tt := range debugTests {
		testInfo(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
		testInfo(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
	}
}

func TestOnce(t *testing.T) {
	logger := New(ioutil.Discard, Ldatetime)
	logger.Info("test this", &LogMap{"test": 1})
}
