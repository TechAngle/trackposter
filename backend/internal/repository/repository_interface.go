// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package repository

import "trackposter/internal/server/models"

type Repository interface {
	// Add new track to line. Returns track id and error.
	AddTrack(*models.Track) (string, error)

	// Removes track from line by ID
	RemoveTrack(string) error

	// Get track by ID. If nothing is found, returns nil.
	TrackByID(string) *models.Track

	// Gets tracks queue
	Queue() []*models.Track
}
