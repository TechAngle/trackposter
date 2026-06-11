package ytdlp

import (
	"errors"
	"io"
	"strings"
	"sync"

	"trackposter/internal/errs"
	"trackposter/internal/validator"
)

const (
	argsMultiplier  = 2
	staticArgsCount = 2 // for '-i' and URL
)

// CommandRequest provides a builder for yt-dlp arguments.
type CommandRequest struct {
	mu sync.RWMutex

	url string

	// args used for downloading and can select
	args map[string]string

	stdout io.Writer
	stderr io.Writer
}

// NewRequest returns built request.
func NewRequest(url string, args ...string) *CommandRequest {
	initArgs := defaultRequestArgs()
	req := CommandRequest{
		mu:     sync.RWMutex{},
		url:    url,
		args:   map[string]string{},
		stdout: nil,
		stderr: nil,
	}
	req.args = initArgs

	for _, arg := range args {
		req.AddArgument(arg, true)
	}

	return &req
}

// Stdout returns currently used writer to stdout.
func (r *CommandRequest) Stdout() io.Writer {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.stdout
}

// Stderr returns currently used writer to stderr.
func (r *CommandRequest) Stderr() io.Writer {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.stderr
}

// SetStdout updates stdout output.
func (r *CommandRequest) SetStdout(w io.Writer) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.stdout = w
}

// SetStderr updates stderr output.
func (r *CommandRequest) SetStderr(w io.Writer) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.stderr = w
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

	key, value, ok := strings.Cut(trimmed, " ")
	if !ok {
		value = ""
	}

	if strings.HasPrefix(key, "--") {
		if _, ok := r.args[key]; ok && !replace {
			return
		}
	}

	r.args[key] = value
}

// BuildArguments returns slice of arguments  for exec.CommandContext (or
// similar) with URL at the end.
//
// IMPORTANT: If failed to validate provided URL it will return an empty slice.
func (r *CommandRequest) BuildArguments() []string {
	err := r.Validate()
	if err != nil {
		return []string{}
	}

	pairs := len(r.args) * argsMultiplier
	total := pairs + staticArgsCount

	args := make([]string, 0, total)

	for k, v := range r.args {
		args = append(args, k)
		if v != "" {
			args = append(args, v)
		}
	}

	// Adding url as latest argument.
	args = append(args, "-i", r.url)

	return args
}

// Validate checks if provided URL in valid format.
func (r *CommandRequest) Validate() error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	err := validator.ValidateURL(r.url)
	if err != nil {
		return errors.Join(errs.ErrValidate, err)
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
		useFFMpegConversion: "",
		addMetadata:         "",
		embedMetadata:       "",
		embedThumbnail:      "",
		"--audio-format":    "flac",
	}

	return args
}
