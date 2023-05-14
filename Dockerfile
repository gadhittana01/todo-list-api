#!/bin/bash
FROM golang:alpine

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /app/cmd/todolist-http ./cmd/todolist-http
RUN apk add && apk add make

EXPOSE 3030

CMD [ "sh", "-c", "/app/cmd/todolist-http/todolist-http"]