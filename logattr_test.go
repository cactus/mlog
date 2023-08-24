package mlog

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func BenchmarkLogAttrWriteBuf(b *testing.B) {
	buf := &discardSliceWriter{}
	attrs := make([]*Attr, b.N)
	for i := 0; i < 100; i++ {
		attrs = append(attrs, &Attr{
			Key:   randString(10, false),
			Value: randString(25, true),
		})
	}
	b.ResetTimer()
	for i, attr := range attrs {
		attr.writeBuf(buf)
		if i != len(attrs)-1 {
			buf.WriteByte(' ')
		}
	}
}

func TestLogAttrWriteTo(t *testing.T) {
	attr := Attr{"test", "this is \"a test\" of \t some \n a"}
	buf := &sliceBuffer{make([]byte, 0, 1024)}
	attr.writeBuf(buf)
	n := `test="this is \"a test\" of \t some \n a"`
	l := buf.String()
	assert.Check(t, is.Equal(n, l), "did not match")
}
