// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package telegram

import (
	"strings"

	"trackposter/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
