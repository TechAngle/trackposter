// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPythonInstalled(t *testing.T) {
	assert.True(t, IsPythonInstalled())
}
