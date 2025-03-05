import {client} from "$lib/pocketbase/index.svelte";
import type {LayoutLoad} from "../../../../.svelte-kit/types/src/routes/(protected)/account/$types";
import type {CustomAuthModel, ExpandedClub, ExpandedTeam} from "$lib/model/ExpandedResponse";

export const load = (async ({fetch, depends, parent}) => {
  const model = client.authStore.record as CustomAuthModel;
  const data = await parent()

  if (!client.authStore.isValid) {
    return
  }

  let clubs = data.clubs
  let teams = data.teams

  if (!teams) {
    teams = await client.collection("teams").getFullList<ExpandedTeam>({
      filter: `"${model?.teams}" ?~ id`,
      expand: "club,admins",
      fetch: fetch,
      sort: "+name",
    })
  }

  if (!clubs) {
    clubs = await client.collection("clubs").getFullList<ExpandedClub>({
      filter: `"${model?.club}" ?~ id`,
      fetch: fetch,
      expand: "admins",
    });
  }

  depends("teams:list")

  return {
    clubs: clubs,
    teams: teams,
  }
}) satisfies LayoutLoad
