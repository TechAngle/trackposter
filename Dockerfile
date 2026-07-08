# Build
FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git

WORKDIR /build
COPY go.mod go.sum /build/
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o trackposter -ldflags="-s -w" ./cmd/trackposter || (echo "BUILD FAILED" && exit 1)

# Run binary
FROM alpine:latest

RUN apk add --no-cache yt-dlp ffmpeg

WORKDIR /app
COPY --from=builder /build/trackposter /app/trackposter

CMD [ "/app/trackposter" ]
