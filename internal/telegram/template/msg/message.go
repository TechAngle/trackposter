// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in
// LICENSE file.
//
// Author: https://github.com/TechAngle

// Package msg provides templates with HTML formatting for Telegram.
package msg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"trackposter/internal/domain"
)

// DefaultMessage returns default message config with pre-defined HTML parsing.
func DefaultMessage[T domain.MessageTemplate | string](
	chatID int64,
	text T,
) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, string(text))
	msg.ParseMode = tgbotapi.ModeHTML

	return msg
}

// DefaultEditMessage returns default edit message config with predefined HTML
// parsing.
func DefaultEditMessage[T domain.MessageTemplate | string](
	chatID int64,
	msgID int,
	text T,
) tgbotapi.EditMessageTextConfig {
	config := tgbotapi.NewEditMessageText(
		chatID,
		msgID,
		string(text),
	)
	config.ParseMode = tgbotapi.ModeHTML

	return config
}
