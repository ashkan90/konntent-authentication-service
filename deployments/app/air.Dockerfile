FROM golang:1.18-alpine

RUN apk --no-cache add build-base git gcc openssh-client

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY .. .

RUN go install github.com/cosmtrek/air@latest

COPY ./.air.app.toml ./
COPY ./.env ./.env

ENTRYPOINT ["air", "-c", ".air.app.toml"]