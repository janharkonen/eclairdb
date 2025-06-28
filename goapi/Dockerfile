# Build stage
ARG GO_VERSION=1.24-alpine
FROM golang:${GO_VERSION} AS build
WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -a -installsuffix cgo -o main .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -a -installsuffix cgo -o main .

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