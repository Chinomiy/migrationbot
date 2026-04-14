package bot

import (
	"context"
	"migtationbot/application/app"
	"migtationbot/application/keyboard"
	"migtationbot/logger"

	"github.com/go-telegram/bot"
)

func (a *Application) HandlerFavorite(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID

	userFavorite, err := a.BookmarkSVc.GetUserFavorites(ctx, userID)
	if err != nil {
		logger.Error(err)
		return err
	}
	kb := keyboard.UserBookmarks(userFavorite)
	_, err = a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      userID,
		MessageID:   msgID,
		Text:        "закладки:",
		ReplyMarkup: kb,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *Application) HandlerAddFavorite(ctx context.Context, args HandlerArgs) error {
	userID := args.UserID
	code := args.Code
	trip := args.Trip

	err := a.BookmarkSVc.AddFavorite(ctx, userID, code, trip)
	if err != nil {
		return err
	}
	return nil
}

func (a *Application) HandlerBookmarkDetails(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID
	code := data.Code
	trip := data.Trip

	content, err := a.CountrySVC.GetCountryContentByTrip(ctx, code, trip)
	if err != nil {
		return err
	}
	kb := keyboard.BookmarkDetails(code, trip)
	_, err = a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
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

func (a *Application) HandlerRemoveBookmark(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID
	code := data.Code
	trip := data.Trip

	err := a.BookmarkSVc.RemoveFavorite(ctx, userID, code, trip)
	if err != nil {
		return err
	}
	a.renderState(ctx, userID, msgID, app.StateAccount)
	return nil
}
