package config

import (
	"log"

	"github.com/b2r2/link-remover/pkg"
)

const TOKEN = ".env"

type bot struct {
	token string
}

func newBot() (*bot, error) {
	token, err := pkg.GetEnv(TOKEN)
	log.Println("TOKEN", token)
	if err != nil {
		return nil, err
	}
	return &bot{
		token: token,
	}, nil
}
