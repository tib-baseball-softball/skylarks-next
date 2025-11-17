import type {ClubFunction} from "bsm.js";
import {env} from "$env/dynamic/public";
import {client} from "$lib/dp/client.svelte.ts";
import type {LayoutLoad} from "../../../../../.svelte-kit/types/src/routes";

export const load: LayoutLoad = async ({fetch}) => {
  const clubOfficials = client.send<ClubFunction[]>("/api/bsm/relay", {
    fetch: fetch,
    query: {
      url: `https://bsm.baseball-softball.de/clubs/${env.PUBLIC_CLUB_ID}/club_functions.json`,
      club: env.PUBLIC_CLUB_ID,
    },
    requestKey: "officials",
  });

  return {
    clubOfficials: clubOfficials,
  };
};
