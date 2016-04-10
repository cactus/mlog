// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkFSLoggingLikeStdlib(b *testing.B) {
	logger := New(ioutil.Discard, Lstd)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infof("this is a test: %s", "test")
	}
}

func BenchmarkFSStdlibLog(b *testing.B) {
	logger := log.New(ioutil.Discard, "info: ", log.LstdFlags)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Printf("this is a test: %s", "test")
	}
}

func BenchmarkFPLoggingLikeStdlib(b *testing.B) {
	logger := New(ioutil.Discard, Lstd)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Infof("this is a test: %s", "test")
		}
	})
}

func BenchmarkFPStdlibLog(b *testing.B) {
	logger := log.New(ioutil.Discard, "info: ", log.LstdFlags)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Printf("this is a test: %s", "test")
		}
	})
}
