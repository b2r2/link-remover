package app

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"
	tele "gopkg.in/tucnak/telebot.v3"
)

type bot struct {
	log *logrus.Logger
	bot *tele.Bot
	m   chan tele.Editable
	rg  *regexp.Regexp
}

func New(log *logrus.Logger, t string) (*bot, error) {
	b, err := tele.NewBot(tele.Settings{
		Token:  t,
		Poller: &tele.LongPoller{Timeout: time.Second * 5},
	})

	if err != nil {
		return nil, err
	}

	r, err := regexp.Compile(`(?:(?:https?|ftp):)?[\w\-?=%.]+\.[\w\-?=%.]+`)
	if err != nil {
		return nil, err
	}

	return &bot{bot: b, log: log, m: make(chan tele.Editable, 100), rg: r}, nil
}

func (b *bot) Start(ctx context.Context) {
	go b.Remover(ctx)
	b.bot.Handle(tele.OnText, func(c tele.Context) error {
		m := c.Message()
		if m.Chat.Username != "nuancesprog" || m.ReplyTo.Chat.Username != "nuancesprog" {
			if b.rg.Match([]byte(m.Text)) {
				b.m <- m
			}
			if len(m.Entities) > 0 {
				for _, e := range m.Entities {
					if b.rg.Match([]byte(e.URL)) {
						b.m <- m
					}
				}
			}
		}
		return nil
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
				fmt.Println(err)
			}
		}
	}
}

func (b *bot) Stop() {
	fmt.Println("exited")
	close(b.m)
	b.bot.Stop()
}
