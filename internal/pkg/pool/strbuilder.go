package pool

import (
	"strings"
	"sync"
)

var builderPool = sync.Pool{
	New: func() any {
		return &strings.Builder{}
	},
}

// GetBuilder returns pointer to strings.Builder in pool.
// If pool type casting has failed it will return new allocated strings.Builder.
func GetBuilder() *strings.Builder {
	builder, ok := builderPool.Get().(*strings.Builder)
	if !ok {
		return new(strings.Builder)
	}

	return builder
}

// PutBuilder resets builder and puts it back to pool.
func PutBuilder(b *strings.Builder) {
	b.Reset()
	builderPool.Put(b)
}
