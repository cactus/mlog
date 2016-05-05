// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"io/ioutil"
	"testing"
)

func BenchmarkJSONFormatWriterBase(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &JSONFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkJSONFormatWriterStd(b *testing.B) {
	logger := New(ioutil.Discard, Lstd)
	logWriter := &JSONFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkJSONFormatWriterTime(b *testing.B) {
	logger := New(ioutil.Discard, Ltimestamp)
	logWriter := &JSONFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkJSONFormatWriterShortfile(b *testing.B) {
	logger := New(ioutil.Discard, Lshortfile)
	logWriter := &JSONFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkJSONFormatWriterLongfile(b *testing.B) {
	logger := New(ioutil.Discard, Llongfile)
	logWriter := &JSONFormatWriter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkJSONFormatWriterMap(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &JSONFormatWriter{}
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}

func BenchmarkJSONFormatWriterHugeMap(b *testing.B) {
	logger := New(ioutil.Discard, 0)
	logWriter := &JSONFormatWriter{}
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(6, false)] = randString(10, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}
