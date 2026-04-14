package keyboard

import (
	"migtationbot/application/app"

	"github.com/go-telegram/bot/models"
)

func AccountMainMenu() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         app.AccountMenuBookmarks,
					CallbackData: app.CallbackFavorite,
				},
			},
			BackButton(),
		},
	}
}
