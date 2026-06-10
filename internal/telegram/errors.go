package telegram

import "errors"

// Defining errors.
var (
	ErrHandlerInternal = errors.New("handler err")
)

// Telegram bot error.
var (
	ErrQueueProcess  = errors.New("queue process err")
	ErrClientInit    = errors.New("bot client init")
	ErrInvalidClient = errors.New("invalid client")
	ErrInvalidToken  = errors.New("invalid token")
	ErrEmptyQueue    = errors.New("empty queue")
	ErrTelegramAPI   = errors.New("telegram api err")
	ErrHandlerInit   = errors.New("handler init err")
)
