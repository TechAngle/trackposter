// Package msg provides templates with HTML formatting for Telegram.
package msg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"trackposter/internal/telegram/template"
)

// DefaultMessage returns default message config with pre-defined HTML parsing.
func DefaultMessage[T template.MessageTemplate | string](
	chatID int64,
	text T,
) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, string(text))
	msg.ParseMode = tgbotapi.ModeHTML

	return msg
}

// DefaultEditMessage returns default edit message config with predefined HTML
// parsing.
func DefaultEditMessage[T template.MessageTemplate | string](
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
