package codec

import (
	"bytes"

	"github.com/soumya-codes/flashwire/internal/varint"
)

// Writer accumulates binary data in an internal buffer.
type Writer struct {
	buf *bytes.Buffer
}

// NewWriter returns a new Writer with an empty buffer.
func NewWriter() *Writer {
	return &Writer{
		buf: GetBuffer(),
	}
}

func NewWriterFromBuffer(buf *bytes.Buffer) *Writer {
	return &Writer{
		buf: buf,
	}
}

func (w *Writer) Reset() {
	w.buf.Reset()
}

// Bytes returns the current contents of the Writer's buffer.
func (w *Writer) Bytes() []byte {
	return w.buf.Bytes()
}

// WriteInt32 writes an int32 using ZigZag + Varint encoding.
func (w *Writer) WriteInt32(val int32) error {
	u := ZigzagEncode32(val)
	var buf [5]byte
	n := varint.EncodeVarint32(buf[:], u)
	_, err := w.buf.Write(buf[:n])
	return err
}
