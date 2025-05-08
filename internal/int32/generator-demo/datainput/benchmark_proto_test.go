package main

import (
	"testing"

	"github.com/soumya-codes/flashwire/internal/int32/proto"
	"google.golang.org/protobuf/proto"
)

func BenchmarkProtoMarshalTestInt32(b *testing.B) {
	m := &benchmarkpb.TestInt32{
		Value: 12345,
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(m)
	}
}

func BenchmarkProtoUnmarshalTestInt32(b *testing.B) {
	m := &benchmarkpb.TestInt32{
		Value: 12345,
	}
	data, _ := proto.Marshal(m)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var out benchmarkpb.TestInt32
		_ = proto.Unmarshal(data, &out)
	}
}
