package keyboard

import (
	"migtationbot/internal/app"

	"github.com/go-telegram/bot/models"
)

func ManagerMenuKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         "Главное меню",
					CallbackData: app.CallbackMainMenu,
				},
			},
			{
				{
					Text:         "Создать страну",
					CallbackData: app.CallbackManagerCreateCountry,
				},
			},
			{
				{
					Text:         "Создать тип поездки",
					CallbackData: "242",
				},
				{
					Text:         "Добавить тип поездки к стране",
					CallbackData: "5215",
				},
				{
					Text:         "Добавить контент к типу поездки стране",
					CallbackData: "5215",
				},
			},
		},
	}
}
