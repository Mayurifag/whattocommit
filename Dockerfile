FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o ./application

######################

FROM alpine

WORKDIR /app

RUN addgroup -g 1000 application && adduser -u 1000 -D -G application -s /sbin/nologin application
USER application

COPY --from=builder --chown=1000:1000 /app/application application
COPY --from=builder --chown=1000:1000 /app/commit_messages.txt commit_messages.txt

EXPOSE 8080
CMD ["./application"]
