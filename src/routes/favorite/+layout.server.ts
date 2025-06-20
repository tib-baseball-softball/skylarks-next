import {ClubTeamsAPIRequest} from "bsm.js";
import {env} from "$env/dynamic/private";
import {env as publicEnv} from "$env/dynamic/public";

export const load = (async ({url}) => {
  const clubTeamRequest = new ClubTeamsAPIRequest(env.BSM_API_KEY);
  const season = url.searchParams.get("season") ?? clubTeamRequest.defaultSeason;

  return {
    clubTeams: clubTeamRequest.getTeamsForClub(Number(publicEnv.PUBLIC_CLUB_ID), Number(season)),
  };
});