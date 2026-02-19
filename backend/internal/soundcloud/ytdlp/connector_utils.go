// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package ytdlp

import (
	"fmt"
	"os/exec"
	"trackposter/internal/utils"
)

// Get actual yt-dlp command. Returns empty string and error if it was not found
func ytDlpCommand() (string, error) {
	if !utils.IsPythonInstalled() {
		return "", fmt.Errorf("install python first")
	}

	ytDlpPath, err := exec.LookPath("yt-dlp")
	if err != nil {
		return "", fmt.Errorf("yt-dlp executable not found: %v", err)
	}

	return ytDlpPath, nil
}

// Get ffmpeg command. Returns empty string and error if it was not found
func ffmpegCommad() (string, error) {
	ffmpegPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		return "", fmt.Errorf("ffmpeg not found: %v", err)
	}

	return ffmpegPath, nil
}
