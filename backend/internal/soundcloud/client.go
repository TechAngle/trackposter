// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package soundcloud

import (
	"context"
	"fmt"
	"trackposter/internal/soundcloud/types"
	"trackposter/internal/soundcloud/ytdlp"
)

// SoundCloud client structure
type SoundcloudClient struct {
	// SoundCloud connector which provides all features
	connector types.SoundcloudConnector
}

// Get track metadata from URL.
func (c *SoundcloudClient) TrackMetadata(ctx context.Context, url string) (*types.TrackMetadata, error) {
	metadata, err := c.connector.TrackMetadataFromURL(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %v", err)
	}

	return metadata, nil
}

// Get track bytes from URL.
func (c *SoundcloudClient) Track(ctx context.Context, url string) ([]byte, error) {
	trackContent, err := c.connector.TrackFromURL(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get track data: %v", err)
	}

	return trackContent, nil
}

// Whether track is valid
func (c *SoundcloudClient) ValidTrack(ctx context.Context, url string) bool {
	valid := c.connector.IsTrackValid(ctx, url)
	return valid
}

// Creates new soundcloud client based on connector.
// TODO: Add another connector, without python tool
func NewSCClient(ctx context.Context) (*SoundcloudClient, error) {
	connectorOptions, err := ytdlp.DefaultOptions()
	if err != nil {
		return nil, fmt.Errorf("failed to get default options: %v", err)
	}

	connector, err := ytdlp.NewConnector(ctx, connectorOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to init connector: %v", err)
	}

	return &SoundcloudClient{
		connector: connector,
	}, nil
}
