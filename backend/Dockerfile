# Première étape : construire l'application
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o websocket-server cmd/main.go

# Deuxième étape : créer une image légère
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/ .

CMD ./websocket-server