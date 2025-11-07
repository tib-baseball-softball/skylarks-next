import type { LayoutLoad } from "./$types"
import { client } from "$lib/pocketbase/index.svelte.ts"
import { env } from "$env/dynamic/public"
import type { ClubFunction } from "bsm.js"

export const load: LayoutLoad = async ({ fetch }) => {
  const clubOfficials = client.send<ClubFunction[]>("/api/bsm/relay", {
    fetch: fetch,
    query: {
      url: `https://bsm.baseball-softball.de/clubs/${env.PUBLIC_CLUB_ID}/club_functions.json`,
      club: env.PUBLIC_CLUB_ID,
    },
    requestKey: "officials",
  })

  return {
    clubOfficials: clubOfficials,
  }
}
