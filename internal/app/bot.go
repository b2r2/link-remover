package app

import (
	"context"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode"

	"mvdan.cc/xurls/v2"

	"github.com/sirupsen/logrus"
	tele "gopkg.in/tucnak/telebot.v3"
)

type bot struct {
	log *logrus.Logger
	bot *tele.Bot
	rg  *regexp.Regexp
	sync.RWMutex
	m            chan message
	removedLinks []user
}

type user struct {
	removedAt string
	url       string
}

type message struct {
	m   tele.Message
	url string
}

const domain = "nuancesprog"

func New(log *logrus.Logger, t string) (*bot, error) {
	b, err := tele.NewBot(tele.Settings{
		Token:  t,
		Poller: &tele.LongPoller{Timeout: time.Second * 8},
	})

	if err != nil {
		return nil, err
	}

	return &bot{bot: b, log: log, m: make(chan message, 100), rg: xurls.Relaxed(), removedLinks: make([]user, 0, 10_000)}, nil
}

func (b *bot) Start(ctx context.Context) {
	b.bot.Handle(tele.OnText, func(c tele.Context) error {
		b.checkMessage(c)
		return nil
	})

	b.bot.Handle(tele.OnEdited, func(c tele.Context) error {
		b.checkMessage(c)
		return nil
	})

	go b.bot.Start()

	<-ctx.Done()
}

func (b *bot) Remover(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case m := <-b.m:
				if err := b.bot.Delete(&m.m); err != nil {
					b.log.Errorln(err)
				} else {
					u := user{
						removedAt: time.Now().String(),
						url:       m.url,
					}

					b.RLock()
					b.removedLinks = append(b.removedLinks, u)
					b.RUnlock()
				}
			}
		}
	}()
}

func (b *bot) Stop() {
	b.bot.Stop()
	b.Lock()
	close(b.m)
	b.log.Println(b.removedLinks)
	b.Unlock()
}

func (b *bot) checkMessage(c tele.Context) {
	m := c.Message()

	m.Text = strings.ToLower(strings.TrimFunc(c.Message().Text, func(r rune) bool {
		if !unicode.IsSpace(r) || !unicode.IsMark(r) {
			return true
		}
		return false
	}))
	if m.Chat.Username != domain || m.ReplyTo.Chat.Username != domain {
		if len(b.rg.FindAllString(m.Text, -1)) > 0 {
			b.push(m, m.Text)
		}
		if len(m.Entities) > 0 {
			for _, e := range m.Entities {
				if _, err := url.ParseRequestURI(e.URL); err == nil {
					b.push(m, e.URL)
					break
				} else {
					b.log.Println(err)
				}
			}
		}
		if strings.Contains(m.Text, "@") {
			b.push(m, m.Text)
		}
	}
}

func (b *bot) push(m *tele.Message, url string) {
	var msg message
	msg.m = *m
	msg.url = url
	b.m <- msg
}
