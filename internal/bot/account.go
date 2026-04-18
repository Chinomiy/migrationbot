package bot

import (
	"context"
	"migtationbot/internal/app"
	"migtationbot/internal/keyboard"
)

func (a *Handler) Account(ctx context.Context, args Args) error {

	err := a.editMassage(ctx, args.userID, args.msgID, app.AccountMainMenuText, keyboard.AccountMainMenu())

	if err != nil {
		return err
	}
	return nil
}
