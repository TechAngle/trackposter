// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package telegram

import (
	"context"
	"errors"
	"log/slog"
	"strings"

	"trackposter/internal/domain"
	"trackposter/internal/telegram/handlers/command"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Queue chan tgbotapi.Update

type BotOptions struct {
	AllowedIDs []int64
	Connector  domain.Connector
	Repository domain.Repository
	Logger     *slog.Logger
	APIToken   string
}

type Client struct {
	client         *tgbotapi.BotAPI
	commandHandler domain.TelegramHandler
	// urlHandler     domain.TelegramHandler
	allowedIDs   []int64
	messageQueue Queue
	repository   domain.Repository
	connector    domain.Connector
	logger       *slog.Logger
}

// NewClient creates new Bot structure with all defined fields.
//
// If token was found empty returns ErrInvalidToken.
// If failed to init new Telegram Bot API client returns ErrClientInit.
func NewClient(options BotOptions) (*Client, error) {
	if strings.TrimSpace(options.APIToken) == "" {
		return nil, domain.ErrInvalidToken
	}

	client, err := newClient(options.APIToken)
	if err != nil {
		return nil, errors.Join(domain.ErrClientInit, err)
	}

	commandHandler, err := command.NewHandler(command.HandlerOptions{
		Client:     client,
		Connector:  options.Connector,
		Repository: options.Repository,
		Logger:     options.Logger.With("module", "CommandHandler"),
	})
	if err != nil {
		return nil, errors.Join(domain.ErrHandlerInit, err)
	}

	return &Client{
		client:         client,
		commandHandler: commandHandler,
		connector:      options.Connector,
		repository:     options.Repository,
		allowedIDs:     options.AllowedIDs,
		logger:         options.Logger.With("module", "Client"),
		messageQueue:   newQueue(),
	}, nil
}

// Start starts polling for client with default timeout.
//
// If client was not init returns ErrNilClient.
func (c *Client) Start(ctx context.Context) error {
	if c.client == nil {
		return domain.ErrNilClient
	}

	u := tgbotapi.NewUpdate(-1)
	u.Timeout = 45

	updates := c.client.GetUpdatesChan(u)

	c.handleUpdates(ctx, updates)
	c.logger.InfoContext(ctx, "started updates handling")

	c.processQueue(ctx, &c.messageQueue)
	c.logger.InfoContext(ctx, "started queue processing")

	<-ctx.Done()

	return nil
}

// Stop closes all queues.
func (c *Client) Stop(ctx context.Context) {
	close(c.messageQueue)

	c.logger.InfoContext(ctx, "queues were stopped")
}

func newClient(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, errors.Join(domain.ErrTelegramAPI, err)
	}

	return bot, nil
}
