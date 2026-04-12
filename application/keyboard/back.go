package keyboard

import (
	"migtationbot/application/app"

	"github.com/go-telegram/bot/models"
)

func BackKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         "⬅️ Назад",
					CallbackData: app.CallbackBack,
				},
			},
		},
	}
}
func BackButton() []models.InlineKeyboardButton {
	return []models.InlineKeyboardButton{
		{
			Text:         "⬅️ Назад",
			CallbackData: app.CallbackBack,
		},
	}
}
