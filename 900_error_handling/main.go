package main

import (
	"io"
)

type errWriter struct {
	w   io.Writer
	err error //keeping first error
}

func (ew *errWriter) writeln(buf []byte) {
	if ew.err != nil {
		return //skips writing, if earlier was error
	}
	_, ew.err = ew.w.Write(buf) // Writes line and holds any error
}

func main() {
}
