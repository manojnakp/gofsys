# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine AS build

WORKDIR /app

COPY . ./

RUN go build -o /pwd

## Deploy
FROM scratch

WORKDIR /

COPY --from=build /pwd /pwd

ENTRYPOINT ["/pwd"]
