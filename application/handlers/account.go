package handlers

import (
	"context"
	"migtationbot/application/country"
	"migtationbot/fsm"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// TODO
type AccountHandler struct {
	f          *fsm.FSM
	countrySvc *country.CountryService
	// userSvc
}

func NewAccountHandler(
	fsm *fsm.FSM,
	countrySvc *country.CountryService,
) *AccountHandler {
	return &AccountHandler{
		f:          fsm,
		countrySvc: countrySvc,
		//	userSvc
	}
}

const (
	accountMenu = "account"
)

func (с *AccountHandler) Name() fsm.StateID {
	return accountMenu
}
func (c *AccountHandler) Handle(ctx context.Context, b *bot.Bot, u *models.Update) (fsm.StateID, error) {
	return fsm.StateID(""), nil
}
