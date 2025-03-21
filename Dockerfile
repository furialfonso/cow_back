# Etapa de construcción
FROM --platform=$BUILDPLATFORM golang:1.20.6-alpine3.18 AS builder
ARG SCOPE
RUN apk add --no-cache git upx && \
    mkdir -p /app
WORKDIR /app

# Copiar y descargar dependencias
COPY ["go.mod", "go.sum", "./"]
RUN go mod download -x

# Copiar el código fuente y compilar
COPY . .
RUN go build -o main api/cmd/main.go && \
    upx --best --ultra-brute main

# Etapa de ejecución
FROM alpine:3.18 AS runner
ARG SCOPE
RUN apk --no-cache add ca-certificates && \
    mkdir -p /app
WORKDIR /app

# Copiar binario y configuración desde la etapa de construcción
COPY --from=builder /app/main .
COPY --from=builder /app/infrastructure/config/prod.yml .

# Configurar el punto de entrada
ENTRYPOINT ["./main"]

# Ejemplo de comandos para construir y ejecutar
# docker build --build-arg SCOPE=local -t furialfonso/cow_project:latest .
# docker run -e PORT=9000 -p 8080:8080 furialfonso/cow_project:latest