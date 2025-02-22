import {env} from "$env/dynamic/public";
import type {Schema} from "$lib/model/pb-types";
import {TypedPocketBase} from "@tigawanna/typed-pocketbase";
import type {AuthRecord, RecordAuthResponse, RecordModel} from "pocketbase";
import {browser, dev} from "$app/environment";
import type {CustomAuthModel} from "$lib/model/ExpandedResponse";

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

client.authStore.onChange((_token: string, record: AuthRecord) => {
  if (browser) {
    if (dev && env.PUBLIC_LOG_LEVEL === "DEBUG") {
      console.log("Auth Store changed:", record);
    }
    authSettings.record = record;
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