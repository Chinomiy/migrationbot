package bot

import (
	"context"
	"migtationbot/fsm"
	"migtationbot/internal/app"
	"migtationbot/internal/keyboard"
	"migtationbot/logger"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Args struct {
	rawText   string
	rawCBData string
	userID    int64
	userName  string
	msgID     int
}

type Router struct {
	f *fsm.FSM

	handlers map[fsm.StateID]HandleFunc

	actionHandlers map[string]HandleFunc
}

func NewRouter(f *fsm.FSM) *Router {
	return &Router{
		f:              f,
		handlers:       make(map[fsm.StateID]HandleFunc),
		actionHandlers: make(map[string]HandleFunc),
	}
}
func (r *Router) RegisterHandler(h *Handler) {
	// меняющие состояние хенждеры
	r.handlers[app.StateMainMenu] = h.MainMenu
	r.handlers[app.StateCountryMenu] = h.CountryMenu
	r.handlers[app.StateCountryDetailsMenu] = h.CountryDetails
	r.handlers[app.StateCountry] = h.CountryTrip
	r.handlers[app.StateAccount] = h.Account
	r.handlers[app.StateFavorite] = h.Favorite
	r.handlers[app.StateBookmarkDetails] = h.BookmarkDetails
	r.handlers[app.StateManagerMenu] = h.ManagerMenu

	//хендлеры которые не меняет состояние
	r.actionHandlers[app.CallbackAddFavorite] = h.AddFavorite
	r.actionHandlers[app.CallbackRemoveBookmark] = h.RemoveBookmark
}

func (r *Router) TextRoute(ctx context.Context, b *bot.Bot, u *models.Update) {
	if u.Message == nil {
		return
	}
	args := Args{
		rawText:  u.Message.Text,
		userID:   u.Message.From.ID,
		msgID:    u.Message.ID,
		userName: u.Message.From.Username,
	}
	current, err := r.f.Current(args.userID)
	if err != nil {
		logger.Error(err)
		return
	}
	switch current.ID {
	case app.StateMainMenu:
		kb := keyboard.MainMenuKeyboard()
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      args.userID,
			Text:        app.MainText,
			ReplyMarkup: kb,
		})
	default:
		kb := keyboard.MainMenuKeyboard()
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      args.userID,
			Text:        app.MainText,
			ReplyMarkup: kb,
		})
	}

}

func (r *Router) CallbackRoute(ctx context.Context, b *bot.Bot, u *models.Update) {
	if u.CallbackQuery == nil {
		return
	}

	logger.Infof("RAW CALLBACK QUERY: %s", u.CallbackQuery.Data)
	args := setupArgsFromCallback(u)

	defer func(b *bot.Bot, ctx context.Context, params *bot.AnswerCallbackQueryParams) {
		_, err := b.AnswerCallbackQuery(ctx, params)
		if err != nil {
			logger.Error(err)
		}
	}(b, ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: u.CallbackQuery.ID,
	})

	if callbackData(args.rawCBData) == app.CallbackBack {
		err := r.f.Back(args.userID)
		if err != nil {
			logger.Error(err)
			return
		}
		r.reRender(ctx, args.userID, u)
		return
	}

	if cb, ok := r.actionHandlers[callbackData(u.CallbackQuery.Data)]; ok {
		err := cb(ctx, args)
		if err != nil {
			logger.Error(err)
		}
		if callbackData(u.CallbackQuery.Data) == app.CallbackRemoveBookmark {
			err := r.f.Back(args.userID)
			if err != nil {
				logger.Error(err)
				return
			}
			r.reRender(ctx, args.userID, u)
			return
		}
		return
	}

	newState := getStateByCallback(callbackData(args.rawCBData))
	logger.Infof("NEW STATE: %v", newState)
	if err := r.f.Transition(args.userID, newState, args); err != nil {
		logger.Error(err)
		return
	}
	if cb, ok := r.handlers[newState]; ok {
		err := cb(ctx, args)
		if err != nil {
			logger.Error(err)
			return
		}
	}
}

func (r *Router) reRender(ctx context.Context, userID int64, u *models.Update) {
	current, _ := r.f.Current(userID)
	if handler, ok := r.handlers[current.ID]; ok {
		if err := handler(ctx, setupArgsFromCache(u, current.Data)); err != nil {
			logger.Error(err)
			return
		}
	}
}
