package bot

import (
	"migtationbot/fsm"
	"strings"

	"github.com/go-telegram/bot/models"
)

func getStateByCallback(callback string) fsm.StateID {

	newState, ok := callbackStateMap[callback]
	if !ok {
		return ""
	}
	return newState
}
func callbackData(data string) string {
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

func setupArgsFromCache(u *models.Update, data []any) Args {
	if len(data) == 0 {
		return Args{
			userID:   u.CallbackQuery.From.ID,
			userName: u.CallbackQuery.From.Username,
			msgID:    u.CallbackQuery.Message.Message.ID,
		}
	}
	return Args{
		userID:    u.CallbackQuery.From.ID,
		userName:  u.CallbackQuery.From.Username,
		msgID:     u.CallbackQuery.Message.Message.ID,
		rawCBData: data[0].(Args).rawCBData,
	}

}

func setupArgsFromCallback(u *models.Update) Args {
	return Args{
		userID:    u.CallbackQuery.From.ID,
		userName:  u.CallbackQuery.From.Username,
		msgID:     u.CallbackQuery.Message.Message.ID,
		rawCBData: u.CallbackQuery.Data,
	}
}
