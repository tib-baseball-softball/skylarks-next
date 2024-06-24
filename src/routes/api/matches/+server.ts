import {Gameday, MatchAPIRequest} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";
import {PUBLIC_CLUB_ID} from "$env/static/public";
import {json} from "@sveltejs/kit";

export async function GET({ url }): Promise<Response> {
    const matchRequest = new MatchAPIRequest(BSM_API_KEY)
    const currentYear = new Date().getFullYear()

    const season = Number(url.searchParams.get("season") ?? currentYear)
    const gameday = url.searchParams.get("gameday") ?? "current"

    if (url.searchParams.has("leagueID")) {
        return new Response()
    } else {
        const response = await matchRequest.loadGamesForClub(Number(PUBLIC_CLUB_ID), season, gameday)
        return json(response)
    }
}