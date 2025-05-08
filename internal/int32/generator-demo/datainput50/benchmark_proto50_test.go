package main

import (
	"testing"

	"github.com/soumya-codes/flashwire/internal/int32/proto"
	"google.golang.org/protobuf/proto"
)

func BenchmarkProtoMarshalTestInt32_50(b *testing.B) {
	m := &benchmarkpb.TestInt32_50{}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(m)
	}
}

func BenchmarkProtoUnmarshalTestInt32_50(b *testing.B) {
	m := &benchmarkpb.TestInt32_50{}
	data, _ := proto.Marshal(m)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var out benchmarkpb.TestInt32_50
		_ = proto.Unmarshal(data, &out)
	}
}
