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

	pool, err := pgxpool.New(ctx, cfg.DBURL)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	//FSM
	f := fsm.New(app.StateMainMenu)

	router := b.NewUpdateHandler(f)
	opts := []bot.Option{
		bot.WithMessageTextHandler("", bot.MatchTypePrefix, router.UpdateTextHandler),
		bot.WithCallbackQueryDataHandler("", bot.MatchTypePrefix, router.UpdateCallbackHandler),
	}

	trManager, err := trmanager.New(trmpgx.NewDefaultFactory(pool))
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	// страны
	countryRepo := country.NewCountryRepository(pool)
	countrySvc := country.NewCountryService(countryRepo, trManager)

	// закладки
	bookmarkRepo := bookmark.NewBookmarkRepository(pool)
	bookmarkSvc := bookmark.NewBookMarkService(nil, countrySvc, bookmarkRepo)

	//юзер
	userRepo := user.NewUserRepository(pool)
	userSvc := user.NewUserService(userRepo, trManager)

	h := b.NewHandler(
		userSvc,
		bookmarkSvc,
		countrySvc,
	)
	h.B, err = bot.New(cfg.TgToken, opts...)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	router.RegisterHandler(h)

	h.B.Start(ctx)
}
