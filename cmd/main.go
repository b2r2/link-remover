package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/b2r2/link-remover/internal/app"
	"github.com/b2r2/link-remover/internal/config"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	if err := config.Load(); err != nil {
		log.Fatal("err load config", err)
	}

	cfg := config.Get()

	bot, err := app.New(cfg.GetLogger(), cfg.GetToken())
	if err != nil {
		log.Fatal("error instance application: ", err)
	}

	bot.Remover(ctx)
	bot.Start(ctx)

	<-ctx.Done()

	bot.Stop()
}
