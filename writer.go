package logger

import (
	"bytes"
	"io"
)

// writer combines a buffer and a writer,
// adding a flush function to flush the buffer
// to the underlying writer. This is used
// to avoid file I/O, making the logger faster.
type writer struct {
	*bytes.Buffer
	w io.Writer
}

// Flush writes the buffer contents to the
// underlying writer
func (w writer) Flush() error {
	_, err := io.Copy(w.w, w.Buffer)
	return err
}
