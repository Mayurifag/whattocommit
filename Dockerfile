FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download && \
    echo "nobody:x:65534:65534:Nobody:/:" > /passwd_nobody

ADD . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags="-w -s" -o whattocommit main.go

######################

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch

WORKDIR /app

COPY --from=builder /app/whattocommit /app/whattocommit
COPY --from=builder /app/commit_messages.txt /app/commit_messages.txt
COPY --from=builder /passwd_nobody /etc/passwd

USER nobody

EXPOSE 8080
ENTRYPOINT ["/app/whattocommit"]
