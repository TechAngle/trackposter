// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://codeberg.com/TechAngle

package repository

import (
	"testing"
	"trackposter/internal/server/models"
)

var (
	mockQueue   = NewMemoryQueue()
	mockTrack   = new(models.Track)
	mockTrackId = ""
)

func setBasicMockTrackData(mockTrack *models.Track) {
	mockTrack.Title = "Hello World"
	mockTrack.Author = "TechAngle"
	mockTrack.Duration = 12345678
	mockTrack.URL = "https://codeberg.org/TechAngle"
}

func addTrackToQueue(t *testing.T) {
	t.Logf("Adding mock track to queue: %v", mockTrack)
	trackId, err := mockQueue.AddTrack(mockTrack)
	if err != nil {
		t.Fatalf("failed to add track: %v", err)
	}
	t.Logf("Track ID: %s", trackId)

	mockTrackId = trackId
}

func findTrack(t *testing.T) {
	t.Logf("Getting track back again by id %s", mockTrackId)
	track := mockQueue.TrackByID(mockTrackId)
	if track == nil {
		t.Fatalf("track mustn't be nil!")
	}

	t.Logf("Returned track: %v", track)

	if track.URL != mockTrack.URL || track.Duration != mockTrack.Duration {
		t.Fatalf("returned track is not the same as added one! %v != %v", track, mockTrack)
	}
}

func TestMemoryQueueOperations(t *testing.T) {
	setBasicMockTrackData(mockTrack)

	// Adding to queue
	addTrackToQueue(t)

	// Getting track by id
	findTrack(t)
}
