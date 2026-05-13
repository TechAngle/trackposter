// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package mock

import (
	"context"
	"sync"

	"trackposter/internal/domain"
)

// Implements SoundCloud connector.
//
// Mock connector that returns pre-defined values from config.
type Connector struct {
	mu sync.RWMutex

	format domain.AudioFormat
}

var _ domain.Connector = (*Connector)(nil)

// TrackMetadataFromURL retrieves track metadata from URL.
func (c *Connector) TrackMetadataFromURL(ctx context.Context, url string) (*domain.TrackMetadata, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := domain.ValidateURL(url); err != nil {
		return nil, err
	}

	// cloning and adding url
	mockTrack := MockTrack
	mockTrack.URL = url

	return &mockTrack, nil
}

// TrackFromURL retrieves track bytes from URL.
// Uses format that was set in options.
func (c *Connector) TrackFromURL(ctx context.Context, url string) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := domain.ValidateURL(url); err != nil {
		return nil, err
	}

	return MockTrackContent, nil
}

// IsTrackValid checks if track is valid.
func (c *Connector) IsTrackValid(ctx context.Context, url string) bool {
	if err := ctx.Err(); err != nil {
		return false
	}

	err := domain.ValidateURL(url)
	return err != nil
}

// SetFormat updates audio format for downloading.
func (c *Connector) SetFormat(format domain.AudioFormat) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.format = format
}

// AudioFormat returns currently used audio format for downloading.
func (c *Connector) AudioFormat() domain.AudioFormat {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.format
}
