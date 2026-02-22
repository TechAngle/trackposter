// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package server

import "testing"

func TestPortValidation(t *testing.T) {
	invalidPort := 24
	validPort := 443

	if isPortValid(invalidPort) {
		t.Fatalf("invalid port is valid! (%d is %v)", invalidPort, isPortValid(invalidPort))
	}

	if !isPortValid(validPort) {
		t.Fatalf("valid port is invalid (%d is %v)", validPort, isPortValid(validPort))
	}
}
