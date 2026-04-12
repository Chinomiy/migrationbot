package handlers

import (
	"context"
	"migtationbot/application/keyboard"

	"github.com/go-telegram/bot"
)

func (a *Application) HandlerCountryDetails(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID
	code := data.Code
	trip := data.Trip

	content, err := a.CountrySVC.GetCountryContentByTrip(ctx, *code, *trip)
	if err != nil {
		return err
	}
	kb := keyboard.BackKeyboard()
	_, err = a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      userID,
		MessageID:   msgID,
		Text:        content,
		ReplyMarkup: kb,
	})
	return nil
}
