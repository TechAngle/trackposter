// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://codeberg.com/TechAngle

package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"trackposter/internal/server/models"

	"github.com/stretchr/testify/assert"
)

// get mock track
func mockTrack() *models.Track {
	mockTrack := new(models.Track)

	mockTrack.Title = "Paracosm"
	mockTrack.Author = "Xtrullor"
	mockTrack.Duration = 12345678
	mockTrack.URL = "https://soundcloud.com/xtrullor/paracosm"

	return mockTrack
}

func TestGetNonExistingTrack(t *testing.T) {
	w := httptest.NewRecorder()

	// trying to get non-existing track
	req := httptest.NewRequest("GET", "/api/tracks/track/12345", nil)
	testServer.router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestAddTrack(t *testing.T) {
	// using random track
	track := mockTrack()

	w := httptest.NewRecorder()

	trackJson, _ := json.Marshal(track)
	req := httptest.NewRequest("POST", "/api/tracks/addTrack", bytes.NewReader(trackJson))
	testServer.router.ServeHTTP(w, req)

	// if we added track
	assert.Equal(t, http.StatusAccepted, w.Code)

	// if returned track id not empty
	body, err := unmarshalBody[models.AddTrackResponse](w.Body)
	if err != nil {
		t.Fatalf("cannot unmashal body: %v", err)
	}

	assert.NotEqual(t, body.TrackID, "")

	t.Logf("Track added: trackId=%v", body.TrackID)
}

func TestAddInvalidTrack(t *testing.T) {

	w := httptest.NewRecorder()

	// using just empty track
	track := models.Track{}
	trackJson, _ := json.Marshal(track)

	req := httptest.NewRequest("PUT", "/api/tracks/addTrack", bytes.NewReader(trackJson))
	testServer.router.ServeHTTP(w, req)

	assert.NotEqual(t, w.Code, http.StatusOK)
}

func TestQueue(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/api/tracks/queue", nil)
	testServer.router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)

	body, err := unmarshalBody[models.TracksListResponse](w.Body)
	if err != nil {
		t.Fatalf("cannot unmarshal body: %v", err)
	}

	// it must be 1 because we added track to the repo in TestAddTrack
	assert.Equal(t, 1, len(*body))
}
