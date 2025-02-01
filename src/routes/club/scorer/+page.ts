import type {PageLoad} from './$types';
import {client} from "$lib/pocketbase/index.svelte.ts";
import type {License} from "bsm.js";
import {env} from "$env/dynamic/public";

export const load: PageLoad = async ({fetch}) => {
  const scorerLicenses = client.send<License[]>("api/bsm/relay", {
    fetch: fetch,
    query: {
      url: `https://bsm.baseball-softball.de/clubs/${env.PUBLIC_CLUB_ID}/licenses.json?filters[categories][]=scorer`,
      club: env.PUBLIC_CLUB_ID,
    }
  });

  return {
    scorerLicenses: scorerLicenses,
  };
};