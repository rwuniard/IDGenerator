# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go get github.com/gin-gonic/gin
RUN go get github.com/joho/godotenv
RUN go mod download


COPY *.go ./
COPY ./algorithms ./algorithms
COPY ./controllers ./controllers
COPY ./initializers ./initializers
COPY .env ./

RUN go build -o ./idgenerator

## Deploy
## This will copy the binary and .env files from the build stage
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/idgenerator ./idgenerator
COPY --from=build /app/.env ./

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["./idgenerator"]