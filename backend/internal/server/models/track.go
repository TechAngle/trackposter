// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package models

type Track struct {
	Title    string `json:"trackTitle,omitempty"`
	Author   string `json:"trackAuthor,omitempty"`
	Duration int    `json:"trackDuration,omitempty"`
	URL      string `json:"trackUrl,omitempty"`
}
