import { defineEnvVars } from "@sveltejs/kit/env";

export const variables = defineEnvVars({
  PUBLIC_APPLICATION_CONTEXT: { public: true, static: true },
  PUBLIC_POCKETBASE_URL: { public: true, schema: (input) => input ?? "" },
  PUBLIC_LOG_LEVEL: { public: true, schema: (input) => input ?? "" },
  PUBLIC_VAPID_PUBLIC_KEY: { public: true, static: true },
  PUBLIC_TILE_SERVER: { public: true, schema: (input) => input ?? "" },
});
