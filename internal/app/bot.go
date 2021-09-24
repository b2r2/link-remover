package app

import (
	"context"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

type bot struct {
	log *logrus.Logger
	bot *tb.Bot
	m   chan tb.Editable
	rg  *regexp.Regexp
}

func New(log *logrus.Logger, t string) (*bot, error) {
	b, err := tb.NewBot(tb.Settings{
		Token: t,
		Poller: &tb.LongPoller{
			Timeout: 10 * time.Second,
		},
	})

	if err != nil {
		return nil, err
	}

	r, err := regexp.Compile(`(?:(?:https?|ftp)://)?[\w\-?=%.]+\.[\w\-?=%.]+`)
	if err != nil {
		return nil, err
	}

	return &bot{bot: b, log: log, m: make(chan tb.Editable, 100), rg: r}, nil
}

func (b *bot) Start(ctx context.Context) {
	go b.Remover(ctx)
	b.bot.Handle(tb.OnText, func(m *tb.Message) {
		if m.Chat.Username != "nuancesprog" || m.ReplyTo.Chat.Username != "nuancesprog" {
			if b.rg.Match([]byte(m.Text)) {
				b.m <- m
			}
		}
	})

	go b.bot.Start()

	<-ctx.Done()
}

func (b *bot) Remover(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case m := <-b.m:
			if err := b.bot.Delete(m); err != nil {
				b.log.Errorln(err)
			}
		}
	}
}

func (b *bot) Stop() {
	close(b.m)
	b.bot.Stop()
}
