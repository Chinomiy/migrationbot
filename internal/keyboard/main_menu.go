package keyboard

import (
	"migtationbot/internal/app"

	"github.com/go-telegram/bot/models"
)

func MainMenuKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         app.MainMenuCountry,
					CallbackData: app.CallbackCountryMenu,
				},
			},
			{
				{
					Text:         app.MainMenuAccount,
					CallbackData: app.CallbackAccount,
				},
			},
			{
				{
					Text:         app.MainMenuFAQ,
					CallbackData: "242", // доделать просто экран с текстом + назад
				},
				{
					Text:         app.MainMenuHelps,
					CallbackData: "5215", // доделать, просто экран с моим тг + назад
				},
				{
					Text:         "менеджерское меню",
					CallbackData: app.CallbackManagerMenu,
				},
			},
		},
	}
}
