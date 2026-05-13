// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package domain

import "errors"

// Global error.
var (
	ErrDefaultOptions = errors.New("default options err")
	ErrNewConnector   = errors.New("new connector")
	ErrUUID           = errors.New("uuid err")
	ErrUnmarshal      = errors.New("unmarshal err")
	ErrInvalidEnvID   = errors.New("parse int64 err")
	ErrLoadEnv        = errors.New("load .env err")
)

// Connector/repository error.
var (
	ErrInvalidURL    = errors.New("invalid url")
	ErrYtDlp         = errors.New("yt-dlp command err")
	ErrTrackNotFound = errors.New("track not found")
	ErrMissingYtDlp  = errors.New("missing yt-dlp path")
	ErrMissingFFMpeg = errors.New("missing ffmpeg path")
	ErrHandlerInit   = errors.New("handler init err")
)

// Telegram bot error.
var (
	ErrQueueProcess = errors.New("queue process err")
	ErrNilClient    = errors.New("client is nil")
	ErrClientInit   = errors.New("bot client init")
	ErrInvalidToken = errors.New("invalid token")
	ErrEmptyQueue   = errors.New("empty queue")
)
