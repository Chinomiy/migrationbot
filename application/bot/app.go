package bot

import (
	"migtationbot/application/bookmark"
	"migtationbot/application/country"
	"migtationbot/application/user"
	"migtationbot/fsm"

	"github.com/go-telegram/bot"
)

type Application struct {
	B           *bot.Bot
	F           *fsm.FSM
	CountrySVC  country.CountryService
	BookmarkSVc bookmark.BookmarkService
	UserSVC     user.UserService
}
