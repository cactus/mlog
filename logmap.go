// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"bytes"
	"fmt"
	"io"
	"sort"
)

/*
// for the love of all that is sane, you probably
// don't really want to use this. only "safe"
// when you *know* that the []byte will never be
// mutated, and you don't care about holding onto the ref
// beyond the string owner lifetime
func stringtoslicebytetmp(s *string) []byte {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(s))
	sh.Len = len(*s)
	sh.Cap = sh.Len
	return *(*[]byte)(unsafe.Pointer(sh))
}
*/

type LogMap map[string]interface{}

func (lm *LogMap) Keys() []string {
	var keys []string
	for k := range *lm {
		keys = append(keys, k)
	}
	return keys
}

func (lm *LogMap) WriteTo(w io.Writer) (int64, error) {
	i := 0
	ilen := len(*lm)
	for k, v := range *lm {
		/*
			// this is a bit grotesque, but it avoids
			// an allocation. Since write will not mutate
			// the string, this *should* be safe.
			p := stringtoslicebytetmp(&k)
			w.Write(p)
			p = nil
		*/
		w.Write([]byte(k))
		w.Write(EQUAL_QUOTE)
		fmt.Fprint(w, v)
		w.Write(QUOTE)
		if i < ilen-1 {
			w.Write(SPACE)
		}
		i++
	}
	// int64 to be compat with io.WriterTo
	return int64(ilen), nil
}

func (lm *LogMap) SortedWriteTo(w io.Writer) (int64, error) {
	keys := lm.Keys()
	sort.Strings(keys)

	i := 0
	ilen := len(keys)
	for _, k := range keys {
		/*
			// this is a bit grotesque, but it avoids
			// an allocation. Since write will not mutate
			// the string, this *should* be safe.
			p := stringtoslicebytetmp(&k)
			w.Write(p)
			p = nil
		*/
		w.Write([]byte(k))
		w.Write(EQUAL_QUOTE)
		fmt.Fprint(w, (*lm)[k])
		w.Write(QUOTE)
		if i < ilen-1 {
			w.Write(SPACE)
		}
		i++
	}
	// int64 to be compat with WriterTo above
	return int64(ilen), nil
}

func (lm *LogMap) String() string {
	var buf bytes.Buffer
	lm.WriteTo(&buf)
	return buf.String()
}

func (lm *LogMap) SortedString() string {
	var buf bytes.Buffer
	lm.SortedWriteTo(&buf)
	return buf.String()
}
