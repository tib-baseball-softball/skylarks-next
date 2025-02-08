FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o pocketbase

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/pocketbase ./

EXPOSE 8090

CMD ["./pocketbase", "serve", "--http=0.0.0.0:8090"]