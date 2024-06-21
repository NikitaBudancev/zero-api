FROM golang:1.22-alpine AS builder

RUN apk --no-cache add bash make git gcc musl-dev

RUN go install github.com/cespare/reflex@latest

WORKDIR /usr/local/src

COPY ../../go.mod go.sum ./

RUN go mod download

COPY ../.. .

CMD ["reflex", "-r", "\\.go$", "-s", "--", "go", "run", "cmd/app/main.go"]
