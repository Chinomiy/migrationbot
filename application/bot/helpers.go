package bot

import (
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
