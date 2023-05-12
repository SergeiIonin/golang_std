package main

import (
	"io"
	"os"
)

type Logger struct {
	io.WriteCloser
}

func main() {
	l := Logger{WriteCloser: os.Stdout}
	// methods of WriteCloser are available directly, which is desirable in this case
	_, _ = l.Write([]byte("foo"))
	_ = l.Close()
}
