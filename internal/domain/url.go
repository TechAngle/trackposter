// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package domain

import (
	"net/url"
)

// ValidateURL checks if provided URL can be parsed.
func ValidateURL(link string) error {
	_, err := url.Parse(link)
	if err != nil {
		return ErrInvalidURL
	}

	return nil
}
