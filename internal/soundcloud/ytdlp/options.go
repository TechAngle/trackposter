// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package ytdlp

import (
	"errors"
	"os/exec"

	"trackposter/internal/domain"
)

// ConnectorOptions represents options for initializing yt-dlp connector.
type ConnectorOptions struct {
	// YtDlpPath is an absolute path to yt-dlp executable.
	YtDlpPath string

	// FFMpegPath is an absolute path to ffmpeg executable.
	FFMpegPath string

	// Format is audio type that will be used for downloading.
	Format domain.AudioFormat
}

// DefaultOptions returns default options for connector.
// Tries automatically find yt-dlp and ffmpeg in system path, but it can return an error
// if one of them cannot be accessed.
func DefaultOptions() (options ConnectorOptions, err error) {
	ffmpegPath, err := findExec("ffmpeg")
	if err != nil {
		return ConnectorOptions{}, err
	}

	ytDlpPath, err := findExec("yt-dlp")
	if err != nil {
		return ConnectorOptions{}, err
	}

	return ConnectorOptions{
		FFMpegPath: ffmpegPath,
		YtDlpPath:  ytDlpPath,
		Format:     domain.FLAC,
	}, nil
}

// findExec tries to find an absolute path to an executable in system path.
//
// Can return any exec.LookPath error if it was not found.
func findExec(execName string) (string, error) {
	path, err := exec.LookPath(execName)
	if err != nil {
		return "", errors.Join(domain.ErrLookPath, err)
	}

	return path, nil
}
