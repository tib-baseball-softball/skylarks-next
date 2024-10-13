import type {PageLoad} from "./$types";
import {client} from "$lib/pocketbase";
import type {ExpandedClub} from "$lib/model/ExpandedResponse";

export const load = (async ({fetch, params, depends}) => {
  const club = await client.collection('clubs').getOne<ExpandedClub>(params.club, {
    expand: "admins",
    fetch: fetch,
  })

  depends("club:single")

  return {
    club: club
  }
}) satisfies PageLoad