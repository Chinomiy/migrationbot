package handlers

import (
	"context"
	"log"
	"migtationbot/application/app"
	"migtationbot/application/keyboard"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type HandlerArgs struct {
	UserID int64
	MsgID  int
	Code   *string
	Trip   *string
}

func (a *Application) TextRouter(
	ctx context.Context,
	b *bot.Bot,
	u *models.Update,
) {
	if u.Message == nil {
		return
	}
	userID := getUserID(u)
	currentState, _ := a.F.Current(userID)
	switch currentState.ID {
	case app.StateMainMenu:
		kb := keyboard.MainMenuKeyboard()
		_, _ = a.B.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      userID,
			Text:        app.MainText,
			ReplyMarkup: kb,
		})
	}
	return
}

func (a *Application) CallbackRouter(
	ctx context.Context,
	b *bot.Bot,
	u *models.Update,
) {
	if u.CallbackQuery == nil {
		return
	}
	userID := getUserID(u)
	msgID := u.CallbackQuery.Message.Message.ID
	data := u.CallbackQuery.Data

	defer b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: u.CallbackQuery.ID,
	})

	switch getRawCallbackData(data) {
	case app.CallbackCountryMenu:
		if err := a.F.Transition(
			ctx,
			userID,
			app.StateCountryMenu,
			HandlerArgs{
				UserID: userID,
				MsgID:  msgID,
			},
		); err != nil {
			log.Println(err)
			return
		}
	case app.CallbackCountry:
		if err := a.F.Transition(
			ctx,
			userID,
			app.StateCountry,
			HandlerArgs{
				UserID: userID,
				MsgID:  msgID,
				Code:   getCodeFromCallbackData(data),
			},
		); err != nil {
			log.Println(err)
			return
		}
		return
	case app.CallbackCountryDetailsMenu:
		if err := a.F.Transition(
			ctx,
			userID,
			app.StateCountryDetailsMenu,
			HandlerArgs{
				UserID: userID,
				MsgID:  msgID,
				Code:   getCodeFromCallbackData(data),
				Trip:   getTripFromCallbackData(data),
			},
		); err != nil {
		}
		return

	case app.CallbackBack:
		if err := a.F.Back(
			ctx,
			userID,
		); err != nil {
			log.Println(err)
			return
		}
		current, _ := a.F.Current(userID)
		a.renderState(
			ctx,
			userID,
			msgID,
			current.ID,
		)

	default:
		return
	}
}

func getRawCallbackData(data string) string {
	return strings.Split(data, ":")[0]
}

func getCodeFromCallbackData(data string) *string {
	rawData := strings.Split(data, ":")
	if len(rawData) == 2 {
		return &rawData[1]
	}
	if len(rawData) == 3 {
		return &rawData[2]
	}
	return nil
}
func getTripFromCallbackData(data string) *string {
	rawData := strings.Split(data, ":")
	if len(rawData) == 3 {
		return &rawData[1]
	}
	return nil
}
