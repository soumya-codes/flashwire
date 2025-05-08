package codec

import (
	"testing"
)

func TestWriteAndReadInt32(t *testing.T) {
	cases := []struct {
		name  string
		input int32
	}{
		{"zero", 0},
		{"positive", 123456789},
		{"maxInt32", 2147483647},
		{"minInt32", -2147483648},
		{"negativeOne", -1},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			w := NewWriter()
			if err := w.WriteInt32(tc.input); err != nil {
				t.Fatalf("WriteInt32 failed: %v", err)
			}

			r := NewReader(w.Bytes())
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

func TestWriteInt32Sequential(t *testing.T) {
	w := NewWriter()
	if err := w.WriteInt32(1); err != nil {
		t.Fatalf("WriteInt32(1) error: %v", err)
	}
	if err := w.WriteInt32(2); err != nil {
		t.Fatalf("WriteInt32(2) error: %v", err)
	}

	r := NewReader(w.Bytes())
	val1, err1 := r.ReadInt32()
	val2, err2 := r.ReadInt32()
	if err1 != nil || err2 != nil {
		t.Fatalf("Unexpected error during sequential ReadInt32: %v, %v", err1, err2)
	}
	if val1 != 1 || val2 != 2 {
		t.Errorf("Sequential ReadInt32 after sequential WriteInt32: got (%d, %d), want (1, 2)", val1, val2)
	}
}
