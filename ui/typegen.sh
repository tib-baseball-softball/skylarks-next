#!/usr/bin/env bash

source ./.env

npx @tigawanna/typed-pocketbase \
    --type ts \
    --dir src/lib/dp/types
