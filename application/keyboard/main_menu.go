package keyboard

import (
	"migtationbot/application/app"

	"github.com/go-telegram/bot/models"
)

// пока только страны реализовано
func MainMenuKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         "🌍 Выбрать страну",
					CallbackData: string(app.StateCountryMenu),
				},
			},
			{
				{
					Text:         "💼 Личный кабинет",
					CallbackData: "123",
				},
			},
			{
				{
					Text:         "📚 FAQ",
					CallbackData: "242",
				},
				{
					Text:         "🛠 Поддержка",
					CallbackData: "5215",
				},
			},
		},
	}
}
