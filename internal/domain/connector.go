// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in
// LICENSE file.
//
// Author: https://github.com/TechAngle

package domain

import (
	"context"
)

type Connector interface {
	// TrackMetadataFromURL retrieves track metadata from URL.
	TrackMetadataFromURL(
		ctx context.Context,
		url string,
	) (*TrackMetadata, error)

	// TrackFromURL retrieves track bytes from URL.
	// Uses format that was set in options.
	TrackFromURL(ctx context.Context, url string) ([]byte, error)

	// IsTrackValid checks if track is valid.
	IsTrackValid(ctx context.Context, url string) bool

	// SetFormat updates audio format for downloading.
	SetFormat(format AudioFormat)

	// AudioFormat returns currently used audio format for downloading.
	AudioFormat() AudioFormat
}

type AudioFormat string

// Audio format.
const (
	MP3  AudioFormat = "mp3"
	M4A  AudioFormat = "m4a"
	WAV  AudioFormat = "wav"
	FLAC AudioFormat = "flac"
	OPUS AudioFormat = "opus"
	AAC  AudioFormat = "aac"
)

// Formats returns slice of available audio formats.
func Formats() []AudioFormat {
	return []AudioFormat{
		MP3,
		M4A,
		WAV,
		FLAC,
		OPUS,
		AAC,
	}
}
