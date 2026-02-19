// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package models

type SoundcloudClient interface {
	// Get track metadata
	TrackMetadataFromURL(url string) (*TrackMetadata, error)

	// Get track bytes from URL
	TrackFromURL(url string) (*[]byte, error)
}
