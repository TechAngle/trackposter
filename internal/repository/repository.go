// Package repository provides logic for track queue management.
package repository

import "trackposter/internal/model"

// Repository defines the contract for persisting and retrieving domain
// entities.
type Repository interface {
	// Add new track to line. Returns track id and error.
	AddTrack(track *model.Track) (string, error)

	// Removes track from line by ID
	RemoveTrack(trackID string) error

	// Get track by ID. If nothing is found, returns nil.
	TrackByID(trackID string) *model.Track

	// Gets tracks queue
	Queue() []*model.Track
}
