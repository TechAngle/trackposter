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
	"log"
	"os/exec"
	types "trackposter/internal/soundcloud/types"
	"trackposter/internal/soundcloud/utils"
)

// implements interface SoundcloudConnector in models
// NOTE: YtDlp connector requires yt-dlp binary in PATH to work
type YtDlpConnector struct {
	// Path to yt-dlp
	ytDlpPath string

	// Path to ffmpeg
	ffmpegPath string

	// download format
	downloadFormat Format
}

// Check for yt-dlp updates
func (c *YtDlpConnector) checkForUpdates(ctx context.Context) error {
	args := []string{
		"-U", // update to master
	}

	cmd := exec.CommandContext(ctx, c.ytDlpPath, args...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cannot update yt-dlp to the latest version: %v", err)
	}

	return nil
}

// apply all download options
func (c *YtDlpConnector) applyDownloadOptions(options *YtDlpConnectorOptions) {
	c.downloadFormat = options.format
}

// apply all yt-dlp options
func (c *YtDlpConnector) applyYtDlpOption(options *YtDlpConnectorOptions) (err error) {
	// apply yt-dlp path
	ytDlpPath := options.YtDlp()
	if ytDlpPath == "" {
		// trying to get yt-dlp from system
		ytDlpPath, err = ytDlpCommand()
		if err != nil {
			return fmt.Errorf("failed to locate yt-dlp: %v", err)
		}

		// at least something was set in path
		if options.YtDlpPath != "" {
			log.Println("WARNING! Options yt-dlp path was not found, but it was found in system, so module will use it.")
		}
	}

	c.ytDlpPath = ytDlpPath

	return
}

// apply all ffmpeg options
func (c *YtDlpConnector) applyFFmpegOption(options *YtDlpConnectorOptions) (err error) {
	// apply ffmpeg path
	ffmpegPath := options.FFMpeg()
	if ffmpegPath == "" {
		// trying to get ffmpeg from system
		ffmpegPath, err = ffmpegCommad()
		if err != nil {
			return fmt.Errorf("failed to locate ffmpeg: %v", err)
		}

		// if at least something was set in path
		if options.FFMpegPath != "" {
			log.Println("WARNING! Options ffmpeg path was not found, but it was found in system, so module will use it.")
		}
	}

	c.ffmpegPath = ffmpegPath

	return
}

// Apply options
func (c *YtDlpConnector) applyOptions(options *YtDlpConnectorOptions) (err error) {
	err = c.applyYtDlpOption(options)
	if err != nil {
		return fmt.Errorf("failed to apply yt-dlp options: %v", err)
	}

	err = c.applyFFmpegOption(options)
	if err != nil {
		return fmt.Errorf("failed to apply ffmpeg options: %v", err)
	}

	return
}

// create yt-dlp command with url and arguments
func (c *YtDlpConnector) newYtDlpCommand(ctx context.Context, url string, args []string, stdout io.Writer, stderr io.Writer) *exec.Cmd {
	args = append(args, url) // adding url as the last argument

	cmd := exec.CommandContext(ctx, c.ytDlpPath, args...)
	cmd.Stderr = stderr
	cmd.Stdout = stdout

	return cmd
}

// Get track metadata from URL
func (c *YtDlpConnector) TrackMetadataFromURL(ctx context.Context, url string) (metadata *types.TrackMetadata, err error) {
	buffer := bytes.Buffer{}

	args := []string{
		"-j",   // json metadata
		"-o -", // write directly to stdout
	}

	cmd := c.newYtDlpCommand(ctx, url, args, &buffer, nil)
	if err = cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run yt-dlp command (arguments: %v): %v", args, err)
	}

	var response types.YtDlpMetadataResponse
	err = json.Unmarshal(buffer.Bytes(), &response)

	description, ok := response.Description.(string)
	if !ok {
		description = ""
	}

	metadata = &types.TrackMetadata{
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

// Create new soundcloud connector based on yt-dlp.
// Context is used for stop command execution
func NewConnector(ctx context.Context, options *YtDlpConnectorOptions) (*YtDlpConnector, error) {
	connector := &YtDlpConnector{}
	err := connector.applyOptions(options)
	if err != nil {
		return nil, fmt.Errorf("failed to apply options: %v", err)
	}

	if !options.SkipUpdates {
		err := connector.checkForUpdates(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to check yt-dlp updates: %v", err)
		}
	}

	return connector, nil
}
