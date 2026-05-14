// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package domain

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramHandler represents default handler for Telegram updates.
type TelegramHandler interface {
	// Handle processes incoming update and returns error if something went wrong.
	Handle(ctx context.Context, update *tgbotapi.Update) error
}

// MessageTemplate represents type for Telegram messages.
type MessageTemplate string

// String returns default representation for type.
func (t MessageTemplate) String() string {
	return string(t)
}
