package telegram

import (
	"context"
	"errors"
	"slices"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"trackposter/internal/domain"
	"trackposter/internal/telegram/template/msg"
)

// addUpdate validates update and adds if it possible.
//
// Requirements for update:
//   - Update Message is found non-empty.
//   - ID of user who sent the URL is allowed.
func (c *Client) addUpdate(update *tgbotapi.Update) {
	if !validMessage(update.Message) && update.SentFrom() != nil {
		return
	}

	if !slices.Contains(c.allowedIDs, update.SentFrom().ID) {
		return
	}

	c.messageQueue <- *update
}

func (c *Client) handleUpdates(
	ctx context.Context,
	updates tgbotapi.UpdatesChannel,
) {
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
				c.logger.InfoContext(
					ctx,
					"updated queue with update",
					"update_id",
					update.UpdateID,
				)
			}
		}
	}()
}

func (c *Client) processQueue(ctx context.Context, queue *Queue) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case update, ok := <-*queue:
				if !ok {
					return
				}

				if !validMessage(update.Message) {
					continue
				}

				c.logger.InfoContext(
					ctx,
					"processing update",
					"update_id",
					update.UpdateID,
				)

				var err error

				switch {
				case update.Message.IsCommand():
					c.logger.InfoContext(
						ctx,
						"processing command",
						"command",
						update.Message.Command(),
					)
					err = c.commandHandler.Handle(ctx, &update)
				default:
					c.logger.InfoContext(
						ctx,
						"processing default message",
						"text",
						messageText(update.Message),
					)
					err = c.handleURL(ctx, &update)
				}

				if err != nil {
					c.logger.ErrorContext(ctx, "queue error", "error", err)
				}
			}
		}
	}()
}

func (c *Client) handleURL(ctx context.Context, update *tgbotapi.Update) error {
	url := messageText(update.Message)
	if !c.connector.IsTrackValid(ctx, url) {
		_, err := c.client.Request(msg.DefaultMessage(
			update.FromChat().ID,
			msg.ErrorInvalidURL(),
		))
		if err != nil {
			return errors.Join(domain.ErrTelegramAPI, err)
		}

		return nil
	}

	message, err := c.client.Send(msg.DefaultMessage(
		update.FromChat().ID,
		msg.MetadataSearchInfo(),
	))
	if err != nil {
		_, err := c.client.Request(msg.DefaultEditMessage(
			update.FromChat().ID,
			message.MessageID,
			msg.ErrorInternal(err),
		))
		if err != nil {
			c.logger.ErrorContext(ctx, "edit message err", "error", err)
		}

		return errors.Join(domain.ErrTelegramAPI, err)
	}

	trackID, err := c.processTrack(ctx, url)
	if err != nil {
		return err
	}

	_, err = c.client.Request(msg.DefaultEditMessage(
		update.FromChat().ID,
		message.MessageID,
		string(msg.TrackAdded(trackID)),
	))
	if err != nil {
		return errors.Join(domain.ErrTelegramAPI, err)
	}

	return nil
}

func (c *Client) processTrack(ctx context.Context, url string) (string, error) {
	metadata, err := c.connector.TrackMetadataFromURL(ctx, url)
	if err != nil {
		return "", errors.Join(domain.ErrConnectorInternal, err)
	}

	trackID, err := c.repository.AddTrack(metadata.AsTrack())
	if err != nil {
		return "", errors.Join(domain.ErrRepositoryInternal, err)
	}

	return trackID, nil
}
