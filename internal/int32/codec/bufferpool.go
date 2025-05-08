package codec

import (
	"bytes"
	"sync"
)

// BufferPoolConfig controls how the encoding buffer pool behaves.
type BufferPoolConfig struct {
	InitialCapacity int // starting buffer size in bytes
	MaxCapacity     int // max allowed buffer size to reuse (larger ones are discarded)
}

// Default buffer pool settings
var poolConfig = BufferPoolConfig{
	InitialCapacity: 4096,      // 4KB
	MaxCapacity:     64 * 1024, // 64KB
}

// ConfigureBufferPool allows users to customize buffer pooling behavior.
func ConfigureBufferPool(cfg BufferPoolConfig) {
	poolConfig = cfg
}

var bufferPool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, poolConfig.InitialCapacity))
	},
}

// GetBuffer retrieves a reset buffer from the pool.
func GetBuffer() *bytes.Buffer {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

// PutBuffer returns a buffer to the pool if it's not too large.
func PutBuffer(buf *bytes.Buffer) {
	if buf.Cap() <= poolConfig.MaxCapacity {
		bufferPool.Put(buf)
	}
}
