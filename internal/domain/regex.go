// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package domain

import "regexp"

var (
	// SCRegex is a RegExp for soundcloud.com URLs
	SCRegex = regexp.MustCompile(`^https://(soundcloud\.com/[A-Za-z0-9\-_]+/[A-Za-z0-9\-_]+|on\.soundcloud\.com/[A-Za-z0-9]+)(\?.*)?$`)
)
