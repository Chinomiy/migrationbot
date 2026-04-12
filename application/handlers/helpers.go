package handlers

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

func getCallbackData(u *models.Update) string {
	if u.CallbackQuery != nil {
		return u.CallbackQuery.Data
	}
	return ""
}
