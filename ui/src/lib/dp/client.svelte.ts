import {TypedPocketBase} from "@tigawanna/typed-pocketbase";
import type {AuthRecord, RecordAuthResponse, RecordModel} from "pocketbase";
import {browser, dev} from "$app/environment";
import {invalidateAll} from "$app/navigation";
import {env} from "$env/dynamic/public";
import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
import type {Schema} from "$lib/dp/types/pb-types.ts";
import {Collection} from "$lib/dp/enum/Collection.ts";

/**
 * baseURL is only explicitly defined in dev, use relative URLs for Prod.
 *
 * No server is involved, so a single shared instance is fine.
 */
export const client = new TypedPocketBase<Schema>(dev ? env.PUBLIC_POCKETBASE_URL : undefined);

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
    if (dev && env.PUBLIC_LOG_LEVEL === "DEBUG") {
      console.log("Auth Store changed:", record);
    }
    authSettings.record = record;
  }
});

client.authStore.onChange(async (_token, record) => {
  if (browser) {
    if (record) {
      const {email} = record;
      toastController.trigger({
        message: `Signed in as ${email}`,
        background: "preset-filled-success-500",
      });
      await invalidateAll();
    } else {
      toastController.trigger({
        message: "Logout successful",
        background: "preset-filled",
      });
      await invalidateAll();
    }
  }
});

/**
 * In addition to onChange func, also provide a manual option to refresh the current user.
 */
export async function manualAuthRefresh(): Promise<RecordAuthResponse<RecordModel>> {
  const authData = await client.collection(Collection.Users).authRefresh();
  authSettings.record = authData.record;
  return authData;
}

/** ####################################################### */
