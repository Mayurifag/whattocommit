FROM golang:alpine

WORKDIR /app

COPY . .

EXPOSE 8080
