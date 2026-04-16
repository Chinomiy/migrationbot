package bot

import (
	"context"
	"migtationbot/fsm"
	"migtationbot/internal/app"
	"migtationbot/internal/keyboard"
	"migtationbot/logger"
)

func (a *Application) renderState(
	ctx context.Context,
	userID int64,
	msgID int,
	stateID fsm.StateID,
) error {
	switch stateID {

	case app.StateMainMenu:
		err := a.editMassage(ctx, userID, msgID, app.MainText, keyboard.MainMenuKeyboard())
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
		err = a.editMassage(ctx, userID, msgID, app.CountryMenuText, keyboard.CountryMenu(countries))
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
		err = a.editMassage(ctx, userID, msgID, app.CountryMenuText, keyboard.CountryTripVariants(countries))
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil

	case app.StateAccount:
		err := a.editMassage(ctx, userID, msgID, app.AccountMainMenuText, keyboard.AccountMainMenu())
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
		err = a.editMassage(ctx, userID, msgID, app.AccountBookmarksText, keyboard.UserBookmarks(userBookmark))
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil

	default:
		err := a.editMassage(ctx, userID, msgID, app.MainText, keyboard.MainMenuKeyboard())
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil
	}
}
