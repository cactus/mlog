package mlog

import "time"

func writeTime(sb intSliceWriter, t *time.Time, flags FlagSet) {
	year, month, day := t.Date()
	sb.AppendIntWidth(year, 4)
	sb.WriteByte('-')
	sb.AppendIntWidth(int(month), 2)
	sb.WriteByte('-')
	sb.AppendIntWidth(day, 2)

	sb.WriteByte('T')

	hour, min, sec := t.Clock()
	sb.AppendIntWidth(hour, 2)
	sb.WriteByte(':')
	sb.AppendIntWidth(min, 2)
	sb.WriteByte(':')
	sb.AppendIntWidth(sec, 2)

	sb.WriteByte('.')
	sb.AppendIntWidth(t.Nanosecond(), 9)

	_, offset := t.Zone()
	if offset == 0 {
		sb.WriteByte('Z')
	} else {
		if offset < 0 {
			sb.WriteByte('-')
			offset = -offset
		} else {
			sb.WriteByte('+')
		}
		sb.AppendIntWidth(offset/3600, 2)
		sb.WriteByte(':')
		sb.AppendIntWidth(offset%3600, 2)
	}
}

var (
	// http://maia.usno.navy.mil/ser7/tai-utc.dat
	// http://www.stjarnhimlen.se/comp/time.html
	tia64nDifferences = []struct {
		t      time.Time
		offset int64
	}{
		{time.Date(1972, 1, 1, 0, 0, 0, 0, time.UTC), 10},
		{time.Date(1972, 7, 1, 0, 0, 0, 0, time.UTC), 11},
		{time.Date(1973, 1, 1, 0, 0, 0, 0, time.UTC), 12},
		{time.Date(1974, 1, 1, 0, 0, 0, 0, time.UTC), 13},
		{time.Date(1975, 1, 1, 0, 0, 0, 0, time.UTC), 14},
		{time.Date(1976, 1, 1, 0, 0, 0, 0, time.UTC), 15},
		{time.Date(1977, 1, 1, 0, 0, 0, 0, time.UTC), 16},
		{time.Date(1978, 1, 1, 0, 0, 0, 0, time.UTC), 17},
		{time.Date(1979, 1, 1, 0, 0, 0, 0, time.UTC), 18},
		{time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), 19},
		{time.Date(1981, 7, 1, 0, 0, 0, 0, time.UTC), 20},
		{time.Date(1982, 7, 1, 0, 0, 0, 0, time.UTC), 21},
		{time.Date(1983, 7, 1, 0, 0, 0, 0, time.UTC), 22},
		{time.Date(1985, 7, 1, 0, 0, 0, 0, time.UTC), 23},
		{time.Date(1988, 1, 1, 0, 0, 0, 0, time.UTC), 24},
		{time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), 25},
		{time.Date(1991, 1, 1, 0, 0, 0, 0, time.UTC), 26},
		{time.Date(1992, 7, 1, 0, 0, 0, 0, time.UTC), 27},
		{time.Date(1993, 7, 1, 0, 0, 0, 0, time.UTC), 28},
		{time.Date(1994, 7, 1, 0, 0, 0, 0, time.UTC), 29},
		{time.Date(1996, 1, 1, 0, 0, 0, 0, time.UTC), 30},
		{time.Date(1997, 7, 1, 0, 0, 0, 0, time.UTC), 31},
		{time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC), 32},
		{time.Date(2006, 1, 1, 0, 0, 0, 0, time.UTC), 33},
		{time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC), 34},
		{time.Date(2012, 7, 1, 0, 0, 0, 0, time.UTC), 35},
		{time.Date(2015, 7, 1, 0, 0, 0, 0, time.UTC), 36},
	}
	tia64nSize = len(tia64nDifferences)
)

func writeTimeTAI64N(sb intSliceWriter, t *time.Time, flags FlagSet) {
	sb.WriteString("@4")
	offset := int64(0)
	tu := t.UTC()
	for i := tia64nSize - 1; i >= 0; i-- {
		if tu.Before(tia64nDifferences[i].t) {
			continue
		} else {
			offset = tia64nDifferences[i].offset
			break
		}
	}
	sb.AppendIntWidthHex(tu.Unix()+offset, 15)
	sb.AppendIntWidthHex(int64(tu.Nanosecond()), 8)
}
