package models

// /addTrack request model
type AddTrackRequest struct {
	Track Track
}

// /track/{id} model
type GetTrackRequest struct {
	TrackId string `json:"trackId"`
}

// /removeTrack model
type RemoveTrackRequest struct {
	// Track ID
	TrackID string `json:"trackId"`
}

// /addTrack resopnse
type AddTrackResponse struct {
	// Track ID
	TrackID string `json:"trackId,omitempty"`
	// Message
	StatusMessage string `json:"details"`
}

type TracksListResponse []*Track

// /track/{id} response
type GetTrackResponse struct {
	Track Track
}
