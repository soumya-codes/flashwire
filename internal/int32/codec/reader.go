// Package codec implements reading and writing of Flashwire data.
package codec

import (
	"io"

	"github.com/soumya-codes/flashwire/internal/varint"
)

// Reader provides sequential reading of binary data from a byte slice.
type Reader struct {
	data []byte
	pos  int
}

// NewReader returns a new Reader reading from the provided byte slice.
func NewReader(data []byte) *Reader {
	return &Reader{data: data, pos: 0}
}

// readVarint32 reads a varint-encoded uint32 from the Reader.
func (r *Reader) readVarint32() (uint32, error) {
	var result uint32
	var shift uint
	for i := 0; i < 5; i++ {
		if r.pos >= len(r.data) {
			return 0, io.ErrUnexpectedEOF
		}
		b := r.data[r.pos]
		r.pos++
		result |= uint32(b&0x7F) << shift
		if b < 0x80 {
			return result, nil
		}
		shift += 7
	}
	return 0, io.ErrUnexpectedEOF
}

// ReadInt32 reads an int32 that was ZigZag + Varint encoded.
func (r *Reader) ReadInt32() (int32, error) {
	if r.pos >= len(r.data) {
		return 0, io.ErrUnexpectedEOF
	}
	value, n := varint.DecodeVarint32(r.data[r.pos:])
	if n == 0 {
		return 0, io.ErrUnexpectedEOF
	}
	r.pos += n
	return ZigzagDecode32(value), nil
}
