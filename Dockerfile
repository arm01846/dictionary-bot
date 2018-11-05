FROM golang:latest AS builder
COPY . /go/src/github.com/arm01846/dictionary-bot
WORKDIR /go/src/github.com/arm01846/dictionary-bot
RUN go get -t ./...
RUN CGO_ENABLED=0 go build -o bot cmd/main.go
RUN chmod +x bot

FROM alpine
RUN apk update && apk add ca-certificates
COPY --from=builder /go/src/github.com/arm01846/dictionary-bot/bot /app/bot
ENTRYPOINT /app/bot
CMD [""]