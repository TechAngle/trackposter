// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package ytdlp

import (
	"fmt"
)

// Options for connector
type YtDlpConnectorOptions struct {
	// Path to yt-dlp
	YtDlpPath string

	// Path to ffmpeg
	FFMpegPath string

	// Skip yt-dlp updates check
	SkipUpdates bool

	// Download format
	format Format
}

// Get default options.
// Automatically finds yt-dlp and ffmpeg.
// If one of them not found - returns error.
func DefaultOptions() (options *YtDlpConnectorOptions, err error) {
	ffmpegPath, err := ffmpegCommad()
	if err != nil {
		return nil, fmt.Errorf("cannot get ffmpeg command: %v", err)
	}

	ytDlpPath, err := ytDlpCommand()
	if err != nil {
		return nil, fmt.Errorf("cannot get ytdlp command: %v", err)
	}

	return &YtDlpConnectorOptions{
		FFMpegPath:  ffmpegPath,
		YtDlpPath:   ytDlpPath,
		SkipUpdates: false,
		format:      MP3,
	}, nil
}

// Checks if options yt-dlp path is valid and returns it.
// Returns error if it is not.
func (o *YtDlpConnectorOptions) YtDlp() string {
	return o.YtDlpPath
}

// Checks if options ffmpeg path is valid and returns it.
// Returns error if it is not.
func (o *YtDlpConnectorOptions) FFMpeg() string {
	return o.YtDlpPath
}
