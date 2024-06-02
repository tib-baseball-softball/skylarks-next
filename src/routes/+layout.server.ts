import type {LayoutServerLoad} from "../../.svelte-kit/types/src/routes/$types";
import {LeagueGroupAPIRequest} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";
import {get} from "svelte/store";
import {selectedSeason} from "../stores";

export const load = (async () => {
    const leagueGroupRequest = new LeagueGroupAPIRequest(BSM_API_KEY)

    return {
        leagueGroups: leagueGroupRequest.getLeagueGroupsForClub(get(selectedSeason))
    }
}) satisfies LayoutServerLoad