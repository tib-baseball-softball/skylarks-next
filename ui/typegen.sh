#!/usr/bin/env bash
# use `pnpm generate:types` instead, kept for reference

source ./.env

npx @tigawanna/typed-pocketbase \
    --type ts \
    --dir src/lib/dp/types
