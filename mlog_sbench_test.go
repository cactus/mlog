// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkSLoggingBase(b *testing.B) {
	logger := New(ioutil.Discard, Lbase)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("this is a test: %s", lm)
	}
}

func BenchmarkSLoggingBaseSortedKeys(b *testing.B) {
	logger := New(ioutil.Discard, Lsort)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("this is a test: %s", lm)
	}
}

func BenchmarkSLoggingDatetime(b *testing.B) {
	logger := New(ioutil.Discard, Ldatetime)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("this is a test: %s", lm)
	}
}

func BenchmarkSLoggingShortfile(b *testing.B) {
	logger := New(ioutil.Discard, Lshortfile)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("this is a test: %s", lm)
	}
}

func BenchmarkSLoggingLongfile(b *testing.B) {
	logger := New(ioutil.Discard, Llongfile)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("this is a test: %s", lm)
	}
}

func BenchmarkSLoggingDebugWithEnabled(b *testing.B) {
	logger := New(ioutil.Discard, Ldebug)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug("this is a test: %s", lm)
	}
}

func BenchmarkSLoggingDebugWithDisabled(b *testing.B) {
	logger := New(ioutil.Discard, Lbase)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug("this is a test: %s", lm)
	}
}

func BenchmarkSLoggingLikeStdlib(b *testing.B) {
	logger := New(ioutil.Discard, Ldatetime)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("this is a test: %s", lm)
	}
}

func BenchmarkSStdlibLog(b *testing.B) {
	logger := log.New(ioutil.Discard, "info: ", log.LstdFlags)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Print(`msg="%s" %s="%d"`, "this is a test: %s", "x", 42)
	}
}

func BenchmarkSStdlibLogShortfile(b *testing.B) {
	logger := log.New(ioutil.Discard, "info: ", log.LstdFlags|log.Lshortfile)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Print(`msg="%s" %s="%d"`, "this is a test: %s", "x", 42)
	}
}

func BenchmarkSStdlibLogLongfile(b *testing.B) {
	logger := log.New(ioutil.Discard, "info: ", log.LstdFlags|log.Llongfile)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Print(`msg="%s" %s="%d"`, "this is a test: %s", "x", 42)
	}
}
