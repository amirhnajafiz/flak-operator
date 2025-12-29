# Use official go image
FROM golang:1.25-alpine3.21 AS build

WORKDIR /usr/local/app

COPY go.mod go.sum ./

# download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# copy all files
COPY . .

# build the go app
RUN GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build -o flak

# Minimal alpine image
FROM alpine:3.21

COPY --from=build /usr/local/app/flak /usr/local/bin/flak

ENTRYPOINT ["/usr/local/bin/flak"]
