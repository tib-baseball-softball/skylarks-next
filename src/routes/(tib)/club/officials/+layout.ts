import type {ClubFunction} from "bsm.js";
import {env} from "$env/dynamic/public";
import {client} from "$lib/dp/client.svelte.ts";
import type {LayoutLoad} from "./$types";

export const load: LayoutLoad = async ({fetch}) => {
  const clubOfficials = client.send<ClubFunction[]>(`/api/bsm/relay/clubs/${env.PUBLIC_CLUB_ID}/club_functions.json`, {
    fetch: fetch,
    requestKey: "officials",
  });

  return {
    clubOfficials: clubOfficials,
  };
};
