FROM golang:alpine AS builder

WORKDIR /builder

ENV GO111MODULE=on CGO_ENABLED=0

COPY . .

RUN apk update && apk add --no-cache upx \
    && go build -o /builder/main /builder/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /builder/main /app/main

ENTRYPOINT [ "/app/main" ]