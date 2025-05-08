package benchmark

import (
	"testing"

	"github.com/soumya-codes/flashwire/internal/int32/proto"
	"google.golang.org/protobuf/proto"
)

func BenchmarkProtobufMarshalInt32(b *testing.B) {
	msg := &benchmarkpb.TestInt32{Value: 123456789}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(msg)
	}
}

func BenchmarkProtobufUnmarshalInt32(b *testing.B) {
	msg := &benchmarkpb.TestInt32{Value: 123456789}
	data, _ := proto.Marshal(msg)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var out benchmarkpb.TestInt32
		_ = proto.Unmarshal(data, &out)
	}
}
