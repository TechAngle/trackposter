package mock

import (
	"context"
	"errors"
	"sync"

	"trackposter/internal/connector"
	"trackposter/internal/errs"
	"trackposter/internal/model"
	"trackposter/internal/validator"
)

// Connector provides predefined values from constants.
type Connector struct {
	mu sync.RWMutex

	format model.AudioFormat
}

var _ connector.Connector = (*Connector)(nil)

// TrackMetadataFromURL retrieves track metadata from URL.
func (c *Connector) TrackMetadataFromURL(
	ctx context.Context,
	url string,
) (*model.TrackMetadata, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.Join(errs.ErrBadContext, err)
	}

	err = validator.ValidateURL(url)
	if err != nil {
		return nil, errors.Join(errs.ErrValidate, err)
	}

	// cloning and adding url
	mockTrack := MockTrack
	mockTrack.URL = url

	return &mockTrack, nil
}

// TrackFromURL retrieves track bytes from URL.
// Uses format that was set in options.
func (c *Connector) TrackFromURL(
	ctx context.Context,
	url string,
) ([]byte, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.Join(errs.ErrBadContext, err)
	}

	err = validator.ValidateURL(url)
	if err != nil {
		return nil, errors.Join(errs.ErrValidate, err)
	}

	return MockTrackContent, nil
}

// IsTrackValid checks if track is valid.
func (c *Connector) IsTrackValid(ctx context.Context, url string) bool {
	err := ctx.Err()
	if err != nil {
		return false
	}

	err = validator.ValidateURL(url)

	return err != nil
}

// SetFormat updates audio format for downloading.
func (c *Connector) SetFormat(format model.AudioFormat) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.format = format
}

// AudioFormat returns currently used audio format for downloading.
func (c *Connector) AudioFormat() model.AudioFormat {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.format
}
