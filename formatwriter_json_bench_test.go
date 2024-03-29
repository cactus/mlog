// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"io"
	"testing"
)

func BenchmarkFormatWriterJSONBase(b *testing.B) {
	logger := New(io.Discard, 0)
	logWriter := &FormatWriterJSON{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkFormatWriterJSONStd(b *testing.B) {
	logger := New(io.Discard, Lstd)
	logWriter := &FormatWriterJSON{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkFormatWriterJSONTime(b *testing.B) {
	logger := New(io.Discard, Ltimestamp)
	logWriter := &FormatWriterJSON{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkFormatWriterJSONTimeTAI64N(b *testing.B) {
	logger := New(io.Discard, Ltai64n)
	logWriter := &FormatWriterJSON{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkFormatWriterJSONShortfile(b *testing.B) {
	logger := New(io.Discard, Lshortfile)
	logWriter := &FormatWriterJSON{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkFormatWriterJSONLongfile(b *testing.B) {
	logger := New(io.Discard, Llongfile)
	logWriter := &FormatWriterJSON{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", nil)
	}
}

func BenchmarkFormatWriterJSONMap(b *testing.B) {
	logger := New(io.Discard, 0)
	logWriter := &FormatWriterJSON{}
	m := Map{"x": 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}

func BenchmarkFormatWriterJSONAttrs(b *testing.B) {
	logger := New(io.Discard, 0)
	logWriter := &FormatWriterJSON{}
	attr := Attr{"x", 42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.EmitAttrs(logger, 0, "this is a test", &attr)
	}
}

func BenchmarkFormatWriterJSONHugeMap(b *testing.B) {
	logger := New(io.Discard, 0)
	logWriter := &FormatWriterJSON{}
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(6, false)] = randString(10, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.Emit(logger, 0, "this is a test", m)
	}
}

func BenchmarkFormatWriterJSONHugeAttrs(b *testing.B) {
	logger := New(io.Discard, 0)
	logWriter := &FormatWriterJSON{}
	attrs := make([]*Attr, 0, 100)
	for i := 1; i <= 100; i++ {
		attrs = append(attrs, &Attr{randString(6, false), randString(10, false)})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logWriter.EmitAttrs(logger, 0, "this is a test", attrs...)
	}
}
