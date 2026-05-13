// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

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

func GetBuilder() *strings.Builder {
	return builderPool.Get().(*strings.Builder)
}

func PutBuilder(b *strings.Builder) {
	b.Reset()
	builderPool.Put(b)
}
