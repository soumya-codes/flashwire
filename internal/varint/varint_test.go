package varint

import (
	"testing"
)

func TestVarintRoundtrip(t *testing.T) {
	values := []uint32{
		0, 1, 127, 128, 255, 300, 16384, 1 << 21, 1 << 28,
	}

	for _, v := range values {
		var buf [5]byte
		n := EncodeVarint32(buf[:], v)
		decoded, m := DecodeVarint32(buf[:n])
		if n != m {
			t.Errorf("Mismatched size for %d: encoded %d bytes, decoded %d bytes", v, n, m)
		}
		if decoded != v {
			t.Errorf("Varint roundtrip failed: wrote %d, read back %d", v, decoded)
		}
	}
}

func TestVarintSize(t *testing.T) {
	cases := []struct {
		value uint64
		want  int
	}{
		{0, 1},
		{127, 1},
		{128, 2},
		{16383, 2},
		{16384, 3},
		{2097151, 3},
		{2097152, 4},
		{268435455, 4},
		{268435456, 5},
	}

	for _, c := range cases {
		got := VarintSize(c.value)
		if got != c.want {
			t.Errorf("VarintSize(%d) = %d, want %d", c.value, got, c.want)
		}
	}
}
