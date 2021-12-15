# syntax=docker/dockerfile:1.3
FROM golang:1.17.1-alpine3.14 as build

RUN sed -i 's/https\:\/\/dl-cdn.alpinelinux.org/http\:\/\/mirror.clarkson.edu/g' /etc/apk/repositories && apk add git --no-cache
WORKDIR /usr/local/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags "-s -w" -o link_remover_tg_bot ./cmd/main.go


FROM alpine:3.14 as link_remover_tg_bot

RUN sed -i 's/https\:\/\/dl-cdn.alpinelinux.org/http\:\/\/mirror.clarkson.edu/g' /etc/apk/repositories && apk add ca-certificates --no-cache

WORKDIR /usr/local/app

ARG TOKEN
ENV TOKEN=$TOKEN

RUN --mount=type=secret,id=API_ENDPOINT \
    export API_ENDPOINT=$(cat /run/secrets/TOKEN) && \
    echo $TOKEN >> .env

COPY --from=build /usr/local/app/link_remover_tg_bot /bin/link_remover_tg_bot

CMD /bin/link_remover_tg_bot
