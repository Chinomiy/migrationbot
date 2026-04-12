package keyboard

import (
	"fmt"
	"migtationbot/application/app"
	"migtationbot/application/country"

	"github.com/go-telegram/bot/models"
)

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
