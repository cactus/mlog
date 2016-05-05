// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"io/ioutil"
	"testing"
)

func BenchmarkStructuredFormatWriterBase(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &StructuredFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkStructuredFormatWriterStd(b *testing.B) {
	logger := New(ioutil.Discard, Lstd)
	logWriter := &StructuredFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkStructuredFormatWriterTime(b *testing.B) {
	logger := New(ioutil.Discard, Ltimestamp)
	logWriter := &StructuredFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkStructuredFormatWriterShortfile(b *testing.B) {
	logger := New(ioutil.Discard, Lshortfile)
	logWriter := &StructuredFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkStructuredFormatWriterLongfile(b *testing.B) {
	logger := New(ioutil.Discard, Llongfile)
	logWriter := &StructuredFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkStructuredFormatWriterMap(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &StructuredFormatWriter{}
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}

func BenchmarkStructuredFormatWriterHugeMapUnsorted(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &StructuredFormatWriter{}
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(6, false)] = randString(10, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}

func BenchmarkStructuredFormatWriterHugeMapSorted(b *testing.B) {
	logger := New(ioutil.Discard, Lsort)
	logWriter := &StructuredFormatWriter{}
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(6, false)] = randString(10, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}
