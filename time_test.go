package mlog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	b := &sliceBuffer{make([]byte, 0, 1024)}
	loc, _ := time.LoadLocation("America/Los_Angeles")

	// test nanos
	tm := time.Date(2016, time.November, 1, 2, 3, 4, 5, loc)
	writeTime(b, &tm, Lnanoseconds)
	assert.Equal(t, `2016-11-01T02:03:04.000000005-07:00`, b.String(),
		"time written incorrectly")

	b.Truncate(0)
	tm = time.Date(2016, time.January, 11, 12, 13, 14, 15, time.UTC)
	writeTime(b, &tm, Lnanoseconds)
	assert.Equal(t, `2016-01-11T12:13:14.000000015Z`, b.String(),
		"time written incorrectly")

	// test micros
	b.Truncate(0)
	tm = time.Date(2016, time.November, 1, 2, 3, 4, 5000, loc)
	writeTime(b, &tm, Lmicroseconds)
	assert.Equal(t, `2016-11-01T02:03:04.000005-07:00`, b.String(),
		"time written incorrectly")

	b.Truncate(0)
	tm = time.Date(2016, time.January, 11, 12, 13, 14, 15000, time.UTC)
	writeTime(b, &tm, Lmicroseconds)
	assert.Equal(t, `2016-01-11T12:13:14.000015Z`, b.String(),
		"time written incorrectly")

}
