package bot

import (
	"context"
	"migtationbot/internal/app"
	"migtationbot/internal/keyboard"
	"migtationbot/logger"
)

func (h *Handler) MainMenu(ctx context.Context, args Args) error {

	err := h.editMassage(ctx, args.userID, args.msgID, app.MainText, keyboard.MainMenuKeyboard())
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (h *Handler) CountryMenu(ctx context.Context, args Args) error {

	country, err := h.countrySVC.List(ctx)
	if err != nil {
		logger.Error(err)
		return err
	}
	err = h.editMassage(ctx, args.userID, args.msgID, app.CountryMenuText, keyboard.CountryMenu(country))
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (h *Handler) CountryDetails(ctx context.Context, args Args) error {
	code := getCodeFromCallbackData(args.rawCBData)
	trip := getTripFromCallbackData(args.rawCBData)
	content, err := h.countrySVC.GetCountryContentByTrip(ctx, code, trip)

	if err != nil {
		return err
	}
	kb := keyboard.CountryKeyboard(code, trip)
	err = h.editMassage(ctx, args.userID, args.msgID, content, kb)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (h *Handler) CountryTrip(ctx context.Context, args Args) error {
	code := getCodeFromCallbackData(args.rawCBData)

	country, err := h.countrySVC.GetCountryWithTrip(ctx, code)
	if err != nil {
		return err
	}

	kb := keyboard.CountryTripVariants(country)
	err = h.editMassage(ctx, args.userID, args.msgID, app.CountryTripText, kb)

	if err != nil {
		return err
	}
	return nil
}
