FROM golang:1.20.4-alpine3.18 AS builder

RUN apk add --no-cache git upx

WORKDIR /app

COPY ["go.mod","go.sum", "./"]

RUN go mod download -x

COPY . .

RUN go build api/main.go

#final stage
FROM alpine:3.18

LABEL Name=dockerization

RUN apk update

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT [ "./main" ]
# docker build . -t app-pro:v1
# docker run -e PORT=9000 -p 8080:8080 app-pro:v1