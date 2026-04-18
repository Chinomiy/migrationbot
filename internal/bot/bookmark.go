package bot

import (
	"context"
	"migtationbot/internal/app"
	"migtationbot/internal/keyboard"
	"migtationbot/logger"
)

func (h *Handler) Favorite(ctx context.Context, args Args) error {

	userFavorite, err := h.bookmarkSVC.GetUserFavorites(ctx, args.userID)
	if err != nil {
		logger.Error(err)
		return err
	}
	kb := keyboard.UserBookmarks(userFavorite)

	err = h.editMassage(ctx, args.userID, args.msgID, app.AccountBookmarksText, kb)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) AddFavorite(ctx context.Context, args Args) error {
	code := getCodeFromCallbackData(args.rawCBData)
	trip := getTripFromCallbackData(args.rawCBData)

	err := h.bookmarkSVC.AddFavorite(ctx, args.userID, code, trip)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) BookmarkDetails(ctx context.Context, args Args) error {
	code := getCodeFromCallbackData(args.rawCBData)
	trip := getTripFromCallbackData(args.rawCBData)

	content, err := h.countrySVC.GetCountryContentByTrip(ctx, code, trip)
	if err != nil {
		return err
	}
	kb := keyboard.BookmarkDetails(code, trip)

	err = h.editMassage(ctx, args.userID, args.msgID, content, kb)

	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) RemoveBookmark(ctx context.Context, args Args) error {
	code := getCodeFromCallbackData(args.rawCBData)
	trip := getTripFromCallbackData(args.rawCBData)

	err := h.bookmarkSVC.RemoveFavorite(ctx, args.userID, code, trip)
	if err != nil {
		return err
	}

	userFavorite, err := h.bookmarkSVC.GetUserFavorites(ctx, args.userID)
	if err != nil {
		return err
	}
	err = h.editMassage(ctx, args.userID, args.msgID, app.AccountBookmarksText, keyboard.UserBookmarks(userFavorite))
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
