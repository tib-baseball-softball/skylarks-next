# Diamond Planner Frontend

SvelteKit-based SPA with `adapter-static` (no server-side logic). Communicates with Pocketbase API server via REST.

## Development

```bash
cp .env.dist
pnpm dev
```

---

Builds use Vite build modes to inject correct env vars into bundle: 

https://vite.dev/guide/env-and-mode#env-variables-and-modes

Env vars that need to be set are the same as in development (first block in `.env.dist`)

## Staging

```bash
cp .env.dist .env.staging
pnpm build:staging
```

## Production

```bash
cp .env.dist .env.production
pnpm build:prod
```
