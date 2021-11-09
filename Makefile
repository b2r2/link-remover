BIN := "./bin/link_remover_tg_bot"

IMAGE := "link_remover_tg_bot:dev"

build_ubuntu:
	GOOS=linux GOARCH=amd64 go build -o ${BIN} -v cmd/main.go

build:
	docker build -t $(IMAGE) -f Dockerfile .

run:
	go run -race cmd/main.go


PHONY: build build_ubuntu run