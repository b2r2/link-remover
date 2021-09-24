package config

type config struct {
	bot *bot
}

var conf *config

func Load() error {
	b := newBot()

	cfg := &config{bot: b}
	conf = cfg

	return nil
}

func Get() *config {
	return conf
}

func (c *config) GetToken() string {
	return conf.bot.token
}
