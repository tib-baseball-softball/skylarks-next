import { Collection } from "#lib/dp/enum/Collection.ts";
import { toastController } from "#lib/dp/service/ToastController.svelte.ts";
import type { CustomAuthModel } from "#lib/dp/types/ExpandedResponse.ts";
import type { Schema } from "#lib/dp/types/pb-types.ts";
import { browser, dev } from "$app/env";
import { LOG_LEVEL, POCKETBASE_URL } from "$app/env/public";
import { refreshAll } from "$app/navigation";
import { TypedPocketBase } from "@tigawanna/typed-pocketbase";
import type { AuthRecord, RecordAuthResponse, RecordModel } from "pocketbase";
import { LogLevel } from "#lib/dp/types/LogLevel.ts";

/**
 * baseURL is only explicitly defined in dev, use relative URLs for Prod.
 *
 * No server is involved, so a single shared instance is fine.
 */
export const client = new TypedPocketBase<Schema>(
  dev ? POCKETBASE_URL : undefined,
);

/**
 * Global auth state object
 */
class AuthSettings {
  record: AuthRecord = $state(client.authStore.record) as CustomAuthModel;
}

export const authSettings = new AuthSettings();

// seems to work only on explicit logout, not on token expiration
client.authStore.onChange((_token: string, record: AuthRecord) => {
  if (browser) {
    if (dev && LOG_LEVEL === LogLevel.Debug) {
      console.log("Auth Store changed:", record);
    }
    authSettings.record = record;
  }
});

client.authStore.onChange(async (_token, record) => {
  if (browser) {
    if (record) {
      const { email } = record;
      toastController.trigger({
        message: `Signed in as ${email}`,
        background: "preset-filled-success-500",
      });
      await refreshAll();
    } else {
      toastController.trigger({
        message: "Logout successful",
        background: "preset-filled",
      });
      await refreshAll();
    }
  }
});

/**
 * In addition to onChange func, also provide a manual option to refresh the current user.
 */
export async function manualAuthRefresh(): Promise<
  RecordAuthResponse<RecordModel>
> {
  const authData = await client.collection(Collection.Users).authRefresh();
  authSettings.record = authData.record;
  return authData;
}

/** ####################################################### */
