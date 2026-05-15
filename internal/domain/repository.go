package domain

// Repository defines the contract for persisting and retrieving domain
// entities.
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
