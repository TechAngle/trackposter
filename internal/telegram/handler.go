// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package telegram

import (
	"context"
	"slices"

	"trackposter/internal/telegram/template/msg"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// addUpdate validates update and adds if it possible.
//
// Requirements for update:
//   - Update Message is found non-empty.
//   - ID of user who sent the URL is allowed.
func (c *Client) addUpdate(update *tgbotapi.Update) {
	if !(validMessage(update.Message) || update.SentFrom() == nil) {
		return
	}

	if !slices.Contains(c.allowedIDs, update.SentFrom().ID) {
		return
	}

	c.messageQueue <- *update
}

func (c *Client) handleUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case update, ok := <-updates:
				if !ok {
					return
				}

				c.addUpdate(&update)
				c.logger.InfoContext(ctx, "updated queue with update", "update_id", update.UpdateID)
			}
		}
	}()
}

func (c *Client) processQueue(ctx context.Context, q *Queue) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case u, ok := <-*q:
				if !ok {
					return
				}

				if !validMessage(u.Message) {
					continue
				}

				c.logger.InfoContext(ctx, "processing update", "update_id", u.UpdateID)

				var err error

				switch {
				case u.Message.IsCommand():
					c.logger.InfoContext(ctx, "processing command", "command", u.Message.Command())
					err = c.commandHandler.Handle(ctx, &u)
				default:
					c.logger.InfoContext(ctx, "processing default message", "text", messageText(u.Message))
					err = c.handleURL(ctx, &u)
				}

				if err != nil {
					c.logger.ErrorContext(ctx, "queue error", "error", err)
				}
			}
		}
	}()
}

func (c *Client) handleURL(ctx context.Context, u *tgbotapi.Update) error {
	url := messageText(u.Message)
	if !c.connector.IsTrackValid(ctx, url) {
		if _, err := c.client.Request(msg.DefaultMessage(
			u.FromChat().ID,
			msg.ErrorInvalidURL(),
		)); err != nil {
			return err
		}
		return nil
	}

	m, err := c.client.Send(msg.DefaultMessage(
		u.FromChat().ID,
		msg.MetadataSearchInfo(),
	))
	if err != nil {
		if _, err := c.client.Request(msg.DefaultEditMessage(
			u.FromChat().ID,
			m.MessageID,
			msg.ErrorInternal(err),
		)); err != nil {
			c.logger.ErrorContext(ctx, "edit message err", "error", err)
		}
		return err
	}

	trackID, err := c.processTrack(ctx, url)
	if err != nil {
		return err
	}

	if _, err := c.client.Request(msg.DefaultEditMessage(
		u.FromChat().ID,
		m.MessageID,
		string(msg.TrackAdded(trackID)),
	)); err != nil {
		return err
	}

	return nil
}

func (c *Client) processTrack(ctx context.Context, url string) (string, error) {
	metadata, err := c.connector.TrackMetadataFromURL(ctx, url)
	if err != nil {
		return "", err
	}

	id, err := c.repository.AddTrack(metadata.AsTrack())
	if err != nil {
		return "", err
	}

	return id, nil
}
