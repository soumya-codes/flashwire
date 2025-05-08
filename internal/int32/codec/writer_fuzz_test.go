package codec

import "testing"

func FuzzWriteInt32(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(1))
	f.Add(int32(-1))
	f.Add(int32(123456789))
	f.Add(int32(-987654321))

	f.Fuzz(func(t *testing.T, value int32) {
		w := NewWriter()
		if err := w.WriteInt32(value); err != nil {
			t.Fatalf("WriteInt32(%d) unexpected error: %v", value, err)
		}

		r := NewReader(w.Bytes())
		got, err := r.ReadInt32()
		if err != nil {
			t.Fatalf("ReadInt32 failed for written value %d: %v", value, err)
		}
		if got != value {
			t.Fatalf("Round-trip mismatch: wrote %d, read %d", value, got)
		}
	})
}
