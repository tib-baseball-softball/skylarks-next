import {client} from "$lib/pocketbase";
import type {LayoutLoad} from "../../../../.svelte-kit/types/src/routes/(protected)/account/$types";
import type {CustomAuthModel, ExpandedClub, ExpandedTeam} from "$lib/model/ExpandedResponse";
import {authModel} from "$lib/pocketbase/Auth";
import {get} from "svelte/store";

export const load = (async ({fetch, depends}) => {
  const model = get(authModel) as unknown as CustomAuthModel;

  const teams = client.collection("teams").getFullList<ExpandedTeam>({
    expand: "club",
    fetch: fetch
  })

  const clubs = client.collection("clubs").getFullList<ExpandedClub>({
    filter: `id ?= "${model.club}"`,
    fetch: fetch,
    expand: "admins",
  });

  depends("teams:list")

  return {
    clubs: clubs,
    teams: teams,
  }
}) satisfies LayoutLoad
