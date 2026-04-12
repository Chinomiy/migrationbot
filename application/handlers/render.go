package handlers

import (
	"context"
	"migtationbot/application/app"
	"migtationbot/application/keyboard"
	"migtationbot/fsm"

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
		return err
	case app.StateCountryMenu:
		countries, err := a.CountrySVC.List(ctx)
		if err != nil {
			return err
		}
		_, err = a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      userID,
			MessageID:   msgID,
			Text:        app.CountryMenuText,
			ReplyMarkup: keyboard.CountryMenu(countries),
		})
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
	default:
		_, err := a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      userID,
			MessageID:   msgID,
			Text:        app.MainText,
			ReplyMarkup: keyboard.MainMenuKeyboard(),
		})
		return err
	}
	return nil
}
