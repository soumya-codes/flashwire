package main

import (
	"testing"
)

func TestDataInputMarshalUnmarshalRoundTrip(t *testing.T) {
	original := &DataInput{
		Foo: 42,
		Bar: -17,
	}

	// Encode
	data, err := original.MarshalBinary()
	if err != nil {
		t.Fatalf("MarshalBinary failed: %v", err)
	}

	// Decode into a new object
	var decoded DataInput
	if err := decoded.UnmarshalBinary(data); err != nil {
		t.Fatalf("UnmarshalBinary failed: %v", err)
	}

	// Validate fields
	if original.Foo != decoded.Foo || original.Bar != decoded.Bar {
		t.Errorf("Roundtrip mismatch: original %+v, decoded %+v", original, decoded)
	}
}

func TestDataInputSizeMatchesMarshal(t *testing.T) {
	d := &DataInput{
		Foo: 42,
		Bar: -17,
	}

	expectedSize := d.Size()
	data, err := d.MarshalBinary()
	if err != nil {
		t.Fatalf("MarshalBinary failed: %v", err)
	}

	if expectedSize != len(data) {
		t.Errorf("Size mismatch: Size()=%d but MarshalBinary() produced %d bytes", expectedSize, len(data))
	}
}
