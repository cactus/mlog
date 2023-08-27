package mlog

import (
	"testing"
)

func TestTLogger(t *testing.T) {
	tw := &TestingLogWriter{t}
	logger := New(tw, 0)
	logger.Info("test")
	_ = tw
}
