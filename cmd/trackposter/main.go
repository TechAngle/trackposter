// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"trackposter/internal/domain"
	"trackposter/internal/soundcloud/ytdlp"
	"trackposter/internal/telegram"

	"github.com/joho/godotenv"
)

type envConfig struct {
	allowedIDs []int64
	botToken   string
}

// stringToIDList returns parsed IDs list in int64 type.
// Returns empty slice if string is empty.
//
// If some ID cannot be parsed it returns an error.
func stringToIDList(s string) ([]int64, error) {
	allowedIDs := []int64{}

	if strings.TrimSpace(s) == "" {
		return []int64{}, nil
	}

	parts := strings.SplitSeq(s, ", ")
	for part := range parts {
		i, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return nil, errors.Join(domain.ErrInvalidEnvID, err)
		}

		allowedIDs = append(allowedIDs, i)
	}

	return allowedIDs, nil
}

// loadEnvConfig loads environment variables and returns config with filled values.
//
// Also returns an error if failed to load .env or parse allowed IDs list from it.
func loadEnvConfig() (envConfig, error) {
	if err := godotenv.Load(".env"); err != nil {
		return envConfig{}, errors.Join(domain.ErrLoadEnv, err)
	}

	// parsing allowed ids string
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
func initConnector() (domain.SoundcloudConnector, error) {
	options, err := ytdlp.DefaultOptions()
	if err != nil {
		return nil, errors.Join(domain.ErrDefaultOptions, err)
	}

	connector, err := ytdlp.NewConnector(options)
	if err != nil {
		return nil, errors.Join(domain.ErrNewConnector, err)
	}

	return connector, nil
}

func main() {
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

	telegram.NewBot(telegram.BotOptions{
		Connector:  connector,
		AllowedIDs: config.allowedIDs,
		APIToken:   config.botToken,
	})
}
