package codec

// ZigzagEncode32 encodes a signed int32 into a uint32 using ZigZag encoding.
func ZigzagEncode32(n int32) uint32 {
	return uint32((n << 1) ^ (n >> 31))
}

// ZigzagDecode32 decodes a ZigZag-encoded uint32 back into a signed int32.
func ZigzagDecode32(n uint32) int32 {
	return int32((n >> 1) ^ uint32(-(n & 1)))
}
