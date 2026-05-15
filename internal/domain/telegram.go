package domain

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

// MessageTemplate represents type for Telegram messages.
type MessageTemplate string

// String returns default representation for type.
func (t MessageTemplate) String() string {
	return string(t)
}
