package main

import (
	"github.com/soumya-codes/flashwire/internal/int32/codec"
	"testing"
)

func BenchmarkDataInput50MarshalBinary(b *testing.B) {
	d := &DataInput50{}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = d.MarshalBinary()
	}
}

func BenchmarkDataInput50MarshalBinaryBorrowed(b *testing.B) {
	d := &DataInput50{}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf, err := d.MarshalBinaryBorrowed()
		if err != nil {
			b.Fatal(err)
		}
		codec.PutBuffer(buf)
	}
}

func BenchmarkDataInput50UnmarshalBinary(b *testing.B) {
	d := &DataInput50{}
	data, _ := d.MarshalBinary()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var decoded DataInput50
		_ = decoded.UnmarshalBinary(data)
	}
}
