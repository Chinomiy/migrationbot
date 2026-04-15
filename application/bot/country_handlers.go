package bot

import (
	"context"
	"migtationbot/application/app"
	"migtationbot/application/keyboard"
)

func (a *Application) HandlerCountryMenu(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID
	err := a.renderState(ctx, userID, msgID, app.StateCountryMenu)
	if err != nil {
		return err
	}

	return nil
}
func (a *Application) HandlerCountryDetails(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID
	code := data.Code
	trip := data.Trip

	content, err := a.CountrySVC.GetCountryContentByTrip(ctx, code, trip)
	if err != nil {
		return err
	}
	kb := keyboard.CountryKeyboard(code, trip)
	err = a.editMassage(ctx, userID, msgID, content, kb)
	return nil
}

func (a *Application) HandlerCountryTrip(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID
	code := data.Code

	country, err := a.CountrySVC.GetCountryWithTrip(ctx, code)
	if err != nil {
		return err
	}
	kb := keyboard.CountryTripVariants(country)
	err = a.editMassage(ctx, userID, msgID, app.CountryTripText, kb)

	if err != nil {
		return err
	}
	return nil
}
