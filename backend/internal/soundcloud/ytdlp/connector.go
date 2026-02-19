// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package ytdlp

import (
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
	types "trackposter/internal/soundcloud/types"
)

// implements interface SoundcloudClient in models
// NOTE: YtDlp connector requires python installed or yt-dlp binary in PATH to work
type YtDlpConnector struct {
	// Path to yt-dlp
	ytDlpPath string

	// Path to ffmpeg
	ffmpegPath string
}

// Check for yt-dlp updates
func (c *YtDlpConnector) checkForUpdates(ctx context.Context) error {
	args := []string{
		"--update-to master", // update to master
	}

	cmd := exec.CommandContext(ctx, c.ytDlpPath, args...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cannot update yt-dlp to the latest version: %v", err)
	}

	return nil
}

// Apply options
func (c *YtDlpConnector) applyOptions(options *YtDlpConnectorOptions) error {
	// apply yt-dlp path
	ytDlpPath, err := options.YtDlp()
	if err != nil {
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

	// apply ffmpeg path
	ffmpegPath, err := options.FFMpeg()
	if err != nil {
		// trying to get ffmpeg from system
		ffmpegPath, err = ffmpegCommad()
		if err != nil {
			return fmt.Errorf("failed to locate ffmpeg: %v", err)
		}

		if options.FFMpegPath != "" {
			log.Println("WARNING! Options ffmpeg path was not found, but it was found in system, so module will use it.")
		}
	}

	c.ytDlpPath = ytDlpPath
	c.ffmpegPath = ffmpegPath

	return nil
}

func (c *YtDlpConnector) createYtDlpCommand(ctx context.Context, args []string, stdout io.Writer, stderr io.Writer) *exec.Cmd {
	cmd := exec.CommandContext(ctx, c.ytDlpPath, args...)
	cmd.Stderr = stderr
	cmd.Stdout = stdout

	return cmd
}

// Get track metadata from URL
func (c *YtDlpConnector) TrackMetadataFromURL(ctx context.Context, url string) (*types.TrackMetadata, error) {
	metadata := &types.TrackMetadata{}

	// TODO: get metadata from url using yt-dlp command

	return metadata, nil
}

// Get track bytes from URL
func (c *YtDlpConnector) TrackFromURL(ctx context.Context, url string) (*[]byte, error) {
	// TODO: get track buffer from url using yt-dlp command
	//
	return nil, nil
}

// Create new soundcloud connector based on yt-dlp.
// Context is used for stop command execution
func NewSoundcloudConnector(ctx context.Context, options YtDlpConnectorOptions) (*YtDlpConnector, error) {
	connector := &YtDlpConnector{}
	err := connector.applyOptions(&options)
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
