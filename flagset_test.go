package mlog

import (
	"io"
	"testing"
)

func TestFlagSet(t *testing.T) {
	logger := New(io.Discard, 0)

	expected := Ltimestamp | Ldebug
	logger.SetFlags(expected)
	flags := logger.Flags()
	t.Log(flags)
	if flags&(expected) == 0 {
		t.Errorf("flags did not match\n%12s %#v\n%12s %#v",
			"expected:", expected.GoString(),
			"actual:", flags.GoString())
	}

	expected = Ltimestamp | Llongfile
	logger.SetFlags(expected)
	flags = logger.Flags()
	if flags&(expected) == 0 {
		t.Errorf("flags did not match\n%12s %#v\n%12s %#v",
			"expected:", expected.GoString(),
			"actual:", flags.GoString())
	}
}
