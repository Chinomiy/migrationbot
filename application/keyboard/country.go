package keyboard

import (
	"fmt"
	"migtationbot/application/app"
	"migtationbot/application/country"

	"github.com/go-telegram/bot/models"
)

func CountryKeyboard(code, trip string) *models.InlineKeyboardMarkup {
	var keyboard [][]models.InlineKeyboardButton

	keyboard = append(keyboard, BackButton())
	keyboard = append(keyboard, AddBookmarkButton(code, trip))

	return &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

func CountryMenu(c *[]country.Country) *models.InlineKeyboardMarkup {
	var keyboard [][]models.InlineKeyboardButton
	var row []models.InlineKeyboardButton

	for i, ctr := range *c {
		btn := models.InlineKeyboardButton{
			Text:         "🌍 " + ctr.Name,
			CallbackData: fmt.Sprintf("%s:%s", app.CallbackCountry, ctr.Code),
		}

		row = append(row, btn)
		if (i+1)%2 == 0 {
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
		if len(row)%2 == 0 {
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
