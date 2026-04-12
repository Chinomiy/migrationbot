package handlers

import (
	"context"
	"migtationbot/application/app"
	"migtationbot/application/keyboard"

	"github.com/go-telegram/bot"
)

func (a *Application) HandlerCountryTrip(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID
	code := data.Code

	country, err := a.CountrySVC.GetCountryWithTrip(ctx, *code)
	if err != nil {
		return err
	}
	kb := keyboard.CountryTripVariants(country)
	_, err = a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      userID,
		MessageID:   msgID,
		Text:        app.CountryTripText,
		ReplyMarkup: kb,
	})
	if err != nil {
		return err
	}
	return nil
}
