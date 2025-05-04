# Use the official Golang image as a build stage
FROM golang:1.24-alpine AS builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main cmd/main.go

CMD ["/app/main"]