// Package handler provides an interface for handlers.
package handler

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramHandler represents default handler for Telegram updates.
type TelegramHandler interface {
	// Handle processes incoming update and returns error if something went
	// wrong.
	Handle(ctx context.Context, update *tgbotapi.Update) error
}
