import type {LayoutServerLoad} from "../../.svelte-kit/types/src/routes/$types";
import {ClubTeamsAPIRequest, LeagueGroupAPIRequest} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";
import {get} from "svelte/store";
import {preferences} from "$lib/stores";
import {PUBLIC_CLUB_ID} from "$env/static/public";
import type {AppPreferences} from "$lib/types/AppPreferences";

export const load = (async () => {
    const leagueGroupRequest = new LeagueGroupAPIRequest(BSM_API_KEY)
    const clubTeamRequest = new ClubTeamsAPIRequest(BSM_API_KEY)
    const appPreferences = get(preferences) as AppPreferences

    return {
        clubTeams: clubTeamRequest.getTeamsForClub(Number(PUBLIC_CLUB_ID)),
        leagueGroups: leagueGroupRequest.getLeagueGroupsForClub(appPreferences.selectedSeason)
    }
}) satisfies LayoutServerLoad