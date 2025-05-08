package varint

// VarintSize returns how many bytes are needed to encode u as varint.
func VarintSize(u uint64) int {
	size := 1
	for u >= 0x80 {
		u >>= 7
		size++
	}
	return size
}

// EncodeVarint32 encodes a uint32 into a provided buffer and returns number of bytes written.
// Caller must ensure buf has enough capacity (at least 5 bytes).
func EncodeVarint32(buf []byte, u uint32) int {
	i := 0
	for u >= 0x80 {
		buf[i] = byte(u) | 0x80
		u >>= 7
		i++
	}
	buf[i] = byte(u)
	return i + 1
}

// DecodeVarint32 decodes a varint-encoded uint32 from buf.
// Returns decoded uint32 value and number of bytes read.
func DecodeVarint32(buf []byte) (uint32, int) {
	var x uint32
	var s uint
	for i, b := range buf {
		if b < 0x80 {
			if i > 4 {
				break // overflow
			}
			return x | uint32(b)<<s, i + 1
		}
		x |= uint32(b&0x7F) << s
		s += 7
	}
	return 0, 0 // malformed input
}
