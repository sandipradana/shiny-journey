FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOARCH="amd64" \
    GOOS=linux

RUN apk update && apk add --no-cache gcc musl-dev

WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build -o main .

FROM alpine

WORKDIR /app

COPY --from=builder /build/main .

ENTRYPOINT ["./main"]
