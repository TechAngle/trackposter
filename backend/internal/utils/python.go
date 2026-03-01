// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package utils

import "os/exec"

// Check if python installed.
// FIX: Don't work on Unix-like systems
func IsPythonInstalled() bool {
	_, err := exec.LookPath("python")
	return err == nil
}
