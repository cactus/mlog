package mlog

import (
	"bytes"
	"time"
)

func intpow10(e int) int {
	y := 1
	x := 10
	for e > 1 {
		if (e % 2) == 0 {
			x = x * x
			e = e / 2
		} else {
			y = x * y
			x = x * x
			e = (e - 1) / 2
		}
	}
	return x * y
}

func itoaw(b *bytes.Buffer, i int, wid int) {
	if i <= 0 {
		for wid > 0 {
			b.WriteByte('0')
			wid--
		}
		return
	}

	pm := intpow10(wid - 1)
	for pm > 1 {
		b.WriteByte(byte('0' + (i / pm)))
		i = i % pm
		pm = pm / 10
	}
	b.WriteByte(byte('0' + i))
}

func writeTime(buf *bytes.Buffer, t *time.Time, flags FlagSet) {
	year, month, day := t.Date()
	itoaw(buf, year, 4)
	buf.WriteByte('-')
	itoaw(buf, int(month), 2)
	buf.WriteByte('-')
	itoaw(buf, day, 2)
	buf.WriteByte('T')

	hour, min, sec := t.Clock()
	itoaw(buf, hour, 2)
	buf.WriteByte(':')
	itoaw(buf, min, 2)
	buf.WriteByte(':')
	itoaw(buf, sec, 2)

	switch {
	case flags&Lmicroseconds != 0:
		buf.WriteByte('.')
		itoaw(buf, t.Nanosecond()/1e3, 6)
	case flags&Lnanoseconds != 0:
		buf.WriteByte('.')
		itoaw(buf, t.Nanosecond(), 9)
	}

	_, offset := t.Zone()
	if offset == 0 {
		buf.WriteByte('Z')
	} else {
		if offset < 0 {
			buf.WriteByte('-')
			offset = -offset
		} else {
			buf.WriteByte('+')
		}
		itoaw(buf, offset/3600, 2)
		buf.WriteByte(':')
		itoaw(buf, offset%3600, 2)
	}
}
