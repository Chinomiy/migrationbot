package bot

import (
	"context"
	"migtationbot/internal/app"
	"migtationbot/internal/keyboard"
)

func (a *Application) HandlerAccount(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID

	err := a.editMassage(ctx, userID, msgID, app.AccountMainMenuText, keyboard.AccountMainMenu())

	if err != nil {
		return err
	}
	return nil
}
