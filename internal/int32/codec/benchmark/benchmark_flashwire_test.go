package benchmark

import (
	"testing"

	"github.com/soumya-codes/flashwire/internal/int32/codec"
)

func BenchmarkFlashwireWriteInt32(b *testing.B) {
	w := codec.NewWriter()
	b.ReportAllocs() // track memory allocations
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = w.WriteInt32(123456789)
		w.Reset() // Important: reset buffer after each write to avoid growing slice
	}
}

func BenchmarkFlashwireReadInt32(b *testing.B) {
	w := codec.NewWriter()
	_ = w.WriteInt32(123456789)
	data := w.Bytes()

	b.ReportAllocs() // track memory allocations
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r := codec.NewReader(data)
		_, _ = r.ReadInt32()
	}
}
