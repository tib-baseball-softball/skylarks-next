import {env} from "$env/dynamic/public";
import type {Schema} from "$lib/model/pb-types";
import {TypedPocketBase} from "@tigawanna/typed-pocketbase";
import type {AuthRecord, RecordAuthResponse, RecordModel} from "pocketbase";
import {browser, dev} from "$app/environment";
import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
import {toastController} from "$lib/service/ToastController.svelte.ts";
import {invalidateAll} from "$app/navigation";

// necessary because: https://github.com/pocketbase/pocketbase/discussions/5313
export function createPocketBaseInstance() {
  return new TypedPocketBase<Schema>(env.PUBLIC_POCKETBASE_URL);
}

/** ######################################################
 *
 * Instance of the PB client in an app-wide scope.
 *
 * _Reminder_: This absolutely has to stay out of all server-side code.
 */
export let client = createPocketBaseInstance();

/**
 * Global auth state object
 */
class AuthSettings {
  record: AuthRecord = $state(client.authStore.record) as CustomAuthModel;
}

export let authSettings = new AuthSettings();

// seems to work only on explicit logout, not on token expiration
client.authStore.onChange((_token: string, record: AuthRecord) => {
  if (browser) {
    if (dev && env.PUBLIC_LOG_LEVEL === "DEBUG") {
      console.log("Auth Store changed:", record);
    }
    authSettings.record = record;
  }
});

client.authStore.onChange((_token, record) => {
  if (browser) {
    if (record) {
      const {email} = record;
      toastController.trigger({
        message: `Signed in as ${email}`,
        background: "preset-filled-success-500",
      });
      invalidateAll();
    } else {
      toastController.trigger({
        message: "Logout successful",
        background: "preset-filled",
      });
      invalidateAll();
    }
  }
});

/**
 * In addition to onChange func, also provide manual option to refresh current user.
 */
export async function manualAuthRefresh(): Promise<RecordAuthResponse<RecordModel>> {
  const authData = await client.collection('users').authRefresh();
  authSettings.record = authData.record;
  return authData;
}

/** ####################################################### */