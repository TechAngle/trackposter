// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package server

import "testing"

func TestValidateAddress(t *testing.T) {
	validOptions := ServerOptions{
		Host: "0.0.0.0",
		Port: 80,
	}
	invalidOptions := ServerOptions{
		Host: "",
		Port: 10,
	}

	validOptionsAddress := "0.0.0.0:80"
	validAddress := validOptions.ValidAddress()
	if validAddress != validOptionsAddress {
		t.Fatalf("invalid address returned: %s != %s", validAddress, validOptionsAddress)
	}

	invalidAddress := invalidOptions.ValidAddress()
	if invalidAddress == ":10" {
		t.Fatalf("invalid address must be formatted: %s == :10", invalidAddress)
	}

}
