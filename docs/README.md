# ⚡ Flashwire

Flashwire is a **zero-alloc, high-throughput serialization system** written in Go.  
Designed as a modern alternative to Protobuf for internal use cases where:

- performance
- memory control
- codegen flexibility(Needs Work)
- and minimal runtime metadata

are all top priorities.

---

## ✨ Features

- ⚡ **Zero-allocation encoding** via pooled buffer reuse
- ✅ `MarshalBinary()` (safe) and `MarshalBinaryBorrowed()` (zero-alloc)
- 🧱 Clean Go codegen (no reflection, no protoimpl)

---
## 📈 Benchmark Comparison: Flashwire vs Protobuf (Single `int32` Field)

Benchmark results collected on **Apple M2 Max**, Go 1.22+

| Operation                  | Method                         | ns/op   | B/op | Allocs/op | Notes                     |
|---------------------------|--------------------------------|---------|------|------------|---------------------------|
| Encode (zero-alloc)       | Flashwire MarshalBinaryBorrowed | 27.58   | 0    | 0          | Fastest, no copy          |
| Encode                   | Protobuf Marshal               | 55.03   | 3    | 1          | Inlined struct write      |
| Decode                   | Flashwire UnmarshalBinary      | 10.95   | 0    | 0          | Fully zero alloc          |
| Decode                   | Protobuf Unmarshal             | 82.46   | 48   | 1          | Slower and alloc-heavy    |

✅ **Flashwire is significantly faster on decode**,  
✅ and **is twice as fast as Protobuf** with zero-alloc encode using the `Borrowed` API.

