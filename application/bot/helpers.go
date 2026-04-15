package bot

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func getUserID(u *models.Update) int64 {
	if u.CallbackQuery != nil {
		return u.CallbackQuery.From.ID
	}
	if u.Message != nil {
		return u.Message.From.ID
	}
	return 0
}

func (a *Application) editMassage(ctx context.Context, userID int64, msgID int, content string, kb *models.InlineKeyboardMarkup) error {
	_, err := a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      userID,
		MessageID:   msgID,
		Text:        content,
		ReplyMarkup: kb,
	})
	if err != nil {
		return err
	}

	return nil
}
