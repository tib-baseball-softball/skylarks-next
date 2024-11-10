import type {AuthProviderInfo, RecordService} from "pocketbase";
import {save} from "./RecordOperations";

export async function providerLogin(
    provider: AuthProviderInfo,
    authCollection: RecordService,
) {
  const authResponse = await authCollection.authWithOAuth2({
    provider: provider.name,
    createData: {
      // emailVisibility: true,
    },
    query: {
      expand: "club"
    }
  });
  // update user "record" if "meta" has info it doesn't have
  const {meta, record} = authResponse;
  let changes = {} as { [key: string]: any };
  if (!record.name && meta?.name) {
    changes.name = meta.name;
  }
  if (!record.avatar && meta?.avatarUrl) {
    const response = await fetch(meta.avatarUrl);
    if (response.ok) {
      const type = response.headers.get("content-type") ?? "image/jpeg";
      changes.avatar = new File([await response.blob()], "avatar", {type});
    }
  }
  if (Object.keys(changes).length) {
    authResponse.record = await save(authCollection.collectionIdOrName, {
      ...record,
      ...changes,
    });
  }
  return authResponse;
}
