package config

type bot struct {
	token string
}

func newBot() *bot {
	return &bot{
		token: "2022598237:AAETE8MNx4qQ71CdsaWtYtvaq0UVyzHBqsw",
		//token: "455278361:AAFSYcbmNvtshujXKU8oxjIxh3XxPyc_pvo", //debug
	}
}
