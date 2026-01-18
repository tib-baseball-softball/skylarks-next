############################################
# Unified build: Vite (SvelteKit static) + PocketBase (Go)
# Result: single container serving PocketBase and the SPA from /app/pb_public
############################################

# 1) UI build stage
FROM node:24-alpine AS ui-builder
WORKDIR /

RUN npm i -g corepack && corepack enable && corepack prepare pnpm@latest --activate

COPY pnpm-lock.yaml package.json pnpm-workspace.yaml .npmrc* ./
COPY diamond-planner/ui/package.json ./diamond-planner/ui/

RUN pnpm install --frozen-lockfile

ARG BUILD_MODE

COPY . .

RUN pnpm $BUILD_MODE


# 2) Go/PocketBase build stage
FROM golang:1.25-alpine AS go-builder
WORKDIR /backend

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend .

RUN CGO_ENABLED=0 GOOS=linux go build -o diamondplanner


# 3) Final runtime image
FROM alpine:3.20
WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata && update-ca-certificates

COPY --from=go-builder /backend/diamondplanner ./
COPY --from=ui-builder /backend/pb_public ./pb_public

EXPOSE 8090

CMD ["./diamondplanner", "serve", "--http=0.0.0.0:8090"]
