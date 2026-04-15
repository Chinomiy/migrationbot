package bot

import (
	"context"
	"log"
	"migtationbot/application/app"
	"migtationbot/application/keyboard"
	"migtationbot/logger"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type HandlerArgs struct {
	UserID int64
	MsgID  int
	Code   string
	Trip   string
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
	tgUsername := u.Message.From.Username
	_, err := a.UserSVC.GetOrCreateUser(ctx, userID, tgUsername)
	if err != nil {
		logger.Error(err)
	}
	currentState, _ := a.F.Current(userID)

	switch currentState.ID {
	case app.StateMainMenu:
		kb := keyboard.MainMenuKeyboard()
		_, _ = a.B.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      userID,
			Text:        app.MainText,
			ReplyMarkup: kb,
		})
	default:
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

	logger.Infof("RAW CALLBACK DATA: %s", data)

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
			logger.Error(err)
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
			logger.Error(err)
		}
		return
	case app.CallbackAccount:
		if err := a.F.Transition(
			ctx,
			userID,
			app.StateAccount,
			HandlerArgs{
				UserID: userID,
				MsgID:  msgID,
			}); err != nil {
			logger.Error(err)
			return
		}
		return
	case app.CallbackAddFavorite:
		state, err := a.F.Current(userID)
		if err != nil {
			logger.Error(err)
			return
		}
		data := state.Data

		if err = a.HandlerAddFavorite(ctx, HandlerArgs{
			UserID: userID,
			MsgID:  msgID,
			Code:   data[0].(HandlerArgs).Code,
			Trip:   data[0].(HandlerArgs).Trip,
		}); err != nil {
			logger.Error(err)
			return
		}
		// уведомление что добавлено
		return
	case app.CallbackFavorite:
		if err := a.F.Transition(
			ctx,
			userID,
			app.StateFavorite,
			HandlerArgs{
				UserID: userID,
				MsgID:  msgID,
			}); err != nil {
			logger.Error(err)
			return
		}
		return

	case app.CallbackBookmarkDetails:
		if err := a.F.Transition(
			ctx,
			userID,
			app.StateBookmarkDetails,
			HandlerArgs{
				UserID: userID,
				MsgID:  msgID,
				Trip:   getTripFromCallbackData(data),
				Code:   getCodeFromCallbackData(data),
			}); err != nil {
			logger.Error(err)
			return
		}
		return

	case app.CallbackRemoveBookmark:
		if err := a.HandlerRemoveBookmark(
			ctx,
			HandlerArgs{
				UserID: userID,
				MsgID:  msgID,
				Trip:   getTripFromCallbackData(data),
				Code:   getCodeFromCallbackData(data),
			},
		); err != nil {
			logger.Error(err)
			return
		}
		return
	case app.CallbackBack:
		if err := a.F.Back(
			ctx,
			userID,
		); err != nil {
			logger.Error(err)
			return
		}
		current, _ := a.F.Current(userID)
		logger.Infof("BACK STATE: %s", current)
		err := a.renderState(
			ctx,
			userID,
			msgID,
			current.ID,
		)
		if err != nil {
			logger.Error(err)
			return
		}

	default:
		return
	}
}

func getRawCallbackData(data string) string {
	return strings.Split(data, ":")[0]
}

func getCodeFromCallbackData(data string) string {
	rawData := strings.Split(data, ":")
	if len(rawData) == 2 {
		return rawData[1]
	}
	if len(rawData) == 3 {
		return rawData[2]
	}
	return ""
}
func getTripFromCallbackData(data string) string {
	rawData := strings.Split(data, ":")
	if len(rawData) == 3 {
		return rawData[1]
	}
	return ""
}
