package main

import (
	"context"
	"migtationbot/fsm"
	"migtationbot/internal/app"
	"migtationbot/internal/bookmark"
	b "migtationbot/internal/bot"
	"migtationbot/internal/config"
	"migtationbot/internal/country"
	"migtationbot/internal/user"
	"migtationbot/logger"
	"os"
	"os/signal"

	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	trmanager "github.com/avito-tech/go-transaction-manager/trm/v2/manager"
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

	appf := &b.Application{}
	opts := []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, appf.TextRouter),
		bot.WithCallbackQueryDataHandler("", bot.MatchTypePrefix, appf.CallbackRouter),
	}

	pool, err := pgxpool.New(ctx, cfg.DBURL)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	trManager, err := trmanager.New(trmpgx.NewDefaultFactory(pool))
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	// страны
	countryRepo := country.NewCountryRepository(pool)
	countrySvc := country.NewCountryService(countryRepo, trManager)
	appf.CountrySVC = countrySvc

	// закладки
	bookmarkRepo := bookmark.NewBookmarkRepository(pool)
	bookmarkSvc := bookmark.NewBookMarkService(nil, countrySvc, bookmarkRepo)
	appf.BookmarkSVc = bookmarkSvc

	//юзер
	userRepo := user.NewUserRepository(pool)
	userSvc := user.NewUserService(userRepo, trManager)
	appf.UserSVC = userSvc

	//регистрируем хенделры
	appf.F = fsm.New(
		app.StateMainMenu,
		map[fsm.StateID]fsm.Callback{
			app.StateCountryMenu:        appf.HandlerCountryMenu,
			app.StateCountry:            appf.HandlerCountryTrip,
			app.StateCountryDetailsMenu: appf.HandlerCountryDetails,
			app.StateAccount:            appf.HandlerAccount,
			app.StateFavorite:           appf.HandlerFavorite,
			app.StateBookmarkDetails:    appf.HandlerBookmarkDetails,
			app.StateManagerMenu:        appf.HandlerManagerMenu,
		})

	appf.B, err = bot.New(cfg.TgToken, opts...)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	appf.B.Start(ctx)
}
