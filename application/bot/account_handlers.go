package bot

import (
	"context"
	"migtationbot/application/app"
	"migtationbot/application/keyboard"

	"github.com/go-telegram/bot"
)

func (a *Application) HandlerAccount(ctx context.Context, args ...any) error {
	data := args[0].(HandlerArgs)
	userID := data.UserID
	msgID := data.MsgID

	_, err := a.B.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      userID,
		MessageID:   msgID,
		Text:        app.AccountMainMenuText,
		ReplyMarkup: keyboard.AccountMainMenu(),
	})
	if err != nil {
		return err
	}
	return nil
}
