package config

import (
	"github.com/b2r2/link-remover/pkg"
)

type bot struct {
	token string
}

func newBot() *bot {
	return &bot{
		token: pkg.GetEnv("TOKEN", ""),
	}
}
