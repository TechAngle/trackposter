// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package utils

import "trackposter/internal/domain"

// IsSoundcloudURL checks if provided string matches SoundCloud link RegExp.
func IsSoundcloudURL(link string) bool {
	return domain.SCRegex.Match([]byte(link))
}
