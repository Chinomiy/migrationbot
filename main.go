package main

import (
	"context"
	"log"
	"migtationbot/application/app"
	"migtationbot/application/country"
	"migtationbot/application/handlers"
	"migtationbot/config"
	"migtationbot/fsm"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	appf := &handlers.Application{}
	opts := []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, appf.TextRouter),
		bot.WithCallbackQueryDataHandler("", bot.MatchTypePrefix, appf.CallbackRouter),
	}
	pool, err := pgxpool.New(ctx, cfg.DBURL)
	countryRepo := country.NewCountryRepository(pool)
	countrySvc := country.NewCountryService(countryRepo)
	appf.CountrySVC = countrySvc
	appf.F = fsm.New(
		app.StateMainMenu,
		map[fsm.StateID]fsm.Callback{
			app.StateCountryMenu:        appf.HandlerCountryMenu,
			app.StateCountry:            appf.HandlerCountryTrip,
			app.StateCountryDetailsMenu: appf.HandlerCountryDetails,
		})
	appf.B, err = bot.New(cfg.TgToken, opts...)
	if err != nil {
		log.Fatal(err)
	}
	appf.B.Start(ctx)
}
