// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"io/ioutil"
	"testing"
)

func BenchmarkStructuredLogWriterBase(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &StructuredLogWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkStructuredLogWriterStd(b *testing.B) {
	logger := New(ioutil.Discard, Lstd)
	logWriter := &StructuredLogWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkStructuredLogWriterTime(b *testing.B) {
	logger := New(ioutil.Discard, Ltimestamp)
	logWriter := &StructuredLogWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkStructuredLogWriterShortfile(b *testing.B) {
	logger := New(ioutil.Discard, Lshortfile)
	logWriter := &StructuredLogWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkStructuredLogWriterLongfile(b *testing.B) {
	logger := New(ioutil.Discard, Llongfile)
	logWriter := &StructuredLogWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkStructuredLogWriterMap(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &StructuredLogWriter{}
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}

func BenchmarkStructuredLogWriterHugeMapUnsorted(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &StructuredLogWriter{}
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(6, false)] = randString(10, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}

func BenchmarkStructuredLogWriterHugeMapSorted(b *testing.B) {
	logger := New(ioutil.Discard, Lsort)
	logWriter := &StructuredLogWriter{}
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(6, false)] = randString(10, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}
