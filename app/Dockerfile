# syntax=docker/dockerfile:1

FROM golang:1.23.1-alpine3.20

WORKDIR /app

COPY . .

RUN go mod tidy && \
	go install honnef.co/go/tools/cmd/staticcheck@latest && \
	go install github.com/rubenv/sql-migrate/...@latest && \
	go install github.com/volatiletech/sqlboiler/v4@latest && \
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

RUN apk add --update make
