package mock

import "trackposter/internal/model"

const (
	mockDuration    = 214.5
	mockFileSize    = 8589934
	mockReleaseDate = 1735689600
)

var (
	// MockTrack is a random track metadata for tests.
	MockTrack model.TrackMetadata = model.TrackMetadata{
		ID:             "666-dev-null",
		Title:          "Never Gonna Give Your DOM Up",
		Uploader:       "Rick Rolling Gopher",
		Description:    "A song about the pain of writing MutationObservers on Windows 11.",
		UploaderURL:    "https://soundcloud.com/gopher-hater",
		ThumbnailURL:   "https://i1.sndcdn.com/artworks-mock-large.jpg",
		AudioExtension: "mp3",
		FileName:       "rick_rolling_gopher_never_gonna_give_your_dom_up.mp3",
		Duration:       mockDuration,
		FileSize:       mockFileSize,
		ReleaseDate:    mockReleaseDate,
		URL:            "https://soundcloud.com/gopther-hater/i-love-c",
	}

	// MockTrackContent is an example content that might be returned by real
	// tool.
	MockTrackContent = []byte{1, 2, 3, 4, 5}
)
