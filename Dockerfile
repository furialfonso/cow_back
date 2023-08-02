FROM --platform=$BUILDPLATFORM golang:1.20.6-alpine3.18 AS builder
RUN apk add --no-cache git upx
WORKDIR /app
COPY ["go.mod","go.sum", "./"]
RUN go mod download -x
COPY . .
RUN go build api/main.go

#upload compilance
FROM alpine:3.18 AS runner
RUN apk update
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/pkg/config/prod.yml .

ENTRYPOINT [ "./main" ]
# docker build -t furialfonso/cow_project:latest .
# docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t furialfonso/cow_project:latest --push .
# docker run -e PORT=9000 -p 8080:8080 app-pro:v1