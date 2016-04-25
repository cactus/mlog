// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"sort"
	"strings"
)

const (
	Ltimestamp    FlagSet = 1 << iota // log the date+time stamp
	Lmicroseconds                     // use microsecond timestamp granularity in Ltimestamp
	Lnanoseconds                      // use nanosecond timestamp granularity in Ltimestamp. overrides Lmicroseconds
	Llevel                            // log message level
	Llongfile                         // log file path and line number: /a/b/c/d.go:23
	Lshortfile                        // log file name and line number: d.go:23. overrides Llongfile
	Lsort                             // sort Map key value pairs in output
	Ldebug                            // enable debug level log
	Lstd          = Ltimestamp | Llevel | Lsort
)

var flagNames = map[FlagSet]string{
	Ltimestamp:    "Ltimestamp",
	Lmicroseconds: "Lmicroseconds",
	Lnanoseconds:  "Lnanoseconds",
	Llevel:        "Llevel",
	Llongfile:     "Llongfile",
	Lshortfile:    "Lshortfile",
	Lsort:         "Lsort",
	Ldebug:        "Ldebug",
}

type FlagSet uint64

func (f *FlagSet) Has(p FlagSet) bool {
	if *f&p != 0 {
		return true
	}
	return false
}

func (f FlagSet) GoString() string {
	s := make([]byte, 0)
	var p uint64
	for p = 64; p > 0; p >>= 1 {
		if f&FlagSet(p) != 0 {
			s = append(s, '1')
		} else {
			s = append(s, '0')
		}
	}
	return string(s)
}

func (f FlagSet) String() string {
	flags := make([]string, 0)
	for k, v := range flagNames {
		if f&k != 0 {
			flags = append(flags, v)
		}
	}
	sort.Strings(flags)
	return "FlagSet(" + strings.Join(flags, "|") + ")"
}
