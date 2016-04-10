// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateGoroutineUpdate(t *testing.T) {
	t.Parallel()
	d := newiDate()
	n := d.String()
	time.Sleep(2 * time.Second)
	l := d.String()
	assert.NotEqual(t, n, l, "Date did not update as expected: %s == %s", n, l)
}

func TestDateManualUpdate(t *testing.T) {
	t.Parallel()
	d := &iDate{}
	d.Update()
	n := d.String()
	time.Sleep(2 * time.Second)
	d.Update()
	l := d.String()
	assert.NotEqual(t, n, l, "Date did not update as expected: %s == %s", n, l)
}

func TestDateManualUpdateUninitialized(t *testing.T) {
	t.Parallel()
	d := &iDate{}

	n := d.String()
	time.Sleep(2 * time.Second)
	d.Update()
	l := d.String()
	assert.NotEqual(t, n, l, "Date did not update as expected: %s == %s", n, l)
}

func BenchmarkDataString(b *testing.B) {
	d := newiDate()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			d.String()
		}
	})
}
