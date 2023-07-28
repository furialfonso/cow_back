#build stage
FROM golang:1.19-alpine3.18 AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go get -d -v ./...

ENV APPLICATION_PACKAGE=./cmd/api
ENV SCOPE=local
ENV PORT=8080
ENV CONFIG_DIR=/app/pkg/config,
ENV GIN_MODE=release

RUN go build -o /go/bin/app -v ./...

ENTRYPOINT [ "./app" ]

#final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT /app
# LABEL Name=dockergoproject Version=0.0.1
# EXPOSE 8080
