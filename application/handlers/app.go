package handlers

import (
	"migtationbot/application/country"
	"migtationbot/fsm"

	"github.com/go-telegram/bot"
)

type Application struct {
	B          *bot.Bot
	F          *fsm.FSM
	CountrySVC *country.CountryService
}
