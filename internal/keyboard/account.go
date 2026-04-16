package keyboard

import (
	"migtationbot/internal/app"

	"github.com/go-telegram/bot/models"
)

func AccountMainMenu() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         app.BookmarksButton,
					CallbackData: app.CallbackFavorite,
				},
			},
			BackButton(),
		},
	}
}
