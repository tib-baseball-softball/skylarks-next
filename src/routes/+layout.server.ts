import type {LayoutServerLoad} from "../../.svelte-kit/types/src/routes/$types";
import {ClubTeamsAPIRequest, LeagueGroupAPIRequest} from "bsm.js";
import {env} from "$env/dynamic/private";
import {get} from "svelte/store";
import {preferences} from "$lib/stores";
import {env as publicEnv} from "$env/dynamic/public";
import type {AppPreferences} from "$lib/types/AppPreferences";

export const load = (async () => {
  const leagueGroupRequest = new LeagueGroupAPIRequest(env.BSM_API_KEY)
  const clubTeamRequest = new ClubTeamsAPIRequest(env.BSM_API_KEY)
  const appPreferences = get(preferences) as AppPreferences

  return {
    clubTeams: clubTeamRequest.getTeamsForClub(Number(publicEnv.PUBLIC_CLUB_ID)),
    leagueGroups: leagueGroupRequest.getLeagueGroupsForClub(appPreferences.selectedSeason)
  }
}) satisfies LayoutServerLoad