// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package utils

import (
	"os"
)

// Check if path valid
func IsPathValid(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}
