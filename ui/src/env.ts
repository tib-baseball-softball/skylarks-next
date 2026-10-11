import { ApplicationContext } from "#lib/dp/types/ApplicationContext.ts";
import { LogLevel } from "#lib/dp/types/LogLevel.ts";
import { defineEnvVars } from "@sveltejs/kit/env";

/**
 * UI environment variables according to SvelteKit specification.
 *
 * Note to dev: Remember to update documentation descriptions also in .env.dist
 */
export const variables = defineEnvVars({
  APPLICATION_CONTEXT: {
    public: true,
    schema: (input) => {
      if (
        input !== ApplicationContext.Development &&
        input !== ApplicationContext.Staging &&
        input !== ApplicationContext.Production
      ) {
        throw new Error("Invalid application context given");
      }
      return input;
    },
    static: true,
    description:
      "affects verbosity of logging, non-production environments show visual indicator",
  },
  POCKETBASE_URL: {
    public: true,
    schema: (input) => {
      try {
        if (typeof input === "string") {
          const _ = new URL(input);
        }
      } catch (error) {
        if (error instanceof TypeError) {
          throw new Error(
            "POCKETBASE_URL can be empty, but if set, it needs to be a valid URL.",
          );
        }
      }
      return input;
    },
    static: true,
    description:
      "only used for Vite dev server (production environment sends requests to self, ignored for production builds",
  },
  LOG_LEVEL: {
    public: true,
    schema: (input) => {
      if (input !== LogLevel.Debug && input !== LogLevel.Error) {
        throw new Error("invalid log level, see documentation");
      }
      return input;
    },
    static: true,
    description:
      "affects client-side-logging, for possible values see #lib/dp/types/LogLevel.ts",
  },
  VAPID_PUBLIC_KEY: {
    public: true,
    schema: (input) => {
      if (!input) {
        throw new Error(
          "VAPID public key must be set for push notification functionality.",
        );
      }
      return input;
    },
    static: true,
    description: "VAPID public key for web push functionality",
  },
  TILE_SERVER: {
    public: true,
    schema: (input) => {
      if (!input) {
        throw new Error("No Leaflet tile server given.");
      }
      return input;
    },
    static: true,
    description: "tile server to use for leaflet maps",
  },
});
