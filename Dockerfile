FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o broker cmd/server/main.go

FROM scratch
COPY --from=builder /app/broker /broker
COPY .env .env
EXPOSE 8080
ENTRYPOINT ["/broker"]
