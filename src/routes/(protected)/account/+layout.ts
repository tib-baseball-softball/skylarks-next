import { client } from "$lib/pocketbase";
import type { LayoutLoad } from "../../../../.svelte-kit/types/src/routes/(protected)/account/$types";
import type { ExpandedTeam } from "$lib/model/ExpandedResponse";

export const load = (async ({ fetch }) => {
  const teams = client.collection("teams").getFullList<ExpandedTeam>({
    expand: "club",
    fetch: fetch
  })

  return {
    teams: teams
  }
}) satisfies LayoutLoad
