import {LeagueGroupAPIRequest} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";

export async function load({ parent, url }) {
    const data = await parent()

    let leagueGroups = data.leagueGroups

    const leagueGroupRequest = new LeagueGroupAPIRequest(BSM_API_KEY)
    const season = url.searchParams.get("season")?.toString()

    if (season) {
        leagueGroups = leagueGroupRequest.getLeagueGroupsForClub(Number(season))
    }

    return {
        leagueGroups: leagueGroups
    }
}