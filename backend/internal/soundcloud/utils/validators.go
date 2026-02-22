// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package utils

import "regexp"

var (
	soundCloudRegex = regexp.MustCompile(`^https://(soundcloud\.com/[A-Za-z0-9\-_]+/[A-Za-z0-9\-_]+|on\.soundcloud\.com/[A-Za-z0-9]+)(\?.*)?$`)
)

// check if string is SoundCloud link
func IsSoundcloudURL(link string) bool {
	return soundCloudRegex.Match([]byte(link))
}
