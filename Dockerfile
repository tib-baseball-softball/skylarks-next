############################################
# Unified build: Vite (SvelteKit static) + PocketBase (Go)
# Result: single container serving PocketBase and the SPA from /app/pb_public
############################################

# 1) UI build stage
FROM node:24-alpine AS ui-builder
WORKDIR /

RUN npm i -g corepack && corepack enable && corepack prepare pnpm@latest --activate

COPY ui/pnpm-lock.yaml ui/package.json ui/pnpm-workspace.yaml ui/.npmrc* ./

RUN pnpm install --frozen-lockfile

ARG BUILD_MODE

COPY . .

RUN pnpm $BUILD_MODE


# 2) Go/PocketBase build stage
FROM golang:1.25-alpine AS go-builder
WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY backend .

RUN CGO_ENABLED=0 GOOS=linux go build -o diamondplanner


# 3) Final runtime image
FROM alpine:3.20
WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata && update-ca-certificates

COPY --from=go-builder /diamondplanner ./
COPY --from=ui-builder /pb_public ./pb_public

EXPOSE 8090

CMD ["./diamondplanner", "serve", "--http=0.0.0.0:8090"]
