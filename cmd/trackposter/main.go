// Package main defines an entry point for app.
package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"trackposter/internal/connector/ytdlp"
	"trackposter/internal/logger"
	"trackposter/internal/repository"
	"trackposter/internal/telegram"
)

// Init error.
var (
	errDefaultOptions = errors.New("default options err")
	errLoadEnv        = errors.New("load .env err")
	errNewConnector   = errors.New("new connector")
)

// Parse error.
var (
	errInvalidEnvID = errors.New("parse int64 err")
)

type envConfig struct {
	allowedIDs []int64
	botToken   string
}

// stringToIDList returns parsed IDs list in int64 type.
// Returns empty slice if string is empty.
//
// If some ID cannot be parsed it returns an error.
func stringToIDList(str string) ([]int64, error) {
	allowedIDs := []int64{}

	if strings.TrimSpace(str) == "" {
		return []int64{}, nil
	}

	parts := strings.SplitSeq(str, ", ")
	for part := range parts {
		i, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return nil, errors.Join(errInvalidEnvID, err)
		}

		allowedIDs = append(allowedIDs, i)
	}

	return allowedIDs, nil
}

// loadEnvConfig loads environment variables and returns config with filled
// values.
//
// Also returns an error if failed to load .env or parse allowed IDs list from
// it.
func loadEnvConfig() (envConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return envConfig{}, errors.Join(errLoadEnv, err)
	}

	// parsing allowed IDs string
	allowedID := os.Getenv("ALLOWED_ID")

	idList, err := stringToIDList(allowedID)
	if err != nil {
		return envConfig{}, err
	}

	return envConfig{
		allowedIDs: idList,
		botToken:   os.Getenv("TOKEN"),
	}, nil
}

// initConnector creates a new soundcloud connector with default options.
func initConnector() (*ytdlp.Connector, error) {
	options, err := ytdlp.DefaultOptions()
	if err != nil {
		return nil, errors.Join(errDefaultOptions, err)
	}

	connector, err := ytdlp.NewConnector(options)
	if err != nil {
		return nil, errors.Join(errNewConnector, err)
	}

	return connector, nil
}

func initRepository() *repository.MemoryQueue {
	return repository.NewMemoryQueue()
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := logger.NewLogger()

	config, err := loadEnvConfig()
	if err != nil {
		logger.ErrorContext(ctx, "load config err", "error", err)

		return
	}

	connector, err := initConnector()
	if err != nil {
		logger.ErrorContext(ctx, "connector init err", "error", err)

		return
	}

	repository := initRepository()

	client, err := telegram.NewClient(telegram.BotOptions{
		Connector:  connector,
		Repository: repository,
		Logger:     &logger,
		AllowedIDs: config.allowedIDs,
		APIToken:   config.botToken,
	})
	if err != nil {
		logger.ErrorContext(ctx, "bot init err", "error", err)

		return
	}
	defer client.Stop(ctx)

	err = client.Start(ctx)
	if err != nil {
		logger.ErrorContext(ctx, "bot start err", "error", err)
	}
}
