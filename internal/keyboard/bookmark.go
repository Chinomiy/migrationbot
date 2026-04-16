package keyboard

import (
	"fmt"
	"migtationbot/internal/app"
	"migtationbot/internal/bookmark"

	"github.com/go-telegram/bot/models"
)

func AddBookmarkButton(code, trip string) []models.InlineKeyboardButton {
	return []models.InlineKeyboardButton{
		{
			Text:         "⬅️ Добавить в закладки",
			CallbackData: fmt.Sprintf("%s:%s:%s", app.CallbackAddFavorite, code, trip),
		},
	}
}

func UserBookmarks(c []bookmark.UserFavorite) *models.InlineKeyboardMarkup {
	var keyboard [][]models.InlineKeyboardButton
	var row []models.InlineKeyboardButton

	for _, user := range c {
		btn := models.InlineKeyboardButton{
			Text:         "📌 " + user.CountryName + " " + user.TripType,
			CallbackData: fmt.Sprintf("%s:%s:%s", app.CallbackBookmarkDetails, user.TripCallback, user.CountryCode),
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

func BookmarkDetails(code, trip string) *models.InlineKeyboardMarkup {
	var keyboard [][]models.InlineKeyboardButton

	keyboard = append(keyboard, BackButton())
	keyboard = append(keyboard, DeleteBookmarkButton(code, trip))

	return &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}
