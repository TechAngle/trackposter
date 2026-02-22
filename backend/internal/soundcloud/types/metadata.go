// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package types

// Track metadata for yt-dlp
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
	FileSize       int     `json:"fileSize"`
	ReleaseDate    int     `json:"releaseDate"`
}

// yt-dlp response command
type YtDlpMetadataResponse struct {
	ID          string      `json:"id"`
	Uploader    string      `json:"uploader"`
	UploaderID  string      `json:"uploader_id"`
	UploaderURL string      `json:"uploader_url"`
	Timestamp   int         `json:"timestamp"`
	Title       string      `json:"title"`
	Track       string      `json:"track"`
	Description interface{} `json:"description"`
	Thumbnails  []struct {
		ID         string `json:"id"`
		URL        string `json:"url"`
		Width      int    `json:"width,omitempty"`
		Height     int    `json:"height,omitempty"`
		Resolution string `json:"resolution,omitempty"`
		Preference int    `json:"preference,omitempty"`
	} `json:"thumbnails"`
	Duration     float64     `json:"duration"`
	WebpageURL   string      `json:"webpage_url"`
	License      string      `json:"license"`
	ViewCount    int         `json:"view_count"`
	LikeCount    int         `json:"like_count"`
	CommentCount int         `json:"comment_count"`
	RepostCount  int         `json:"repost_count"`
	Genres       interface{} `json:"genres"`
	Tags         []string    `json:"tags"`
	Artists      interface{} `json:"artists"`
	Formats      []struct {
		FormatID       string      `json:"format_id"`
		URL            string      `json:"url"`
		Ext            string      `json:"ext"`
		Acodec         string      `json:"acodec"`
		Vcodec         string      `json:"vcodec"`
		Abr            int         `json:"abr"`
		Protocol       string      `json:"protocol"`
		Container      interface{} `json:"container"`
		Quality        int         `json:"quality"`
		FormatNote     interface{} `json:"format_note"`
		Preference     interface{} `json:"preference"`
		AudioExt       string      `json:"audio_ext"`
		VideoExt       string      `json:"video_ext"`
		Vbr            int         `json:"vbr"`
		Tbr            int         `json:"tbr"`
		Resolution     string      `json:"resolution"`
		AspectRatio    interface{} `json:"aspect_ratio"`
		FilesizeApprox int         `json:"filesize_approx"`
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
	Playlist           interface{} `json:"playlist"`
	PlaylistIndex      interface{} `json:"playlist_index"`
	Thumbnail          string      `json:"thumbnail"`
	DisplayID          string      `json:"display_id"`
	Fulltitle          string      `json:"fulltitle"`
	DurationString     string      `json:"duration_string"`
	UploadDate         string      `json:"upload_date"`
	ReleaseYear        interface{} `json:"release_year"`
	RequestedSubtitles interface{} `json:"requested_subtitles"`
	HasDrm             interface{} `json:"_has_drm"`
	Epoch              int         `json:"epoch"`
	FormatID           string      `json:"format_id"`
	URL                string      `json:"url"`
	Ext                string      `json:"ext"`
	Acodec             string      `json:"acodec"`
	Vcodec             string      `json:"vcodec"`
	Abr                int         `json:"abr"`
	Protocol           string      `json:"protocol"`
	Container          string      `json:"container"`
	Quality            int         `json:"quality"`
	FormatNote         interface{} `json:"format_note"`
	Preference         interface{} `json:"preference"`
	AudioExt           string      `json:"audio_ext"`
	VideoExt           string      `json:"video_ext"`
	Vbr                int         `json:"vbr"`
	Tbr                int         `json:"tbr"`
	Resolution         string      `json:"resolution"`
	AspectRatio        interface{} `json:"aspect_ratio"`
	FilesizeApprox     int         `json:"filesize_approx"`
	HTTPHeaders        struct {
		UserAgent      string `json:"User-Agent"`
		Accept         string `json:"Accept"`
		AcceptLanguage string `json:"Accept-Language"`
		SecFetchMode   string `json:"Sec-Fetch-Mode"`
	} `json:"http_headers"`
	Format   string `json:"format"`
	Filename string `json:"filename"`
	Version  struct {
		Version        string      `json:"version"`
		CurrentGitHead interface{} `json:"current_git_head"`
		ReleaseGitHead string      `json:"release_git_head"`
		Repository     string      `json:"repository"`
	} `json:"_version"`
}
