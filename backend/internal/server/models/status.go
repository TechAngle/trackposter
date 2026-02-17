// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://codeberg.com/TechAngle

package models

// default message response with status code
type StatusResponse struct {
	// Message
	StatusMessage string `json:"details"`
}
