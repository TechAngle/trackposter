// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package ytdlp

import (
	"fmt"
	"os/exec"
	"trackposter/internal/domain"
)

// YtDlpConnectorOptions represents options for initializing yt-dlp connector
type YtDlpConnectorOptions struct {
	// YtDlpPath is an absolute path to yt-dlp executable.
	YtDlpPath string

	// FFMpegPath is an absolute path to ffmpeg executable.
	FFMpegPath string

	// FIX: Hardcoded ATM, will be accessible soon.
	format domain.AudioFormat
}

// DefaultOptions returns default options for connector.
// Tries automatically find yt-dlp and ffmpeg in system path, but it can return an error
// if one of them cannot be accessed.
func DefaultOptions() (options YtDlpConnectorOptions, err error) {
	ffmpegPath, err := findExec("ffmpeg")
	if err != nil {
		return YtDlpConnectorOptions{}, fmt.Errorf("cannot get ffmpeg command: %w", err)
	}

	ytDlpPath, err := findExec("yt-dlp")
	if err != nil {
		return YtDlpConnectorOptions{}, fmt.Errorf("cannot get ytdlp command: %w", err)
	}

	return YtDlpConnectorOptions{
		FFMpegPath: ffmpegPath,
		YtDlpPath:  ytDlpPath,
		format:     domain.MP3,
	}, nil
}

// findExec tries to find an absolute path to an executable in system path.
//
// Can return any exec.LookPath error if it was not found.
func findExec(execName string) (string, error) {
	path, err := exec.LookPath(execName)
	if err != nil {
		return "", err
	}

	return path, nil
}
