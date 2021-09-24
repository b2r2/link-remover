package main

import (
	"context"
	"github.com/b2r2/link-remover/internal/app"
	"github.com/b2r2/link-remover/internal/config"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	if err := config.Load(); err != nil {
		log.Fatal("err load config", err)
	}

	cfg := config.Get()

	bot, err := app.New(&logrus.Logger{Out: os.Stderr}, cfg.GetToken())
	if err != nil {
		log.Fatal("err start new bot", err)
	}

	bot.Start(ctx)

	<-ctx.Done()

	bot.Stop()
}
