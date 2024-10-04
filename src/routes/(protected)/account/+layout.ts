import { client } from "$lib/pocketbase";
import type { LayoutLoad } from "../../../../.svelte-kit/types/src/routes/(protected)/account/$types";
import type { ExpandedTeam } from "$lib/model/ExpandedResponse";

export const load = (async ({ fetch, depends }) => {
  const teams = client.collection("teams").getFullList<ExpandedTeam>({
    expand: "club",
    fetch: fetch
  })

  depends("teams:list")

  return {
    teams: teams
  }
}) satisfies LayoutLoad
