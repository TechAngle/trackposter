package telegram

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"trackposter/internal/config"
)

// validMessage checks if message is not nil and contains text or caption.
func validMessage(msg *tgbotapi.Message) bool {
	return msg != nil && (messageText(msg) != "")
}

// messageText returns trimmed text or caption depends on what message contains.
func messageText(msg *tgbotapi.Message) string {
	if msg == nil {
		return ""
	}

	if msg.Caption != "" {
		return strings.TrimSpace(msg.Caption)
	}

	return strings.TrimSpace(msg.Text)
}

func newQueue() Queue {
	return make(Queue, config.MaxQueueSize)
}
