package mlog

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkLogMapWriteTo(b *testing.B) {
	m := Map{}
	for i := 1; i <= 100; i++ {
		m[randString(10, false)] = randString(25, true)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.WriteTo(ioutil.Discard)
	}
}

func TestLogMapWriteTo(t *testing.T) {
	m := Map{"test": "this is \"a test\" of \t some \n a"}
	buf := &bytes.Buffer{}
	m.WriteTo(buf)
	n := `test="this is \"a test\" of \t some \n a"`
	l := buf.String()
	assert.Equal(t, n, l, "did not match")

}
