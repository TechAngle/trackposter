// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in
// LICENSE file.
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
	ErrLookPath       = errors.New("look path err")
)

// Connector error.
var (
	ErrConnectorInternal = errors.New("connector err")
)

// Repository error.
var (
	ErrRepositoryInternal = errors.New("repository err")
	ErrTrackNotFound      = errors.New("track not found")
)

// Telegram bot error.
var (
	ErrQueueProcess = errors.New("queue process err")
	ErrNilClient    = errors.New("client is nil")
	ErrClientInit   = errors.New("bot client init")
	ErrInvalidToken = errors.New("invalid token")
	ErrEmptyQueue   = errors.New("empty queue")
	ErrTelegramAPI  = errors.New("telegram api err")
	ErrHandlerInit  = errors.New("handler init err")
)

// Context error.
var (
	ErrBadContext = errors.New("context err")
)

// Tool error.
var (
	ErrYtDlpCommand  = errors.New("yt-dlp command err")
	ErrMissingYtDlp  = errors.New("missing yt-dlp path")
	ErrMissingFFMpeg = errors.New("missing ffmpeg path")
)

// Validation error.
var (
	ErrValidation = errors.New("validation err")
	ErrInvalidURL = errors.New("invalid url")
)
