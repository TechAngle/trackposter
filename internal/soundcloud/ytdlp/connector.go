// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in
// LICENSE file.
//
// Author: https://github.com/TechAngle

// Package ytdlp provides connector for work with yt-dlp tool.
package ytdlp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"os/exec"
	"sync"

	"trackposter/internal/domain"
	"trackposter/internal/pkg/pool"
)

// Connector implements interface SoundcloudConnector.
// NOTE: YtDlp connector requires yt-dlp binary in PATH to work.
type Connector struct {
	mu sync.RWMutex

	ytDlpPath      string
	ffmpegPath     string
	downloadFormat domain.AudioFormat
}

var _ domain.Connector = (*Connector)(nil)

// NewConnector creates new soundcloud connector based on yt-dlp.
func NewConnector(options ConnectorOptions) (*Connector, error) {
	if options.YtDlpPath == "" {
		return nil, domain.ErrMissingYtDlp
	}

	if options.FFMpegPath == "" {
		return nil, domain.ErrMissingFFMpeg
	}

	return &Connector{
		mu:             sync.RWMutex{},
		ytDlpPath:      options.YtDlpPath,
		ffmpegPath:     options.FFMpegPath,
		downloadFormat: options.Format,
	}, nil
}

// TrackMetadataFromURL retrieves track metadata from URL.
func (c *Connector) TrackMetadataFromURL(
	ctx context.Context,
	url string,
) (metadata *domain.TrackMetadata, err error) {
	buffer := pool.GetBuffer()
	defer pool.PutBuffer(buffer)

	req := NewRequest(
		url,
		domain.JSONMetadata,
		domain.StdoutOutput,
		domain.NoWarnings,
	)
	req.SetStdout(buffer)

	cmd := c.newCommand(ctx, req)
	err = cmd.Run()
	if err != nil {
		return nil, errors.Join(domain.ErrYtDlpCommand, err)
	}

	var response domain.YtDlpMetadataResponse

	err = json.Unmarshal(buffer.Bytes(), &response)
	if err != nil {
		log.Println(buffer.String())

		return nil, errors.Join(domain.ErrUnmarshal, err)
	}

	metadata = metadataFromResponse(&response)

	return metadata, nil
}

// TrackFromURL retrieves track bytes from URL.
// Uses format that was set in options.
func (c *Connector) TrackFromURL(
	ctx context.Context,
	url string,
) ([]byte, error) {
	// Using default bytes buffer instead of pool to avoid problems with stack.
	buffer := new(bytes.Buffer)

	req := NewRequest(
		url,
		domain.UseFFMpegConversion,
		"--audio-format", string(c.AudioFormat()),
		domain.AddMetadata,
		domain.EmbedMetadata,
		domain.WriteThumbnail,
		domain.StdoutOutput,
	)

	req.SetStdout(buffer)

	cmd := c.newCommand(ctx, req)
	err := cmd.Run()
	if err != nil {
		return nil, errors.Join(domain.ErrYtDlpCommand, err)
	}

	return buffer.Bytes(), nil
}

// IsTrackValid checks if track is valid.
func (c *Connector) IsTrackValid(ctx context.Context, url string) bool {
	err := domain.ValidateURL(url)
	if err != nil {
		return false
	}

	return c.trackValid(ctx, url)
}

// SetFormat updates audio format for downloading.
func (c *Connector) SetFormat(format domain.AudioFormat) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.downloadFormat = format
}

// AudioFormat returns currently used audio format for downloading.
func (c *Connector) AudioFormat() domain.AudioFormat {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.downloadFormat
}

func (c *Connector) newCommand(
	ctx context.Context,
	req *CommandRequest,
) *exec.Cmd {
	args := req.BuildArguments()

	cmd := exec.CommandContext(ctx, c.ytDlpPath, args...) // #nosec G204
	cmd.Stderr = req.stderr
	cmd.Stdout = req.stdout

	return cmd
}

func (c *Connector) trackValid(ctx context.Context, url string) bool {
	req := NewRequest(
		url,
		domain.Simulate,
		domain.Quiet,
		domain.NoWarnings,
	)

	cmd := c.newCommand(ctx, req)
	// if url is not found - yt-dlp returns 404 and error exit code
	return cmd.Run() == nil
}

func metadataFromResponse(
	response *domain.YtDlpMetadataResponse,
) *domain.TrackMetadata {
	description, ok := response.Description.(string)
	if !ok {
		description = ""
	}

	duration, err := response.Duration.Float64()
	if err != nil {
		duration = 0
	}

	fileSize, err := response.FileSizeApprox.Float64()
	if err != nil {
		fileSize = 0
	}

	timestamp, err := response.Timestamp.Float64()
	if err != nil {
		timestamp = 0
	}

	return &domain.TrackMetadata{
		ID:             response.ID,
		Title:          response.Title,
		Uploader:       response.Uploader,
		UploaderURL:    response.UploaderURL,
		Description:    description,
		ThumbnailURL:   response.Thumbnail,
		AudioExtension: response.AudioExt,
		FileName:       response.Filename,
		Duration:       duration,
		FileSize:       fileSize,
		ReleaseDate:    timestamp,
		URL:            response.URL,
	}
}
