package trace_test

import (
	"bytes"
	"testing"

	"github.com/masa-suzu/go-web/trace"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer

	tracer := trace.New(&buf)

	if tracer == nil {
		t.Error("got nil tracer")
	}

	in := "hello, trace package!"
	want := "hello, trace package!\n"
	tracer.Trace(in)

	got := buf.String()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestOff(t *testing.T) {
	var buf bytes.Buffer
	silent := trace.New(&buf, trace.Off)
	silent.Trace("be quiet")

	got := buf.String()

	if got != "" {
		t.Errorf("got '%s', want ''", got)
	}
}
