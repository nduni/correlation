FROM golang:alpine as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./weather/weather-acceptor /app/weather/weather-acceptor

COPY ./common /app/common

WORKDIR /app/weather/weather-acceptor 

RUN go build -o main .

FROM alpine:latest

RUN apk add --no-cache tzdata

COPY --from=builder /app/weather/weather-acceptor/main .

COPY --from=builder /app/weather/weather-acceptor/app/resources/config ./resources/config


ENTRYPOINT [ "/main" ]