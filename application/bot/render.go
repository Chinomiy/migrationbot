package bot

import (
	"context"
	"migtationbot/application/app"
	"migtationbot/application/keyboard"
	"migtationbot/fsm"
	"migtationbot/logger"

	"github.com/go-telegram/bot"
)

func (a *Application) renderState(
	ctx context.Context,
	userID int64,
	msgID int,
	stateID fsm.StateID,
) error {
	switch stateID {
	case app.StateMainMenu:
		_, err := a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      userID,
			MessageID:   msgID,
			Text:        app.MainText,
			ReplyMarkup: keyboard.MainMenuKeyboard(),
		})
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil
	case app.StateCountryMenu:
		countries, err := a.CountrySVC.List(ctx)
		if err != nil {
			logger.Error(err)
			return err
		}
		_, err = a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      userID,
			MessageID:   msgID,
			Text:        app.CountryMenuText,
			ReplyMarkup: keyboard.CountryMenu(countries),
		})
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil
	case app.StateCountry:
		countries, err := a.CountrySVC.FindByCode(ctx, "")
		if err != nil {
			return err
		}
		_, err = a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      userID,
			MessageID:   msgID,
			Text:        app.CountryMenuText,
			ReplyMarkup: keyboard.CountryTripVariants(countries),
		})
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil

	case app.StateAccount:
		_, err := a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      userID,
			MessageID:   msgID,
			Text:        app.AccountMainMenuText,
			ReplyMarkup: keyboard.AccountMainMenu(),
		})
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil
	case app.StateFavorite:
		userBookmark, err := a.BookmarkSVc.GetUserFavorites(ctx, userID)
		if err != nil {
			logger.Error(err)
			return err
		}
		_, err = a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      userID,
			MessageID:   msgID,
			Text:        "Закладки:",
			ReplyMarkup: keyboard.UserBookmarks(userBookmark),
		})
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil

	default:
		_, err := a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      userID,
			MessageID:   msgID,
			Text:        app.MainText,
			ReplyMarkup: keyboard.MainMenuKeyboard(),
		})
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil
	}
}
