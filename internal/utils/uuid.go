// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package utils

import (
	"fmt"

	"github.com/google/uuid"
)

// Generate UUID.
// if something gone wrong - returns empty string and error
func GenerateUUID() (string, error) {
	_uuid, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("failed to create new uuid: %v", err)
	}

	return _uuid.String(), nil
}
