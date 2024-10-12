import {client} from "$lib/pocketbase";
import type {LayoutLoad} from "../../../../.svelte-kit/types/src/routes/(protected)/account/$types";
import type {ExpandedTeam} from "$lib/model/ExpandedResponse";
import type {ClubsResponse} from "$lib/model/pb-types";

export const load = (async ({fetch, depends}) => {

  const teams = client.collection("teams").getFullList<ExpandedTeam>({
    expand: "club",
    fetch: fetch
  })

  const clubs = client.collection("clubs").getFullList<ClubsResponse>({
    // TODO: filter for own - API rule might not be enough
    fetch: fetch,
    expand: "admins",
  });

  depends("teams:list")

  return {
    clubs: clubs,
    teams: teams,
  }
}) satisfies LayoutLoad
