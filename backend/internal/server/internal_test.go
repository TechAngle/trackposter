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
	"time"
	"trackposter/internal/server/models"

	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/api/status", nil)
	testServer.router.ServeHTTP(w, req)

	// if status not 200
	assert.Equal(t, http.StatusOK, w.Code)
	defer req.Body.Close()

	body, err := unmarshalBody[models.StatusResponse](w.Body)
	if err != nil {
		t.Fatalf("cannot unmarshal body: %v", err)
	}

	// checking
	assert.Equal(t, body.StatusMessage, "ok")
}

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()

	testRequest := &models.PingRequest{
		Timestamp: int(time.Now().Unix()),
	}
	requestJson, _ := json.Marshal(testRequest)

	req := httptest.NewRequest("POST", "/api/ping", bytes.NewReader(requestJson))
	testServer.router.ServeHTTP(w, req)

	// decoding response
	body, err := unmarshalBody[models.PingResponse](w.Body)
	if err != nil {
		t.Fatalf("cannot unmarshal body: %v", err)
	}

	// delta must be greater than 0
	assert.GreaterOrEqual(t, body.Delta, 0)
}
