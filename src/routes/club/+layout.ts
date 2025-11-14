import type {Club} from "bsm.js";
import {env} from "$env/dynamic/public";
import {client} from "$lib/dp/client.svelte.ts";
import type {LayoutLoad} from "./$types";

export const load: LayoutLoad = async ({fetch}) => {
  const bsmClubData = await client.send<Club>("/api/bsm/relay", {
    fetch: fetch,
    query: {
      url: `https://bsm.baseball-softball.de/clubs/${env.PUBLIC_CLUB_ID}.json`,
      club: env.PUBLIC_CLUB_ID,
    },
    requestKey: env.PUBLIC_CLUB_ID,
  });

  return {
    bsmClubData: bsmClubData,
  };
};
