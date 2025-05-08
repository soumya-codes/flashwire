package main

import (
	"testing"

	"github.com/soumya-codes/flashwire/internal/int32/codec"
)

func BenchmarkDataInputMarshalBinary(b *testing.B) {
	d := &DataInput{
		Foo: 12345,
		Bar: -67890,
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = d.MarshalBinary()
	}
}

func BenchmarkDataInputMarshalBinaryBorrowed(b *testing.B) {
	d := &DataInput{
		Foo: 12345,
		Bar: -67890,
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf, err := d.MarshalBinaryBorrowed()
		if err != nil {
			b.Fatal(err)
		}
		codec.PutBuffer(buf) // return it after each use
	}
}

func BenchmarkDataInputUnmarshalBinary(b *testing.B) {
	d := &DataInput{
		Foo: 12345,
		Bar: -67890,
	}
	data, _ := d.MarshalBinary()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var decoded DataInput
		_ = decoded.UnmarshalBinary(data)
	}
}
