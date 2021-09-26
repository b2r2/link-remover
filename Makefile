BIN := "./bin/link_remover"

build:
	GOOS=linux GOARCH=amd64 go build -o ${BIN} -v cmd/main.go

run:
	go run -race cmd/main.go


PHONY: build run