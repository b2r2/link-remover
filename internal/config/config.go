package config

type config struct {
	bot *bot
}

var conf *config

func Load() error {
	b, err := newBot()
	if err != nil {
		return err
	}

	conf = &config{bot: b}

	return nil
}

func Get() *config {
	return conf
}

func (c *config) GetToken() string {
	return conf.bot.token
}
