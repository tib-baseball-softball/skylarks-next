import type { PageLoad } from "./$types"
import { client } from "$lib/pocketbase/index.svelte"
import type { LocationsResponse } from "$lib/model/pb-types.js"
import type { ExpandedClub } from "$lib/model/ExpandedResponse.ts"

export const load = (async ({ fetch, params, depends }) => {
  const club = await client.collection("clubs").getOne<ExpandedClub>(params.club, {
    expand: "admins",
    fetch: fetch,
  })

  const locations = await client.collection("locations").getFullList<LocationsResponse>({
    filter: `club.id = "${params.club}"`,
    fetch: fetch,
  })

  depends("club:locations")

  return {
    club: club,
    locations: locations,
  }
}) satisfies PageLoad
