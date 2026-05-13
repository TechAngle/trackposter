// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package uid

import (
	"fmt"

	"github.com/google/uuid"
)

// New generates new UUID.
// If something went wrong - returns empty string and error.
func New() (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("new uuid: %v", err)
	}

	return uid.String(), nil
}
