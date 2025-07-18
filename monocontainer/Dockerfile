ARG GO_VERSION=1.24-alpine
FROM golang:${GO_VERSION} AS gobuilder
WORKDIR /app
COPY goapi/go.mod .
COPY goapi/go.sum .
RUN go mod download -x
COPY ./goapi/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM node:24 AS vuebuilder
WORKDIR /app
COPY ./vuefrontend/package.json ./
COPY ./vuefrontend/package-lock.json ./
RUN npm install
COPY ./vuefrontend/ .
RUN npm run build

# This is for deploying the app in a single container
FROM node:24-bookworm-slim
RUN npm install -g http-server
WORKDIR /app

# Install required packages
RUN apt-get update && apt-get install -y \
    nginx \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copy NGINX config
COPY ./monocontainer/nginx.conf /etc/nginx/nginx.conf

# Copy frontend build
COPY --from=vuebuilder /app/dist ./dist
COPY ./vuefrontend/package.json ./
COPY ./vuefrontend/package-lock.json ./

# Copy Go API binary
COPY --from=gobuilder /app/main /usr/local/bin/goapi/main

# Start script
COPY ./monocontainer/start.sh /start.sh
RUN chmod +x /start.sh

EXPOSE 80

CMD ["/start.sh"]