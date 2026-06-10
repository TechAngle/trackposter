package model

// TrackRecord is a record representation for line.
type TrackRecord struct {
	Track *Track
	ID    string
}

// Track provides simplified structure for tracks storage and management.
type Track struct {
	Title    string   `json:"trackTitle,omitempty"`
	Author   string   `json:"trackAuthor,omitempty"`
	Duration *float64 `json:"trackDuration,omitempty"`
	URL      string   `json:"trackUrl,omitempty"`
}

// TrackMetadata is metadata structure for yt-dlp response.
type TrackMetadata struct {
	ID             string  `json:"id"`
	Title          string  `json:"title"`
	Uploader       string  `json:"uploader"`
	Description    string  `json:"description"`
	UploaderURL    string  `json:"uploaderUrl"`
	ThumbnailURL   string  `json:"thumbnailUrl"`
	AudioExtension string  `json:"audioExt"`
	FileName       string  `json:"fileName"`
	Duration       float64 `json:"duration"`
	FileSize       float64 `json:"fileSize"`
	ReleaseDate    float64 `json:"releaseDate"`
	URL            string
}

// AsTrack returns only important information (such as title, duration, author
// and url)
// of the track.
func (t *TrackMetadata) AsTrack() *Track {
	return &Track{
		Title:    t.Title,
		Duration: &t.Duration,
		Author:   t.Uploader,
		URL:      t.URL,
	}
}
