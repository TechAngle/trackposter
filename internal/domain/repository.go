// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in
// LICENSE file.
//
// Author: https://github.com/TechAngle

package domain

type Repository interface {
	// Add new track to line. Returns track id and error.
	AddTrack(track *Track) (string, error)

	// Removes track from line by ID
	RemoveTrack(trackID string) error

	// Get track by ID. If nothing is found, returns nil.
	TrackByID(trackID string) *Track

	// Gets tracks queue
	Queue() []*Track
}
