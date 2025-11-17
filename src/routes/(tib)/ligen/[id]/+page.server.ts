import {type LeagueGroup, LeagueGroupAPIRequest, TablesAPIRequest} from "bsm.js";
import {env} from "$env/dynamic/private";
import {error} from "@sveltejs/kit";
import type {PageServerLoad} from "../../../../../.svelte-kit/types/src/routes";

export const load: PageServerLoad = async ({parent, params}) => {
  const data = await parent();
  const leagueGroups = await data.leagueGroups;
  let leagueGroup: LeagueGroup | undefined = leagueGroups.find(
      (leagueGroup) => leagueGroup.id === Number(params.id)
  );

  if (!leagueGroup) {
    try {
      const leagueGroupRequest = new LeagueGroupAPIRequest(env.BSM_API_KEY);
      leagueGroup = await leagueGroupRequest.getSingleLeagueGroup(Number(params.id));
    } catch (error) {
      console.error(error);
    }
  }

  if (!leagueGroup) {
    throw error(404, "leagueGroup not found");
  }

  const tableRequest = new TablesAPIRequest(env.BSM_API_KEY);
  const table = tableRequest.getSingleTable(Number(params.id));

  return {
    table: table,
    leagueGroup: leagueGroup,
  };
};
