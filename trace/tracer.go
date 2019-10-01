package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
}

type option func(*tracer)

type tracer struct {
	off bool
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	if t.off {
		return
	}
	_, _ = t.out.Write([]byte(fmt.Sprint(a...)))
	_, _ = t.out.Write([]byte("\n"))
}

func New(w io.Writer, opts ...option) Tracer {
	t := &tracer{out: w}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func Off(t *tracer) {
	t.off = true
}
