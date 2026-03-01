// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package mock

import "trackposter/internal/soundcloud/types"

var (
	// Mock track metadata
	MockTrack types.TrackMetadata = types.TrackMetadata{
		ID:             "666-dev-null",
		Title:          "Never Gonna Give Your DOM Up",
		Uploader:       "Rick Rolling Gopher",
		Description:    "A song about the pain of writing MutationObservers on Windows 11.",
		UploaderURL:    "https://soundcloud.com/gopher-hater",
		ThumbnailURL:   "https://i1.sndcdn.com/artworks-mock-large.jpg",
		AudioExtension: "mp3",
		FileName:       "rick_rolling_gopher_never_gonna_give_your_dom_up.mp3",
		Duration:       214.5,
		FileSize:       8589934,
		ReleaseDate:    1735689600,
		URL:            "https://soundcloud.com/gopther-hater/i-love-c",
	}

	// Mock track content that could be returned
	MockTrackContent []byte = []byte{1, 2, 3, 4, 5}
)
