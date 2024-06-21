# Build stage
FROM docker.io/golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
WORKDIR /go/src/app/server  
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v .

# Final stage
FROM docker.io/alpine:latest

# Instalação do MongoDB e outras dependências necessárias
RUN apk add --no-cache mongodb-tools ca-certificates

# Copia o executável da aplicação Go compilada a partir do builder stage
COPY --from=builder /go/bin/app /app

ENV MONGODB_URL=$MONGODB_URL
ENV MONGODB_USERNAME=$MONGODB_USERNAME
ENV MONGODB_PASSWORD=$MONGODB_PASSWORD
ENV MONGODB_GENEALOGY_DB=$MONGODB_GENEALOGY_DB

# Comando de entrada para iniciar a aplicação
ENTRYPOINT ["/app"]

# Exposição da porta em que a aplicação irá escutar
EXPOSE 8080
