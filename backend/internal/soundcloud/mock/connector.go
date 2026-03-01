// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package mock

import (
	"context"
	"fmt"
	"trackposter/internal/soundcloud/types"
	"trackposter/internal/soundcloud/utils"
)

// Implements SoundCloud connector.
//
// Mock connector that returns pre-defined values from config
type MockConnector struct{}

func validUrl(url string) bool {
	return utils.IsSoundcloudURL(url)
}

// Get track metadata.
func (c *MockConnector) TrackMetadataFromURL(ctx context.Context, url string) (*types.TrackMetadata, error) {
	if !validUrl(url) {
		return nil, fmt.Errorf("invalid url provided")
	}

	// cloning and adding url
	mockTrack := MockTrack
	mockTrack.URL = url

	return &mockTrack, nil
}

// Get track bytes from URL
func (c *MockConnector) TrackFromURL(ctx context.Context, url string) ([]byte, error) {
	if !validUrl(url) {
		return nil, fmt.Errorf("invalid url provided")
	}

	return MockTrackContent, nil
}

// Whether track is valid
func (c *MockConnector) IsTrackValid(ctx context.Context, url string) bool {
	return validUrl(url)
}
