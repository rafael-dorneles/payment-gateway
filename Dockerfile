# Estágio 1: Compilação
# Usando 1.25 para compatibilidade com pgx v5.9.1
FROM golang:1.25-alpine AS builder

# Instala dependências necessárias para compilação (se houver CGO ou ferramentas de rede)
RUN apk add --no-cache git

WORKDIR /app

# Copia arquivos de dependências e baixa os módulos
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código-fonte, incluindo a pasta de migrations
COPY . .

# Compila o binário com flags para reduzir o tamanho e desabilitar CGO (estático)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o payment-api ./cmd/api/main.go

# Estágio 2: Execução (Imagem final ultra leve)
FROM alpine:latest

# Instala certificados CA (necessário para chamadas HTTPS/APIs externas)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 1. Copia o binário do estágio de build
COPY --from=builder /app/payment-api .

# 2. Copia a pasta de migrations (Essencial para não dar erro de "no such file")
# Certifique-se de que o caminho "migrations" condiz com a sua estrutura de pastas
COPY --from=builder /app/migrations ./migrations

# Expõe a porta da API
EXPOSE 8080

# Comando para rodar
CMD ["./payment-api"]