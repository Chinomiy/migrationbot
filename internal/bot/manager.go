package bot

import (
	"context"
	"migtationbot/internal/app"
	"migtationbot/internal/keyboard"
)

func (h *Handler) ManagerMenu(ctx context.Context, args Args) error {

	switch callbackData(args.rawCBData) {
	case app.CallbackManagerMenu:
		kb := keyboard.ManagerMenuKeyboard()
		err := h.editMassage(ctx, args.userID, args.msgID, "Менеджерское меню", kb)
		if err != nil {
			return err
		}
	}
	return nil
}
