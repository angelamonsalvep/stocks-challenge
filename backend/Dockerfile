# Imagen base
FROM golang:1.25-alpine

WORKDIR /app

# Copiar dependencias y código
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar
RUN go build -o main ./cmd

EXPOSE 8081
CMD ["./main"]
