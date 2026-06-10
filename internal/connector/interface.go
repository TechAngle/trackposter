// Package connector provides an interface for creating abstraction over
// different tools.
package connector

import (
	"context"

	"trackposter/internal/model"
)

// Connector provides all logic for work with tracks from external sources and
// their metadata.
type Connector interface {
	// TrackMetadataFromURL retrieves track metadata from URL.
	TrackMetadataFromURL(
		ctx context.Context,
		url string,
	) (*model.TrackMetadata, error)

	// TrackFromURL retrieves track bytes from URL.
	// Uses format that was set in options.
	TrackFromURL(ctx context.Context, url string) ([]byte, error)

	// IsTrackValid checks if track is valid.
	IsTrackValid(ctx context.Context, url string) bool

	// SetFormat updates audio format for downloading.
	SetFormat(format model.AudioFormat)

	// AudioFormat returns currently used audio format for downloading.
	AudioFormat() model.AudioFormat
}
