############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/binozoworks/Zyzz-Motivation-Booster/

COPY . .
WORKDIR main
# Fetch dependencies.
# Using go get.

RUN go get -d -v
# Build the binary.


RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/zyzzmotivationbooster

############################
# STEP 2 build a small image
############################
FROM alpine:latest
# Copy our static executable.
COPY --from=builder /go/bin/zyzzmotivationbooster /go/bin/zyzzmotivationbooster
# Run the hello binary.
ENTRYPOINT ["/go/bin/zyzzmotivationbooster"]