// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package ytdlp

import (
	"fmt"
	"trackposter/internal/utils"
)

// Options for connector
type YtDlpConnectorOptions struct {
	// Path to yt-dlp
	YtDlpPath string

	// Path to ffmpeg
	FFMpegPath string

	// Skip yt-dlp updates check
	SkipUpdates bool
}

// Checks if options yt-dlp path is valid and returns it.
// Returns error if it is not.
func (o *YtDlpConnectorOptions) YtDlp() (string, error) {
	if o.YtDlpPath == "" {
		return "", fmt.Errorf("yt-dlp path is empty")
	}

	if !utils.IsPathValid(o.YtDlpPath) {
		return "", fmt.Errorf("yt-dlp path is invalid")
	}

	return o.YtDlpPath, nil
}

// Checks if options ffmpeg path is valid and returns it.
// Returns error if it is not.
func (o *YtDlpConnectorOptions) FFMpeg() (string, error) {
	if !utils.IsPathValid(o.FFMpegPath) {
		return "", fmt.Errorf("yt-dlp path is invalid")
	}

	return o.YtDlpPath, nil
}
