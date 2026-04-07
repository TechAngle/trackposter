// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package telegram

import (
	"fmt"
	"slices"
	"strings"
	"trackposter/internal/config"
	"trackposter/internal/domain"
	"trackposter/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const ()

type Queue chan tgbotapi.Update

type Bot struct {
	client       *tgbotapi.BotAPI
	allowedIDs   []int64
	messageQueue Queue
	connector    domain.SoundcloudConnector
}

type BotOptions struct {
	AllowedIDs []int64
	Connector  domain.SoundcloudConnector
	APIToken   string
}

// Start starts polling for client with default timeout.
func (b *Bot) Start() error {
	if b.client == nil {
		return fmt.Errorf("client was not initialized")
	}

	u := tgbotapi.NewUpdate(-1)
	u.Timeout = 45
	updates := b.client.GetUpdatesChan(u)

	b.handleUpdates(updates)

	if err := b.handleMessages(newQueue()); err != nil {
		return fmt.Errorf("failed to handle messsages: %w", err)
	}

	return nil
}

// Stop closes all queues
func (b *Bot) Stop() {
	close(b.messageQueue)
}

// NewBot creates new Bot structure with all defined fields
func NewBot(options BotOptions) (*Bot, error) {
	if strings.TrimSpace(options.APIToken) == "" {
		return nil, fmt.Errorf("invalid token")
	}

	client, err := newClient(options.APIToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot client: %w", err)
	}

	return &Bot{
		client:    client,
		connector: options.Connector,
	}, nil
}

// newQueue creates a queue channel for messages.
func newQueue() Queue {
	return make(Queue, config.MaxQueueSize)
}

// newClient creates new bot api client.
//
// Can return an error if failed to create bot api for some reason.
func newClient(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create new bot api: %w", err)
	}

	return bot, err
}

// validMessage checks if message is not nil and contains text or caption.
func validMessage(msg *tgbotapi.Message) bool {
	return msg != nil && (msg.Text != "" || msg.Caption != "")
}

// messageText returns text or caption depends on what message contains.
func messageText(msg *tgbotapi.Message) string {
	if msg == nil {
		return ""
	}

	if msg.Caption != "" {
		return msg.Caption
	}

	return msg.Text
}

// addUpdate validates update and adds if it possible.
//
// # Requirements for update:
//   - Update Message is found non-empty.
//   - ID of user who sent the URL is allowed.
func (b *Bot) addUpdate(update *tgbotapi.Update) {
	if !validMessage(update.Message) {
		return
	}

	if !slices.Contains(b.allowedIDs, update.Message.From.ID) {
		return
	}

	b.messageQueue <- *update
}

// handleUpdates runs goroutine for receiving updates through channel and adds them
// to the queue if they are valid.
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	go func() {
		for update := range updates {
			b.addUpdate(&update)
		}
	}()
}

// handleMessages processes messages queue. If regex URL was found it gets added to
// tracks line.
func (b *Bot) handleMessages(q Queue) error {
	go func() {
		for {
			update := <-q
			text := messageText(update.Message)

			if !utils.IsSoundcloudURL(text) {

			}
		}
	}()

	return nil
}
