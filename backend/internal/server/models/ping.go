// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://codeberg.com/TechAngle

package models

// /ping request
type PingRequest struct {
	Timestamp int `json:"timestamp"`
}

// /ping response
type PingResponse struct {
	Delta int `json:"timestamp"`
}
