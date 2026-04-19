package keyboard

import (
	"fmt"
	"migtationbot/internal/app"

	"github.com/go-telegram/bot/models"
)

func BackKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         app.BackButton,
					CallbackData: app.CallbackBack,
				},
			},
		},
	}
}
func MainMenuButton() []models.InlineKeyboardButton {
	return []models.InlineKeyboardButton{
		{
			Text:         app.MainMenuButton,
			CallbackData: app.CallbackMainMenu,
		},
	}
}
func BackButton() []models.InlineKeyboardButton {
	return []models.InlineKeyboardButton{
		{
			Text:         app.BackButton,
			CallbackData: app.CallbackBack,
		},
	}
}

func DeleteBookmarkButton(code, trip string) []models.InlineKeyboardButton {
	return []models.InlineKeyboardButton{
		{
			Text:         app.DeleteBookmarkButton,
			CallbackData: fmt.Sprintf("%s:%s:%s", app.CallbackRemoveBookmark, trip, code),
		},
	}
}
