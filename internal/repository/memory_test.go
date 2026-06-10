package repository

import (
	"testing"

	"trackposter/internal/connector/mock"
)

var (
	mockQueue   = NewMemoryQueue()
	mockTrack   = mock.MockTrack.AsTrack()
	mockTrackID = ""
)

func addTrackToQueue(t *testing.T) {
	t.Helper()

	t.Logf("Adding mock track to queue: %v", mockTrack)
	trackID, err := mockQueue.AddTrack(mockTrack)
	if err != nil {
		t.Fatalf("failed to add track: %v", err)
	}
	t.Logf("Track ID: %s", trackID)

	mockTrackID = trackID
}

func findTrack(t *testing.T) {
	t.Helper()

	t.Logf("Getting track back again by id %s", mockTrackID)
	track := mockQueue.TrackByID(mockTrackID)
	if track == nil {
		t.Fatalf("track must not be nil!")
	}

	t.Logf("Returned track: %v", track)

	if track.URL != mockTrack.URL || track.Duration != mockTrack.Duration {
		t.Fatalf(
			"returned track is not the same as added one! %v != %v",
			track,
			mockTrack,
		)
	}
}

func TestMemoryQueueOperations(t *testing.T) {
	t.Parallel()

	// Adding to queue
	addTrackToQueue(t)

	// Getting track by id
	findTrack(t)
}
