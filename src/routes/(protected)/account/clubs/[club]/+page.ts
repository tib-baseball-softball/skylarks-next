import type {PageLoad} from "./$types";
import {client} from "$lib/pocketbase/index.svelte";
import type {ExpandedClub, ExpandedTeam, ExpandedUniformSet} from "$lib/model/ExpandedResponse";
import type { LocationsResponse, UniformsetsResponse } from "$lib/model/pb-types.js";

export const load = (async ({fetch, params, depends}) => {

  const club = await client.collection('clubs').getOne<ExpandedClub>(params.club, {
    expand: "admins",
    fetch: fetch,
  })

  const teams = await client.collection("teams").getFullList<ExpandedTeam>({
    filter: `club.id = "${club.id}"`,
    fetch: fetch,
    expand: "club,admins",
    sort: "+name",
  })

  const uniformSets = await client.collection("uniformsets").getFullList<ExpandedUniformSet>({
    filter: `club.id = "${club.id}"`,
    fetch: fetch,
    expand: "club",
  })

  const locations = await client.collection("locations").getFullList<LocationsResponse>({
    filter: `club.id = "${club.id}"`,
    fetch: fetch,
  })

  depends("club:single")

  return {
    club: club,
    teams: teams,
    uniformSets: uniformSets,
    locations: locations,
  }
}) satisfies PageLoad