// Package logger provides logger logic.
package logger

import (
	"log/slog"
	"os"
)

// NewLogger returns predefined logger with output to stdout.
func NewLogger() slog.Logger {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, nil),
	)

	return *logger
}
