# Build stage
ARG GO_VERSION=1.24-alpine
FROM golang:${GO_VERSION} AS build
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download -x
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache curl
RUN adduser -D goapiuser
USER goapiuser
COPY --from=build /app/main .
HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD curl -f http://localhost:8081/ping || exit 1
CMD ["./main"]