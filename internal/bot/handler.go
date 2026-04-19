package bot

import (
	"context"
	"migtationbot/fsm"
	"migtationbot/internal/app"
	"migtationbot/internal/bookmark"
	"migtationbot/internal/country"
	"migtationbot/internal/user"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var callbackStateMap = map[string]fsm.StateID{
	app.CallbackCountryMenu:        app.StateCountryMenu,
	app.CallbackFavorite:           app.StateFavorite,
	app.CallbackCountry:            app.StateCountry,
	app.CallbackCountryDetailsMenu: app.StateCountryDetailsMenu,
	app.CallbackMainMenu:           app.StateMainMenu,
	app.CallbackFAQ:                app.StateFAQ,
	app.CallbackHelp:               app.StateHelp,

	//manager создание страны
	/*
		app.CallbackManagerMenu:                     app.StateManagerMenu,
		app.CallbackBookmarkDetails:                 app.StateBookmarkDetails,
		app.CallbackManagerCreateCountry:            app.StateManagerCreateCountry,
		app.CallbackManagerCreateCountryCode:        app.StateManagerCreateCountryCode,
		app.CallbackManagerCreateCountryDescription: app.StateManagerCreateCountryDescription,
		app.CallbackManagerCreateCountryConfirm:     app.StateManagerCreateCountryConfirm,
		app.CallbackManagerCreateCountryName:        app.StateManagerCreateCountryName,
	*/

	app.CallbackAccount: app.StateAccount,

	// действия которые не должны менять стек состояний
	app.CallbackAddFavorite:    app.StateNoChange,
	app.CallbackBack:           app.StateNoChange,
	app.CallbackRemoveBookmark: app.StateNoChange,
}

type HandleFunc func(ctx context.Context, args Args) error

type Handler struct {
	B *bot.Bot

	userSVC user.UserService

	bookmarkSVC bookmark.BookmarkService

	countrySVC country.CountryService
}

func NewHandler(
	userSVC user.UserService,
	bookmarkSVC bookmark.BookmarkService,
	countrySVC country.CountryService,
) *Handler {
	return &Handler{
		userSVC: userSVC,

		bookmarkSVC: bookmarkSVC,

		countrySVC: countrySVC,
	}
}

func (h *Handler) editMassage(ctx context.Context, userID int64, msgID int, content string, kb *models.InlineKeyboardMarkup) error {
	_, err := h.B.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      userID,
		MessageID:   msgID,
		Text:        content,
		ReplyMarkup: kb,
	})
	if err != nil {
		return err
	}

	return nil
}
