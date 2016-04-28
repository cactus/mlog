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
	logger := New(ioutil.Discard, 0)
	m := Map{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Infom("this is a test: %s", m)
		}
	})
}

func BenchmarkPLoggingBaseSortedKeys(b *testing.B) {
	logger := New(ioutil.Discard, Lstd|Lsort)
	m := Map{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Infom("this is a test: %s", m)
		}
	})
}

func BenchmarkPLoggingDatetime(b *testing.B) {
	logger := New(ioutil.Discard, Lstd)
	m := Map{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Infom("this is a test: %s", m)
		}
	})
}

func BenchmarkPLoggingShortfile(b *testing.B) {
	logger := New(ioutil.Discard, Lstd|Lshortfile)
	m := Map{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Infom("this is a test: %s", m)
		}
	})
}

func BenchmarkPLoggingLongfile(b *testing.B) {
	logger := New(ioutil.Discard, Lstd|Llongfile)
	m := Map{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Infom("this is a test: %s", m)
		}
	})
}

func BenchmarkPLoggingDebugWithEnabled(b *testing.B) {
	logger := New(ioutil.Discard, Lstd|Ldebug)
	m := Map{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Debugm("this is a test: %s", m)
		}
	})
}

func BenchmarkPLoggingDebugWithDisabled(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	m := Map{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Debugm("this is a test: %s", m)
		}
	})
}

func BenchmarkPLoggingLikeStdlib(b *testing.B) {
	logger := New(ioutil.Discard, Lstd)
	m := Map{"x": 42}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Infom("this is a test: %s", m)
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
