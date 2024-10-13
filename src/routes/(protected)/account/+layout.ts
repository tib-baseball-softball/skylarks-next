import {client} from "$lib/pocketbase";
import type {LayoutLoad} from "../../../../.svelte-kit/types/src/routes/(protected)/account/$types";
import type {CustomAuthModel, ExpandedClub, ExpandedTeam} from "$lib/model/ExpandedResponse";
import {authModel} from "$lib/pocketbase/Auth";
import {get} from "svelte/store";

export const load = (async ({fetch, depends}) => {
  const model = get(authModel) as unknown as CustomAuthModel;

  const teams = client.collection("teams").getFullList<ExpandedTeam>({
    filter: `"${model.teams}" ?~ id`,
    expand: "club",
    fetch: fetch
  })

  const clubs = client.collection("clubs").getFullList<ExpandedClub>({
    filter: `"${model.club}" ?~ id`,
    fetch: fetch,
    expand: "admins",
  });

  depends("teams:list")

  return {
    clubs: clubs,
    teams: teams,
  }
}) satisfies LayoutLoad
