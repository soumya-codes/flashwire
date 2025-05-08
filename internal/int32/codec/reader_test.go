package codec

import (
	"testing"
)

func TestInt32RoundTripReaderWriter(t *testing.T) {
	cases := []struct {
		name  string
		input int32
	}{
		{"zero", 0},
		{"positive", 123456789},
		{"negativeOne", -1},
		{"maxInt32", 2147483647},
		{"minInt32", -2147483648},
		{"smallNegative", -15},
		{"smallPositive", 42},
	}
	for _, tc := range cases {
		tc := tc // capture range var
		t.Run(tc.name, func(t *testing.T) {
			w := NewWriter()
			err := w.WriteInt32(tc.input)
			if err != nil {
				t.Fatalf("WriteInt32 failed: %v", err)
			}
			data := w.Bytes()

			r := NewReader(data)
			got, err := r.ReadInt32()
			if err != nil {
				t.Fatalf("ReadInt32 failed: %v", err)
			}
			if got != tc.input {
				t.Errorf("Roundtrip mismatch: wrote %d, read %d", tc.input, got)
			}
		})
	}
}

func TestReadInt32Sequential(t *testing.T) {
	w := NewWriter()
	err1 := w.WriteInt32(1)
	err2 := w.WriteInt32(2)
	if err1 != nil || err2 != nil {
		t.Fatalf("WriteInt32 failed: %v, %v", err1, err2)
	}

	r := NewReader(w.Bytes())
	val1, err1 := r.ReadInt32()
	val2, err2 := r.ReadInt32()
	if err1 != nil || err2 != nil {
		t.Fatalf("Unexpected error on sequential reads: %v, %v", err1, err2)
	}
	if val1 != 1 || val2 != 2 {
		t.Errorf("Sequential ReadInt32 got (%d, %d), want (1, 2)", val1, val2)
	}
}
