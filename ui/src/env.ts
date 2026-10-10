import { defineEnvVars } from '@sveltejs/kit/env';

// @migration-task Review usage of dynamic environment variables. They fall back to the empty string if not present, which may not be what you want.
export const variables = defineEnvVars({
	PUBLIC_APPLICATION_CONTEXT: { public: true, static: true },
	PUBLIC_POCKETBASE_URL: { public: true, schema: (input) => input ?? '' },
	PUBLIC_LOG_LEVEL: { public: true, schema: (input) => input ?? '' },
	PUBLIC_VAPID_PUBLIC_KEY: { public: true, static: true },
	PUBLIC_TILE_SERVER: { public: true, schema: (input) => input ?? '' }
});
