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

type testerF struct {
	level     string
	pattern   string
	message   string
	arguments []interface{}
}

var infoFTests = []testerF{
	{
		"info",
		`level="I" msg="test one"`,
		"test one",
		[]interface{}{},
	},
	{
		"info",
		`level="I" msg="test one 1 2"`,
		"test one %d %s",
		[]interface{}{
			1, "2",
		},
	},
	{
		"info",
		`level="I" msg="test one: \[a b c\]"`,
		"test one: %s",
		[]interface{}{
			[]string{"a", "b", "c"},
		},
	},
}

var debugFTests = []testerF{
	{
		"debug",
		`level="D" msg="test: 1 2"`,
		"test: %s %d",
		[]interface{}{
			"1", 2,
		},
	},
}

func testInfoF(t *testing.T, logger *Logger, level, message, pattern string, arguments []interface{}) {
	buf := &bytes.Buffer{}
	logger.out = buf

	switch level {
	case "debug":
		logger.Debugf(message, arguments...)
	default:
		logger.Infof(message, arguments...)
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

func TestAllFInfo(t *testing.T) {
	logger := New(ioutil.Discard, Llevel|Lsort)
	for _, tt := range infoFTests {
		testInfoF(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
		testInfoF(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
	}
}

func TestAllFDebug(t *testing.T) {
	logger := New(ioutil.Discard, Llevel|Lsort|Ldebug)
	for _, tt := range debugFTests {
		testInfoF(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
		testInfoF(t, logger, tt.level, tt.message, tt.pattern, tt.arguments)
	}
}
