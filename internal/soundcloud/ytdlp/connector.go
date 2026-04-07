// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package ytdlp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"trackposter/internal/domain"
	"trackposter/internal/utils"
)

// implements interface SoundcloudConnector
// NOTE: YtDlp connector requires yt-dlp binary in PATH to work
type YtDlpConnector struct {
	ytDlpPath      string
	ffmpegPath     string
	downloadFormat domain.AudioFormat
}

var _ domain.SoundcloudConnector = (*YtDlpConnector)(nil)

// create yt-dlp command with url and arguments
func (c *YtDlpConnector) newYtDlpCommand(ctx context.Context, url string, args []string, stdout io.Writer, stderr io.Writer) *exec.Cmd {
	args = append(args, url) // adding url as the last argument

	cmd := exec.CommandContext(ctx, c.ytDlpPath, args...)
	cmd.Stderr = stderr
	cmd.Stdout = stdout

	return cmd
}

// Get track metadata from URL
func (c *YtDlpConnector) TrackMetadataFromURL(ctx context.Context, url string) (metadata *domain.TrackMetadata, err error) {
	buffer := bytes.Buffer{}

	args := []string{
		"-j",   // json metadata
		"-o -", // write directly to stdout
	}

	cmd := c.newYtDlpCommand(ctx, url, args, &buffer, nil)
	if err = cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run yt-dlp command (arguments: %v): %v", args, err)
	}

	var response domain.YtDlpMetadataResponse
	err = json.Unmarshal(buffer.Bytes(), &response)

	description, ok := response.Description.(string)
	if !ok {
		description = ""
	}

	metadata = &domain.TrackMetadata{
		ID:             response.ID,
		Title:          response.Title,
		Uploader:       response.Uploader,
		Description:    description,
		ThumbnailURL:   response.Thumbnail,
		AudioExtension: response.AudioExt,
		FileName:       response.Filename,
		Duration:       response.Duration,
		FileSize:       response.FilesizeApprox,
		ReleaseDate:    response.Timestamp,
		URL:            url,
	}

	return
}

// Get track bytes from URL.
// Uses format that was set in options.
func (c *YtDlpConnector) TrackFromURL(ctx context.Context, url string) ([]byte, error) {
	buffer := bytes.Buffer{}
	args := []string{
		"-t", string(c.downloadFormat), // format
		"-o", "-", // write directly to buffer
	}

	cmd := c.newYtDlpCommand(ctx, url, args, &buffer, nil)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run yt-dlp command(arguments: %v): %v", args, err)
	}

	return buffer.Bytes(), nil
}

// checks if track exists using yt-dlp command
func (c *YtDlpConnector) trackValid(ctx context.Context, url string) bool {
	args := []string{
		"--simulate",
		"--quiet",
		"--no-warnings",
	}

	cmd := c.newYtDlpCommand(ctx, url, args, nil, nil)
	// if url is not found - yt-dlp returns 404 and error exit code
	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

// Whether track is valid
func (c *YtDlpConnector) IsTrackValid(ctx context.Context, url string) bool {
	if !utils.IsSoundcloudURL(url) {
		return false
	}

	return c.trackValid(ctx, url)
}

// NewConnector creates new soundcloud connector based on yt-dlp.
func NewConnector(options YtDlpConnectorOptions) (*YtDlpConnector, error) {
	if options.YtDlpPath == "" || options.FFMpegPath == "" {
		return nil, fmt.Errorf(
			"missing paths: yt-dlp='%s' ffmpeg='%s'",
			options.YtDlpPath,
			options.FFMpegPath,
		)
	}

	return &YtDlpConnector{
		ytDlpPath:      options.YtDlpPath,
		ffmpegPath:     options.FFMpegPath,
		downloadFormat: options.format,
	}, nil
}
