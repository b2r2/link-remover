package config

import (
	"fmt"

	"github.com/b2r2/link-remover/pkg"
)

type bot struct {
	token string
}

func newBot() *bot {
	fmt.Println("token:", pkg.GetEnv("TOKEN", ""))
	return &bot{
		token: pkg.GetEnv("TOKEN", ""),
	}
}
