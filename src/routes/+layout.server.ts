import type {LayoutServerLoad} from "../../.svelte-kit/types/src/routes/$types";
import {ClubTeamsAPIRequest, LeagueGroupAPIRequest} from "bsm.js";
import {env} from "$env/dynamic/private";
import {preferences} from "$lib/globals.svelte.ts";
import {env as publicEnv} from "$env/dynamic/public";

export const load = (async () => {
  const leagueGroupRequest = new LeagueGroupAPIRequest(env.BSM_API_KEY)
  const clubTeamRequest = new ClubTeamsAPIRequest(env.BSM_API_KEY)
  const appPreferences = preferences.current

  return {
    clubTeams: clubTeamRequest.getTeamsForClub(Number(publicEnv.PUBLIC_CLUB_ID)),
    leagueGroups: leagueGroupRequest.getLeagueGroupsForClub(appPreferences.selectedSeason)
  }
}) satisfies LayoutServerLoad