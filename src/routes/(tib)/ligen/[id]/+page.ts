import {type LeagueGroup, LeagueGroupAPIRequest, TablesAPIRequest} from "bsm.js";
import {client} from "$lib/dp/client.svelte.ts";
import {RELAY_URL} from "$lib/tib/globals.svelte.ts";
import {error} from "@sveltejs/kit";
import type {PageLoad} from "./$types";

export const load: PageLoad = async ({ parent, params, fetch }) => {
  const data = await parent();
  const leagueGroups = await data.leagueGroups as LeagueGroup[];
  let leagueGroup: LeagueGroup | undefined = leagueGroups.find(
    (lg) => lg.id === Number(params.id)
  );

  const leagueGroupRequest = new LeagueGroupAPIRequest("");
  leagueGroupRequest.setAPIURL(RELAY_URL);

  if (!leagueGroup) {
    const lgURL = leagueGroupRequest.buildRequestURL(`league_groups/${params.id}.json`, []);
    leagueGroup = await client.send<LeagueGroup>(lgURL.pathname + lgURL.search, {
      fetch,
      requestKey: `ligen-id-leagueGroup-${params.id}`,
    });
  }

  if (!leagueGroup) {
    throw error(404, "leagueGroup not found");
  }

  // TODO: Verify exact tables endpoint path; assumed league_groups/{id}/table.json
  const tablesRequest = new TablesAPIRequest("");
  tablesRequest.setAPIURL(RELAY_URL);
  const tableURL = tablesRequest.buildRequestURL(`league_groups/${params.id}/table.json`, []);
  const table = client.send<any>(tableURL.pathname + tableURL.search, {
    fetch,
    requestKey: `ligen-id-table-${params.id}`,
  });

  return {
    table,
    leagueGroup,
  };
};
