FROM golang:1.23.1-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .


FROM alpine:latest


COPY --from=builder /app/server /server


EXPOSE 8000

CMD ["/server"]