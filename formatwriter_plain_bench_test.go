// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"io/ioutil"
	"testing"
)

func BenchmarkPlainFormatWriterBase(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &PlainFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkPlainFormatWriterStd(b *testing.B) {
	logger := New(ioutil.Discard, Lstd)
	logWriter := &PlainFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkPlainFormatWriterTime(b *testing.B) {
	logger := New(ioutil.Discard, Ltimestamp)
	logWriter := &PlainFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkPlainFormatWriterShortfile(b *testing.B) {
	logger := New(ioutil.Discard, Lshortfile)
	logWriter := &PlainFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkPlainFormatWriterLongfile(b *testing.B) {
	logger := New(ioutil.Discard, Llongfile)
	logWriter := &PlainFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkPlainFormatWriterMap(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &PlainFormatWriter{}
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}

func BenchmarkPlainFormatWriterHugeMapUnsorted(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &PlainFormatWriter{}
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(6, false)] = randString(10, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}

func BenchmarkPlainFormatWriterHugeMapSorted(b *testing.B) {
	logger := New(ioutil.Discard, Lsort)
	logWriter := &PlainFormatWriter{}
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(6, false)] = randString(10, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}
