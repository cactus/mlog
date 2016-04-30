// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"io/ioutil"
	"testing"
)

func BenchmarkJSONLogWriterBase(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &JSONLogWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkJSONLogWriterStd(b *testing.B) {
	logger := New(ioutil.Discard, Lstd)
	logWriter := &JSONLogWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkJSONLogWriterTime(b *testing.B) {
	logger := New(ioutil.Discard, Ltimestamp)
	logWriter := &JSONLogWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkJSONLogWriterShortfile(b *testing.B) {
	logger := New(ioutil.Discard, Lshortfile)
	logWriter := &JSONLogWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkJSONLogWriterLongfile(b *testing.B) {
	logger := New(ioutil.Discard, Llongfile)
	logWriter := &JSONLogWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkJSONLogWriterMap(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &JSONLogWriter{}
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}

func BenchmarkJSONLogWriterHugeMap(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &JSONLogWriter{}
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(6, false)] = randString(10, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}
