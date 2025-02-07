FROM golang:1.22.12-alpine AS builder

COPY . /go/src/github.com/irootpro/chat-server
WORKDIR /go/src/github.com/irootpro/chat-server

RUN go mod download
RUN go build -o chat-server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/github.com/irootpro/chat-server/chat-server .

CMD ["./chat-server"]
