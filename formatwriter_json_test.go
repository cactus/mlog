// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

var jsonStringTests = map[string]string{
	"generic":           `test`,
	"quote":             `"this"`,
	"r&n":               "te\r\nst",
	"tab":               "\t what",
	"weird chars":       "\u2028 \u2029",
	"other weird chars": `"\u003c\u0026\u003e"`,
	"invalid utf8":      "\xff\xff\xffhello",
}

func TestFormatWriterJSONEncodeStringMap(t *testing.T) {
	b := &bytes.Buffer{}
	for name, s := range jsonStringTests {
		e, err := json.Marshal(s)
		assert.Check(t, err, "%s: json marshal failed", name)

		b.Truncate(0)
		b.WriteByte('"')
		encodeStringJSON(b, s)
		b.WriteByte('"')
		assert.Check(t, is.Equal(string(e), b.String()), "%s: did not match expectation", name)
	}
}
func TestFormatWriterJSONAttrsNil(t *testing.T) {
	logger := New(io.Discard, 0)
	logWriter := &FormatWriterJSON{}
	attrs := make([]*Attr, 10)
	for i := 1; i <= 5; i++ {
		attrs = append(attrs, &Attr{randString(6, false), randString(10, false)})
	}
	for i := 0; i < len(attrs); i++ {
		logWriter.EmitAttrs(logger, 0, "this is a test", attrs...)
	}
}
