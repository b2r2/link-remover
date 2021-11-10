BIN := "./bin/link_remover_tg_bot"
IMAGE := "link_remover_tg_bot:dev"


build:
	docker build -t $(IMAGE) --secret id=TOKEN,src=secret.env -f Dockerfile .

run:
	docker run $(IMAGE)

PHONY: build run