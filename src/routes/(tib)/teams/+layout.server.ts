import {ClubTeamsAPIRequest} from "bsm.js";
import {env} from "$env/dynamic/private";
import {env as publicEnv} from "$env/dynamic/public";
import type {LayoutServerLoad} from "../../../../.svelte-kit/types/src/routes";

export const load: LayoutServerLoad = async ({parent, url}) => {
  const data = await parent();

  let clubTeams = data.clubTeams;

  const clubTeamRequest = new ClubTeamsAPIRequest(env.BSM_API_KEY);
  const season = url.searchParams.get("season");

  if (season) {
    clubTeams = clubTeamRequest.getTeamsForClub(Number(publicEnv.PUBLIC_CLUB_ID), Number(season));
  }

  return {
    clubTeams: clubTeams,
  };
};
