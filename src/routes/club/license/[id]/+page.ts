import type {PageLoad} from './$types';
import {page} from "$app/state";
import type {License} from "bsm.js";
import {browser} from "$app/environment";
import {error} from "@sveltejs/kit";
import {client} from "$lib/pocketbase/index.svelte.ts";
import {env} from "$env/dynamic/public";

export const load: PageLoad = async ({params, fetch}) => {
  let license: License | undefined;

  if (browser) {
    const umpireLicenses: License[] = await page.data.umpireLicenses;
    const scorerLicenses: License[] = await page.data.scorerLicenses;

    license = scorerLicenses?.find(license => license.id === Number(params.id));

    if (!license) {
      license = umpireLicenses?.find(license => license.id === Number(params.id));
    }

    if (!license) {
      license = await client.send<License>("/api/bsm/relay", {
        fetch: fetch,
        query: {
          url: `https://bsm.baseball-softball.de/licenses/${params.id}.json`,
          club: env.PUBLIC_CLUB_ID
        },
        requestKey: `license-${params.id}`,
      });
    }

    if (!license) {
      error(404, "No license found.");
    }
  }

  return {
    license: license,
  };
};