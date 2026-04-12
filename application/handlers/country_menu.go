package handlers

import (
	"context"
	"migtationbot/application/app"
)

func (a *Application) HandlerCountryMenu(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID
	a.renderState(ctx, userID, msgID, app.StateCountryMenu)

	return nil
}
