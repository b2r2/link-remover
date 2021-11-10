# syntax=docker/dockerfile:1.2
FROM golang:1.17.1-alpine3.14 as build

RUN sed -i 's/https\:\/\/dl-cdn.alpinelinux.org/http\:\/\/mirror.clarkson.edu/g' /etc/apk/repositories && apk add git --no-cache
WORKDIR /usr/local/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN --mount=type=secret,id=TOKEN \
    export TOKEN=$(cat /run/secrets/TOKEN)

RUN go build -ldflags "-s -w" -o link_remover_tg_bot ./cmd/main.go

FROM alpine:3.14 as link_remover_tg_bot

RUN sed -i 's/https\:\/\/dl-cdn.alpinelinux.org/http\:\/\/mirror.clarkson.edu/g' /etc/apk/repositories && apk add ca-certificates --no-cache

RUN --mount=type=secret,id=TOKEN,dst=/home/
#    export TOKEN=$(cat /.env)

WORKDIR /usr/local/app
COPY --from=build /usr/local/app/link_remover_tg_bot /bin/link_remover_tg_bot

CMD /bin/link_remover_tg_bot
