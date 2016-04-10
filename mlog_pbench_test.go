// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkPLoggingBase(b *testing.B) {
	logger := New(ioutil.Discard, Lbase)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("this is a test: %s", lm)
		}
	})
}

func BenchmarkPLoggingBaseSortedKeys(b *testing.B) {
	logger := New(ioutil.Discard, Lsort)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("this is a test: %s", lm)
		}
	})
}

func BenchmarkPLoggingDatetime(b *testing.B) {
	logger := New(ioutil.Discard, Ldatetime)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("this is a test: %s", lm)
		}
	})
}

func BenchmarkPLoggingShortfile(b *testing.B) {
	logger := New(ioutil.Discard, Lshortfile)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("this is a test: %s", lm)
		}
	})
}

func BenchmarkPLoggingLongfile(b *testing.B) {
	logger := New(ioutil.Discard, Llongfile)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("this is a test: %s", lm)
		}
	})
}

func BenchmarkPLoggingDebugWithEnabled(b *testing.B) {
	logger := New(ioutil.Discard, Ldebug)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Debug("this is a test: %s", lm)
		}
	})
}

func BenchmarkPLoggingDebugWithDisabled(b *testing.B) {
	logger := New(ioutil.Discard, Lbase)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Debug("this is a test: %s", lm)
		}
	})
}

func BenchmarkPLoggingLikeStdlib(b *testing.B) {
	logger := New(ioutil.Discard, Ldatetime)
	lm := &LogMap{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("this is a test: %s", lm)
		}
	})
}

func BenchmarkPStdlibLog(b *testing.B) {
	logger := log.New(ioutil.Discard, "info: ", log.LstdFlags)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Print(`msg="%s" %s="%d"`, "this is a test: %s", "x", 42)
		}
	})
}

func BenchmarkPStdlibLogShortfile(b *testing.B) {
	logger := log.New(ioutil.Discard, "info: ", log.LstdFlags|log.Lshortfile)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Print(`msg="%s" %s="%d"`, "this is a test: %s", "x", 42)
		}
	})
}

func BenchmarkPStdlibLogLongfile(b *testing.B) {
	logger := log.New(ioutil.Discard, "info: ", log.LstdFlags|log.Llongfile)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Print(`msg="%s" %s="%d"`, "this is a test: %s", "x", 42)
		}
	})
}
