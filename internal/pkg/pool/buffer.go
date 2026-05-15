package pool

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() any {
		return &bytes.Buffer{}
	},
}

// GetBuffer returns pointer to bytes.Buffer pool.
// If pool type casting has failed it will return new allocated bytes.Buffer.
func GetBuffer() *bytes.Buffer {
	buf, ok := bufferPool.Get().(*bytes.Buffer)
	if !ok {
		return new(bytes.Buffer)
	}

	return buf
}

// PutBuffer resets and puts buffer back.
func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}
