// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mlog

import (
	"sync"
	"sync/atomic"
	"time"
)

const timeFormat = time.RFC3339

// iDate holds current date stamp formatted for logging
type iDate struct {
	dateValue   atomic.Value
	onceUpdater sync.Once
}

func (h *iDate) String() string {
	stamp := h.dateValue.Load()
	if stamp == nil {
		h.Update()
		return time.Now().UTC().Format(timeFormat)
	}
	return stamp.(string)
}

func (h *iDate) Update() {
	h.dateValue.Store(time.Now().UTC().Format(timeFormat))
}

func newiDate() *iDate {
	d := &iDate{}
	d.Update()
	// spawn a single formattedDate updater
	d.onceUpdater.Do(func() {
		go func() {
			for range time.Tick(1 * time.Second) {
				d.Update()
			}
		}()
	})
	return d
}

var formattedDate = newiDate()
