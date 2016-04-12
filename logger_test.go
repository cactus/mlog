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
		`level="I" msg="test one" haz_space="such tests" x="1" y="2" z="3"`,
		"test one",
		[]Map{
			Map{
				"x":         1,
				"y":         2,
				"z":         3,
				"haz space": "such tests",
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

func printUint64Bits(u uint64) string {
	s := make([]byte, 0)
	var p uint64
	for p = 64; p > 0; p >>= 1 {
		if u&p != 0 {
			s = append(s, '1')
		} else {
			s = append(s, '0')
		}
	}
	return string(s)
}

func TestFlags(t *testing.T) {
	logger := New(ioutil.Discard, 0)

	expected := Ldatetime | Ldebug
	logger.SetFlags(expected)
	flags := logger.Flags()
	if flags&(expected) == 0 {
		t.Errorf("flags did not match\n%12s %#v\n%12s %#v",
			"expected:", printUint64Bits(expected),
			"actual:", printUint64Bits(flags))
	}

	expected = Ldatetime | Llongfile
	logger.SetFlags(expected)
	flags = logger.Flags()
	if flags&(expected) == 0 {
		t.Errorf("flags did not match\n%12s %#v\n%12s %#v",
			"expected:", printUint64Bits(expected),
			"actual:", printUint64Bits(flags))
	}
}