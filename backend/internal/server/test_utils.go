// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package server

import (
	"encoding/json"
	"fmt"
	"io"
	"trackposter/internal/repository"
)

var (
	testServer = defaultServer()
)

// creates server with memory queue (for testing)
func defaultServer() *Server {
	server := NewServer()
	server.SetRepository(repository.NewMemoryQueue())

	return server
}

// unmarshal body to given structure
func unmarshalBody[T any](body io.Reader) (*T, error) {
	bodyContent, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err)
	}

	// unmarshaling response
	var response T
	err = json.Unmarshal(bodyContent, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %v", err)
	}

	return &response, nil
}
