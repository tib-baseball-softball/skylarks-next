import {LeagueGroupAPIRequest, type Match, MatchAPIRequest} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";
import {Gameday} from "bsm.js";
import type {LayoutServerLoad} from "../../../.svelte-kit/types/src/routes/gamecenter/$types";
import {error} from "@sveltejs/kit";
import {PUBLIC_CLUB_ID} from "$env/static/public";

export const load = (async ({ url }) => {
    const matchRequest = new MatchAPIRequest(BSM_API_KEY)
    const season = url.searchParams.get("season")?.toString() ?? matchRequest.defaultSeason
    const gameday = url.searchParams.get("gameday") as Gameday ?? Gameday.current
    const leagueGroup = url.searchParams.get("leagueGroup")
    const search = url.searchParams.get("search") as string | undefined

    // may need more validation
    if (!(gameday in Gameday)) {
        throw error(400, "Invalid query parameters given.")
    }

    const leagueGroupRequest = new LeagueGroupAPIRequest(BSM_API_KEY)


    // Our Club API key has access to basically all games in Germany which takes a long time to load and is not relevant
    // most of the time. Restrict games to Club Games except when a specific league is requested.
    let matchesPromise: Promise<Match[]>

    if (leagueGroup && leagueGroup !== "0") {
        matchesPromise = matchRequest.loadAllGames(Number(season), gameday, Number(leagueGroup), search)
    } else {
        matchesPromise = matchRequest.loadGamesForClub(Number(PUBLIC_CLUB_ID), Number(season), gameday)
    }

    return {
        leagueGroups: leagueGroupRequest.getLeagueGroupsForClub(Number(season)),
        streamed: {
            matches: matchesPromise
        }
    }
}) satisfies LayoutServerLoad