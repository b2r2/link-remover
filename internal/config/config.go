package config

import "github.com/sirupsen/logrus"

type config struct {
	bot    *bot
	logger *logrus.Logger
}

var conf *config

func Load() error {
	b, err := newBot()
	if err != nil {
		return err
	}

	conf = &config{bot: b, logger: logrus.New()}

	return nil
}

func Get() *config {
	return conf
}

func (c *config) GetToken() string {
	return conf.bot.token
}

func (c *config) GetLogger() *logrus.Logger {
	return conf.logger
}
