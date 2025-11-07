import type { LayoutLoad } from "./$types"
import { client } from "$lib/pocketbase/index.svelte.ts"
import type { Club } from "bsm.js"
import { env } from "$env/dynamic/public"

export const load: LayoutLoad = async ({ fetch }) => {
  const bsmClubData = await client.send<Club>("/api/bsm/relay", {
    fetch: fetch,
    query: {
      url: `https://bsm.baseball-softball.de/clubs/${env.PUBLIC_CLUB_ID}.json`,
      club: env.PUBLIC_CLUB_ID,
    },
    requestKey: env.PUBLIC_CLUB_ID,
  })

  return {
    bsmClubData: bsmClubData,
  }
}
