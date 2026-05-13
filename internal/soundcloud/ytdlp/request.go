// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package ytdlp

import (
	"io"
	"net/url"
	"strings"
	"sync"

	"trackposter/internal/domain"
)

type CommandRequest struct {
	mu sync.RWMutex

	URL string

	// Args used for downloading and can select
	Args map[string]string

	Stdout io.Writer
	Stderr io.Writer
}

// NewRequest returns built request.
func NewRequest(url string, args ...string) *CommandRequest {
	initArgs := defaultRequestArgs()
	r := CommandRequest{
		URL:    url,
		Args:   map[string]string{},
		Stdout: nil,
		Stderr: nil,
	}
	r.Args = initArgs

	for _, arg := range args {
		r.AddArgument(arg, true)
	}

	return &r
}

// SetStdout updates stdout output.
func (r *CommandRequest) SetStdout(w io.Writer) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Stdout = w
}

func (r *CommandRequest) SetStderr(w io.Writer) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Stderr = w
}

// AddArgument updates arguments list.
//
// If such argument already was set it ignores new one. Also, if more than 1
// argument was provided it will ignore them.
func (r *CommandRequest) AddArgument(arg string, replace bool) {
	trimmed := strings.TrimSpace(arg)
	if trimmed == "" {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	k, v, ok := strings.Cut(trimmed, " ")
	if !ok {
		v = ""
	}

	if strings.HasPrefix(k, "--") {
		if _, ok := r.Args[k]; ok && !replace {
			return
		}
	}

	r.Args[k] = v
}

// BuildArguments returns slice of arguments  for exec.CommandContext (or similar) with URL at the end.
//
// IMPORTANT: If failed to validate provided URL it will return an empty slice.
func (r *CommandRequest) BuildArguments() []string {
	if err := r.Validate(); err != nil {
		return []string{}
	}

	// `+ 2` for '-i' and URL parts.
	total := len(r.Args)*2 + 2

	args := make([]string, 0, total)

	for k, v := range r.Args {
		args = append(args, k)
		if v != "" {
			args = append(args, v)
		}
	}

	// Adding url as latest argument.
	args = append(args, "-i", r.URL)

	return args
}

// Validate checks if provided URL in valid format.
func (r *CommandRequest) Validate() error {
	if _, err := url.ParseRequestURI(r.URL); err != nil {
		return domain.ErrInvalidURL
	}

	return nil
}

// defaultRequestArgs returns default options used for downloading files:
//   - ffmpeg conversion,
//   - included/embedded metadata,
//   - embedded thumbnail,
//     -`.flac` audio format.
func defaultRequestArgs() map[string]string {
	args := map[string]string{
		domain.UseFFMpegConversion: "",
		domain.AddMetadata:         "",
		domain.EmbedMetadata:       "",
		domain.EmbedThumbnail:      "",
		"--audio-format":           "flac",
	}
	return args
}
