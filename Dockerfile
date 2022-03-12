FROM golang:1.17-alpine

RUN apk add --no-cache bash

WORKDIR /home/app
