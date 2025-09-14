package mlog

import (
	"testing"

	"github.com/dropwhile/assert"
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
	assert.Equal(t, n, l, "did not match")
}
