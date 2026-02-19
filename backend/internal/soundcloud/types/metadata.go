// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package models

// Track metadata for yt-dlp
type TrackMetadata struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Uploader       string `json:"uploader"`
	Descriptiong   string `json:"description"`
	UploaderURL    string `json:"uploaderUrl"`
	ThumbnailURL   string `json:"thumbnailUrl"`
	AudioExtension string `json:"audioExt"`
	FileName       string `json:"fileName"`
	Duration       int    `json:"duration"`
	FileSize       int    `json:"fileSize"`
	ReleaseDate    int    `json:"releaseDate"`
}
