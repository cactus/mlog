// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"fmt"
	"io"
	"sort"
)

var (
	// precomputed byte slices, to avoid calling io.WriteString
	// when possible (io.WriteString has an assertion, so just
	// use Write (io.Writer interface) when it makes sense.
	c_Q = []byte{'\\', '"'}
	c_T = []byte{'\\', 't'}
	c_R = []byte{'\\', 'r'}
	c_N = []byte{'\\', 'n'}

	i_SPACE       = []byte{' '}
	i_QUOTE       = []byte{'"'}
	i_EQUAL_QUOTE = []byte{'=', '"'}
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
	// scratch buffer for intermediate writes
	buf := bufPool.Get()
	defer bufPool.Put(buf)

	first := true
	for k, v := range m {
		if first {
			first = false
		} else {
			w.Write(i_SPACE)
		}

		io.WriteString(w, k)
		w.Write(i_EQUAL_QUOTE)

		fmt.Fprint(buf, v)
		// pull out byte slice from buff
		b := buf.Bytes()
		blen := buf.Len()
		p := 0
		for i := 0; i < blen; i++ {
			switch b[i] {
			case '"':
				w.Write(b[p:i])
				w.Write(c_Q)
				p = i + 1
			case '\t':
				w.Write(b[p:i])
				w.Write(c_T)
				p = i + 1
			case '\r':
				w.Write(b[p:i])
				w.Write(c_R)
				p = i + 1
			case '\n':
				w.Write(b[p:i])
				w.Write(c_N)
				p = i + 1
			}
		}
		if p < blen {
			w.Write(b[p:blen])
		}

		w.Write(i_QUOTE)
		// truncate intermediate buf so it is clean for next loop
		buf.Truncate(0)
	}
	// int64 to be compat with io.WriterTo
	return int64(len(m)), nil
}

// SortedWriteTo writes a sorted string representation of
// the Map's key value pairs to w.
func (m Map) SortedWriteTo(w io.Writer) (int64, error) {
	// scratch buffer for intermediate writes
	buf := bufPool.Get()
	defer bufPool.Put(buf)

	keys := m.Keys()
	sort.Strings(keys)

	first := true
	for _, k := range keys {
		if first {
			first = false
		} else {
			w.Write(i_SPACE)
		}

		io.WriteString(w, k)
		w.Write(i_EQUAL_QUOTE)

		fmt.Fprint(buf, m[k])
		b := buf.Bytes()
		blen := buf.Len()
		p := 0
		for i := 0; i < blen; i++ {
			switch b[i] {
			case '"':
				w.Write(b[p:i])
				w.Write(c_Q)
				p = i + 1
			case '\t':
				w.Write(b[p:i])
				w.Write(c_T)
				p = i + 1
			case '\r':
				w.Write(b[p:i])
				w.Write(c_R)
				p = i + 1
			case '\n':
				w.Write(b[p:i])
				w.Write(c_N)
				p = i + 1
			}
		}
		if p < blen {
			w.Write(b[p:blen])
		}

		w.Write(i_QUOTE)
		// truncate intermediate buf so it is clean for next loop
		buf.Truncate(0)
	}
	// int64 to be compat with WriterTo above
	return int64(len(m)), nil
}

// String returns an unsorted string representation of
// the Map's key value pairs.
func (m Map) String() string {
	buf := bufPool.Get()
	defer bufPool.Put(buf)
	m.WriteTo(buf)
	return buf.String()
}

// String returns a sorted string representation of
// the Map's key value pairs.
func (m Map) SortedString() string {
	buf := bufPool.Get()
	defer bufPool.Put(buf)
	m.SortedWriteTo(buf)
	return buf.String()
}
