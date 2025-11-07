import type { LayoutLoad } from "./$types"
import { client } from "$lib/pocketbase/index.svelte.ts"
import type { Field } from "bsm.js"
import { env } from "$env/dynamic/public"

export const load: LayoutLoad = async ({ fetch }) => {
  const fields = client.send<Field[]>("/api/bsm/relay", {
    fetch: fetch,
    query: {
      url: `https://bsm.baseball-softball.de/clubs/${env.PUBLIC_CLUB_ID}/fields.json`,
      club: env.PUBLIC_CLUB_ID,
    },
    requestKey: "ballparks",
  })

  return {
    fields: fields,
  }
}
