FROM golang:1.20.4-bullseye as development

WORKDIR /app
VOLUME /app
ADD . /app

ENV GO_ENV=development

RUN apt-get update && apt-get -y install default-mysql-client
RUN go install github.com/cosmtrek/air@latest
