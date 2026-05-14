// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in
// LICENSE file.
//
// Author: https://github.com/TechAngle

package domain

import "encoding/json"

// Command argument.
const (
	UseFFMpegConversion string = "-x"
	AddMetadata         string = "--add-metadata"
	EmbedMetadata       string = "--embed-metadata"
	EmbedThumbnail      string = "--embed-thumbnail"
	WriteThumbnail      string = "--write-thumbnail"
	Simulate            string = "--simulate"
	StdoutOutput        string = "-o -"
	JSONMetadata        string = "-j"
	Quiet               string = "--quiet"
	NoWarnings          string = "--no-warnings"
)

// YtDlpMetadataResponse is structure which represents a whole response of yt-dlp command.
type YtDlpMetadataResponse struct {
	ID          string      `json:"id"`
	Uploader    string      `json:"uploader"`
	UploaderID  string      `json:"uploader_id"`
	UploaderURL string      `json:"uploader_url"`
	Timestamp   json.Number `json:"timestamp"` // Может быть float
	Title       string      `json:"title"`
	Track       string      `json:"track"`
	Description any         `json:"description"`
	Thumbnails  []struct {
		ID         string      `json:"id"`
		URL        string      `json:"url"`
		Width      json.Number `json:"width,omitempty"` 
		Height     json.Number `json:"height,omitempty"`
		Resolution string      `json:"resolution,omitempty"`
		Preference json.Number `json:"preference,omitempty"`
	} `json:"thumbnails"`
	Duration     json.Number `json:"duration"`
	WebpageURL   string      `json:"webpage_url"`
	License      string      `json:"license"`
	ViewCount    json.Number `json:"view_count"`
	LikeCount    json.Number `json:"like_count"`
	CommentCount json.Number `json:"comment_count"`
	RepostCount  json.Number `json:"repost_count"`
	Genres       any         `json:"genres"`
	Tags         []string    `json:"tags"`
	Artists      any         `json:"artists"`
	Formats      []struct {
		FormatID       string      `json:"format_id"`
		URL            string      `json:"url"`
		Ext            string      `json:"ext"`
		Acodec         string      `json:"acodec"`
		Vcodec         string      `json:"vcodec"`
		Abr            json.Number `json:"abr"`
		Protocol       string      `json:"protocol"`
		Container      any         `json:"container"`
		Quality        json.Number `json:"quality"`
		FormatNote     any         `json:"format_note"`
		Preference     any         `json:"preference"`
		AudioExt       string      `json:"audio_ext"`
		VideoExt       string      `json:"video_ext"`
		Vbr            json.Number `json:"vbr"`
		Tbr            json.Number `json:"tbr"`
		Resolution     string      `json:"resolution"`
		AspectRatio    any         `json:"aspect_ratio"`
		FilesizeApprox json.Number `json:"filesize_approx"`
		HTTPHeaders    struct {
			UserAgent      string `json:"User-Agent"`
			Accept         string `json:"Accept"`
			AcceptLanguage string `json:"Accept-Language"`
			SecFetchMode   string `json:"Sec-Fetch-Mode"`
		} `json:"http_headers"`
		Format string `json:"format"`
	} `json:"formats"`
	OriginalURL        string      `json:"original_url"`
	WebpageURLBasename string      `json:"webpage_url_basename"`
	WebpageURLDomain   string      `json:"webpage_url_domain"`
	Extractor          string      `json:"extractor"`
	ExtractorKey       string      `json:"extractor_key"`
	Playlist           any         `json:"playlist"`
	PlaylistIndex      any         `json:"playlist_index"`
	Thumbnail          string      `json:"thumbnail"`
	DisplayID          string      `json:"display_id"`
	Fulltitle          string      `json:"fulltitle"`
	DurationString     string      `json:"duration_string"`
	UploadDate         string      `json:"upload_date"`
	ReleaseYear        any         `json:"release_year"`
	RequestedSubtitles any         `json:"requested_subtitles"`
	HasDrm             any         `json:"_has_drm"`
	Epoch              json.Number `json:"epoch"`
	FormatID           string      `json:"format_id"`
	URL                string      `json:"url"`
	Ext                string      `json:"ext"`
	Acodec             string      `json:"acodec"`
	Vcodec             string      `json:"vcodec"`
	Abr                json.Number `json:"abr"`
	Protocol           string      `json:"protocol"`
	Container          string      `json:"container"`
	Quality            json.Number `json:"quality"`
	FormatNote         any         `json:"format_note"`
	Preference         any         `json:"preference"`
	AudioExt           string      `json:"audio_ext"`
	VideoExt           string      `json:"video_ext"`
	Vbr                json.Number `json:"vbr"`
	Tbr                json.Number `json:"tbr"`
	Resolution         string      `json:"resolution"`
	AspectRatio        any         `json:"aspect_ratio"`
	FileSizeApprox     json.Number `json:"filesize_approx"`
	HTTPHeaders        struct {
		UserAgent      string `json:"User-Agent"`
		Accept         string `json:"Accept"`
		AcceptLanguage string `json:"Accept-Language"`
		SecFetchMode   string `json:"Sec-Fetch-Mode"`
	} `json:"http_headers"`
	Format   string `json:"format"`
	Filename string `json:"filename"`
	Version  struct {
		Version        string `json:"version"`
		CurrentGitHead any    `json:"current_git_head"`
		ReleaseGitHead string `json:"release_git_head"`
		Repository     string `json:"repository"`
	} `json:"_version"`
}
