package bot

import (
	"migtationbot/fsm"
	"migtationbot/internal/bookmark"
	"migtationbot/internal/country"
	"migtationbot/internal/user"

	"github.com/go-telegram/bot"
)

type Application struct {
	B           *bot.Bot
	F           *fsm.FSM
	CountrySVC  country.CountryService
	BookmarkSVc bookmark.BookmarkService
	UserSVC     user.UserService
}
