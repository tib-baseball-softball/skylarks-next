import {LeagueGroupAPIRequest, MatchAPIRequest} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";
import {Gameday} from "bsm.js";
import type {LayoutServerLoad} from "../../../.svelte-kit/types/src/routes/gamecenter/$types";
import {error} from "@sveltejs/kit";

export const load = (async ({ url }) => {
    const matchRequest = new MatchAPIRequest(BSM_API_KEY)
    const season = url.searchParams.get("season")?.toString() ?? matchRequest.defaultSeason
    const gameday = url.searchParams.get("gameday") as Gameday ?? Gameday.current
    const leagueGroup = url.searchParams.get("leagueGroup") ?? "0"
    const search = url.searchParams.get("search") as string | undefined

    // may need more validation
    if (!(gameday in Gameday)) {
        throw error(400, "Invalid query parameters given.")
    }

    const leagueGroupRequest = new LeagueGroupAPIRequest(BSM_API_KEY)

    return {
        leagueGroups: leagueGroupRequest.getLeagueGroupsForClub(Number(season)),
        streamed: {
            matches: matchRequest.loadAllGames(Number(season), gameday, Number(leagueGroup), search)
        }
    }
}) satisfies LayoutServerLoad