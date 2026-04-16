package bot

import (
	"context"
	"migtationbot/internal/app"
	"migtationbot/internal/keyboard"
)

func (a *Application) HandlerManagerMenu(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID
	cb := data.RawCallbackData
	switch cb {
	case app.CallbackManagerMenu:
		kb := keyboard.ManagerMenuKeyboard()
		a.editMassage(ctx, userID, msgID, "Менеджерское меню", kb)
	}
	return nil
}
