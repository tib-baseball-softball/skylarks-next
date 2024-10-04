import { PUBLIC_POCKETBASE_URL } from "$env/static/public";
import type { Schema } from "$lib/model/pb-types";
import { TypedPocketBase } from "typed-pocketbase";

// necessary because: https://github.com/pocketbase/pocketbase/discussions/5313
export function createPocketBaseInstance() {
  return new TypedPocketBase<Schema>(PUBLIC_POCKETBASE_URL);
}

export const client = createPocketBaseInstance();
