// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package types

import "context"

type SoundcloudConnector interface {
	// Get track metadata
	TrackMetadataFromURL(ctx context.Context, url string) (*TrackMetadata, error)

	// Get track bytes from URL
	TrackFromURL(ctx context.Context, url string) ([]byte, error)

	// Whether track is valid
	IsTrackValid(ctx context.Context, url string) bool
}
