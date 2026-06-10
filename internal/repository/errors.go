package repository

import "errors"

// General repository errors that could be used for wrapping across all app.
var (
	ErrInternal = errors.New("repository internal err")
)

// Track errors.
var (
	ErrTrackNotFound = errors.New("track not found")
)
