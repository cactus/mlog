// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"io/ioutil"
	"log"
	"math/rand"
	"testing"
)

const (
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBytesAlt = letterBytes + "\"\t\r\n"
	letterIdxBits  = 6                    // 6 bits to represent a letter index
	letterIdxMask  = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax   = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// uses unseeded rand (seed(1))...only use for testing!
func randString(n int, altchars bool) string {
	lb := letterBytes
	if altchars {
		lb = letterBytesAlt
	}
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(lb) {
			b[i] = lb[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func BenchmarkSLoggingBase(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infom("this is a test: %s", m)
	}
}

func BenchmarkSLoggingBaseSortedKeys(b *testing.B) {
	logger := New(ioutil.Discard, Lstd|Lsort)
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infom("this is a test: %s", m)
	}
}

func BenchmarkSLoggingBaseHugeMapUnsortedKeys(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(6, false)] = randString(10, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infom("this is a test: %s", m)
	}
}

func BenchmarkSLoggingBaseHugeMapSortedKeys(b *testing.B) {
	logger := New(ioutil.Discard, Lstd|Lsort)
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(6, false)] = randString(10, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infom("this is a test: %s", m)
	}
}

func BenchmarkSLoggingDatetime(b *testing.B) {
	logger := New(ioutil.Discard, Lstd)
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infom("this is a test: %s", m)
	}
}

func BenchmarkSLoggingShortfile(b *testing.B) {
	logger := New(ioutil.Discard, Lstd|Lshortfile)
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infom("this is a test: %s", m)
	}
}

func BenchmarkSLoggingLongfile(b *testing.B) {
	logger := New(ioutil.Discard, Lstd|Llongfile)
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infom("this is a test: %s", m)
	}
}

func BenchmarkSLoggingDebugWithEnabled(b *testing.B) {
	logger := New(ioutil.Discard, Lstd|Ldebug)
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debugm("this is a test: %s", m)
	}
}

func BenchmarkSLoggingDebugWithDisabled(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debugm("this is a test: %s", m)
	}
}

func BenchmarkSLoggingLikeStdlib(b *testing.B) {
	logger := New(ioutil.Discard, Lstd)
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infom("this is a test: %s", m)
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
