// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package telegram

import (
	"fmt"
	"strings"
)

type Bot struct {
}

// Initializate new bot instannce
func NewBot(options BotOptions) (*Bot, error) {
	if strings.TrimSpace(options.Token) == "" {
		return nil, fmt.Errorf("invalid token")
	}

	return &Bot{}, nil
}
