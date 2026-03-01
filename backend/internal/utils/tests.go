// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package utils

import (
	"os"
	"testing"
)

// Skips testing if current enviroment is CI.
//
// NOTE: Useful for some functions that check binaries paths
func SkipIfCI(t testing.TB) {
	if os.Getenv("CI") != "" {
		t.Skipf("Skipped %s in CI mode", t.Name())
	}
}
