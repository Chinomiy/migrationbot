package bot

import (
	"context"
	"migtationbot/internal/app"
	"migtationbot/internal/keyboard"
)

func (h *Handler) Account(ctx context.Context, args Args) error {

	err := h.editMassage(ctx, args.userID, args.msgID, app.AccountMainMenuText, keyboard.AccountMainMenu())

	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) FAQMenu(ctx context.Context, args Args) error {

	err := h.editMassage(ctx, args.userID, args.msgID, app.FAQMenuText, keyboard.FAQMenu())
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) HelpMenu(ctx context.Context, args Args) error {

	err := h.editMassage(ctx, args.userID, args.msgID, app.HelpText, keyboard.FAQMenu())
	if err != nil {
		return err
	}
	return nil
}
