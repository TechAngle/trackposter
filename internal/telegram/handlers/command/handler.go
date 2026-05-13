// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package command

import (
	"context"
	"log/slog"
	"strings"

	"trackposter/internal/domain"
	"trackposter/internal/telegram/template/msg"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Func func(context.Context, *tgbotapi.Update) error

type HandlerOptions struct {
	Client     *tgbotapi.BotAPI
	Repository domain.Repository
	Connector  domain.Connector
	Logger     *slog.Logger
}

type Handler struct {
	client     *tgbotapi.BotAPI
	repository domain.Repository
	connector  domain.Connector
	commands   map[string]Func
	logger     *slog.Logger
}

var _ domain.TelegramHandler = (*Handler)(nil)

// NewHandler creates new handler with preregistered commands.
// Uses slog.Default() if logger value is nil.
func NewHandler(opts HandlerOptions) (*Handler, error) {
	if opts.Client == nil {
		return nil, domain.ErrNilClient
	}

	logger := opts.Logger
	if logger == nil {
		logger = slog.Default()
	}

	h := &Handler{
		client:     opts.Client,
		repository: opts.Repository,
		connector:  opts.Connector,
		logger:     logger,
	}
	registerCommands(h)

	return h, nil
}

// Handle processes incoming update and returns error if something went wrong.
func (h *Handler) Handle(ctx context.Context, u *tgbotapi.Update) error {
	return h.handleCommand(ctx, u)
}

func (h *Handler) handleCommand(ctx context.Context, u *tgbotapi.Update) (err error) {
	command := strings.ToLower(u.Message.Command())

	if fn, ok := h.commands[command]; ok {
		return fn(ctx, u)
	}

	return
}

func registerCommands(h *Handler) {
	h.commands = map[string]Func{
		"queue":    h.getQueueCmd,
		"download": h.downloadCmd,
	}
}

func (h *Handler) getQueueCmd(ctx context.Context, u *tgbotapi.Update) error {
	queue := h.repository.Queue()
	if len(queue) == 0 {
		h.sendRequest(ctx, msg.DefaultMessage(
			u.FromChat().ID,
			msg.ErrorEmptyQueue(),
		))
		return nil
	}

	h.sendRequest(ctx, msg.DefaultMessage(
		u.FromChat().ID,
		msg.TracksQueue(queue)))

	return nil
}

func (h *Handler) downloadCmd(ctx context.Context, u *tgbotapi.Update) error {
	files, err := h.downloadQueue(ctx)
	if err != nil {
		h.sendRequest(ctx, msg.DefaultMessage(
			u.FromChat().ID,
			msg.ErrorDownload(err),
		))
		return nil
	}

	h.sendRequest(ctx, tgbotapi.NewMediaGroup(
		u.FromChat().ID,
		files,
	))

	return nil
}

func (h *Handler) downloadQueue(ctx context.Context) ([]any, error) {
	queue := h.repository.Queue()
	if len(queue) == 0 {
		return nil, domain.ErrEmptyQueue
	}

	files := make([]any, 0, len(queue))

	for _, track := range queue {
		h.logger.InfoContext(ctx,
			"processing track",
			"author", track.Author,
			"title", track.Title,
		)
		inputMedia, err := h.trackBytes(ctx, track)
		if err != nil {
			h.logger.ErrorContext(ctx, "track bytes resolve err", "error", err)
			continue
		}

		files = append(files, inputMedia)
	}

	return files, nil
}

func (h *Handler) trackBytes(ctx context.Context, track *domain.Track) (tgbotapi.InputMediaAudio, error) {
	trackContent, err := h.connector.TrackFromURL(ctx, track.URL)
	if err != nil {
		return tgbotapi.InputMediaAudio{}, err
	}

	name := formatTrackName(track)

	return tgbotapi.NewInputMediaAudio(tgbotapi.FileBytes{
		Name:  name,
		Bytes: trackContent,
	}), nil
}

func (h *Handler) sendRequest(ctx context.Context, c tgbotapi.Chattable) {
	_, err := h.client.Request(c)
	if err != nil {
		h.logger.ErrorContext(ctx, "telegram request err", "error", err)
	}
}

func formatTrackName(track *domain.Track) string {
	var name strings.Builder
	name.Grow(len(track.Author) + len(track.Title) + 3)
	name.WriteString(track.Author)
	name.WriteString(" - ")
	name.WriteString(track.Title)

	return name.String()
}
