package main

import (
	"context"
	"migtationbot/application/app"
	"migtationbot/application/bookmark"
	bot2 "migtationbot/application/bot"
	"migtationbot/application/country"
	"migtationbot/config"
	"migtationbot/fsm"
	"migtationbot/logger"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	logger.Init()
	cfg, err := config.MustLoad()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	appf := &bot2.Application{}
	opts := []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, appf.TextRouter),
		bot.WithCallbackQueryDataHandler("", bot.MatchTypePrefix, appf.CallbackRouter),
	}
	pool, err := pgxpool.New(ctx, cfg.DBURL)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	// кантри
	countryRepo := country.NewCountryRepository(pool)
	countrySvc := country.NewCountryService(countryRepo)
	// закладки
	bookmarkRepo := bookmark.NewBookmarkRepository(pool)
	bookmarkSvc := bookmark.NewBookMarkService(nil, countrySvc, bookmarkRepo)
	appf.CountrySVC = countrySvc
	appf.BookmarkSVc = bookmarkSvc
	appf.F = fsm.New(
		app.StateMainMenu,
		map[fsm.StateID]fsm.Callback{
			app.StateCountryMenu:        appf.HandlerCountryMenu,
			app.StateCountry:            appf.HandlerCountryTrip,
			app.StateCountryDetailsMenu: appf.HandlerCountryDetails,
			app.StateAccount:            appf.HandlerAccount,
			app.StateFavorite:           appf.HandlerFavorite,
			app.StateBookmarkDetails:    appf.HandlerBookmarkDetails,
		})
	appf.B, err = bot.New(cfg.TgToken, opts...)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	appf.B.Start(ctx)
}
