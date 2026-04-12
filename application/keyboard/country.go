package keyboard

import (
	"fmt"
	"migtationbot/application/app"
	"migtationbot/application/country"

	"github.com/go-telegram/bot/models"
)

func CountryTripVariants(c *country.Country) *models.InlineKeyboardMarkup {
	var keyboard [][]models.InlineKeyboardButton
	var row []models.InlineKeyboardButton

	tt := c.TripTypes
	for i, t := range tt.Data {
		btn := models.InlineKeyboardButton{
			Text:         "📌 " + t,
			CallbackData: fmt.Sprintf("%s:%s:%s", app.CallbackCountryDetailsMenu, i, c.Code),
		}
		row = append(row, btn)

		// по 2 кнопки в ряд
		if len(tt.Data)%2 == 0 {
			keyboard = append(keyboard, row)
			row = []models.InlineKeyboardButton{}
		}
	}
	if len(row) > 0 {
		keyboard = append(keyboard, row)
	}
	keyboard = append(keyboard, BackButton())

	return &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}
