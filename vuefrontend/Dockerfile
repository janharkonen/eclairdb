FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM node:20-alpine AS production
RUN npm install -g http-server
WORKDIR /app
COPY --from=builder /app/dist ./dist
COPY package*.json ./
CMD ["http-server", "dist"]