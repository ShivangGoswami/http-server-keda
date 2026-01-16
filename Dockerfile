# ---- Build stage ----
FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o http-server .

# ---- Run stage ----
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/http-server .
EXPOSE 8080
CMD ["./http-server"]
