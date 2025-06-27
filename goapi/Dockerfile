# Build stage
FROM golang:tip-bullseye AS builder
WORKDIR /app
COPY go.mod .
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOEXPERIMENT=jsonv2 go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache curl
RUN adduser -D goapiuser
USER goapiuser
COPY --from=builder /app/main .
HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD curl -f http://localhost:8080/ping || exit 1
CMD ["./main"]