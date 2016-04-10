// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Map is a key value element used to pass
// data to the Logger functions.
type Map map[string]interface{}

// Return an unsorted list of keys in the Map as a []string.
func (m Map) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// WriteTo writes an unsorted string representation of
// the Map's key value pairs to w.
func (m Map) WriteTo(w io.Writer) (int64, error) {
	first := true
	for k, v := range m {
		if first {
			first = false
		} else {
			w.Write(i_SPACE)
		}

		w.Write([]byte(strings.Replace(k, " ", "_", -1)))
		w.Write(i_EQUAL_QUOTE)
		fmt.Fprint(w, v)
		w.Write(i_QUOTE)
	}
	// int64 to be compat with io.WriterTo
	return int64(len(m)), nil
}

// SortedWriteTo writes a sorted string representation of
// the Map's key value pairs to w.
func (m Map) SortedWriteTo(w io.Writer) (int64, error) {
	keys := m.Keys()
	sort.Strings(keys)

	first := true
	for _, k := range keys {
		if first {
			first = false
		} else {
			w.Write(i_SPACE)
		}

		w.Write([]byte(strings.Replace(k, " ", "_", -1)))
		w.Write(i_EQUAL_QUOTE)
		fmt.Fprint(w, m[k])
		w.Write(i_QUOTE)
	}
	// int64 to be compat with WriterTo above
	return int64(len(m)), nil
}

// String returns an unsorted string representation of
// the Map's key value pairs.
func (m Map) String() string {
	var buf bytes.Buffer
	m.WriteTo(&buf)
	return buf.String()
}

// String returns a sorted string representation of
// the Map's key value pairs.
func (m Map) SortedString() string {
	var buf bytes.Buffer
	m.SortedWriteTo(&buf)
	return buf.String()
}
