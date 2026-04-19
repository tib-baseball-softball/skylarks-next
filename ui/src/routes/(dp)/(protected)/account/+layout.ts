import {authSettings, client} from "$lib/dp/client.svelte.js";
import type {CustomAuthModel, ExpandedClub, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
import type {LayoutLoad} from "./$types";
import {Collection} from "$lib/dp/enum/Collection.ts";

export const load = (async ({fetch, depends, parent}) => {
  const model = authSettings.record as CustomAuthModel;
  const data = await parent();

  if (!client.authStore.isValid) {
    return;
  }

  let clubs = data.clubs;
  let teams: ExpandedTeam[] | undefined;

  if (!teams) {
    teams = await client.collection(Collection.Teams).getFullList<ExpandedTeam>({
      filter: `"${model?.teams}" ?~ id`, // here we only display memberships
      expand: "club,admins",
      fetch: fetch,
      sort: "+name",
      requestKey: null,
    });
  }

  if (!clubs) {
    clubs = await client.collection(Collection.Clubs).getFullList<ExpandedClub>({
      filter: `"${model?.club}" ?~ id`,
      fetch: fetch,
      expand: "admins",
      requestKey: null,
    });
  }

  depends("teams:list");

  return {
    clubs: clubs,
    teams: teams,
  };
}) satisfies LayoutLoad;
