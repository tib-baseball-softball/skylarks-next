import type { Match } from "bsm.js"
import {error} from "@sveltejs/kit";
import {BSM_API_KEY} from "$env/static/private";
import {MatchAPIRequest} from "bsm.js";

export async function load({ parent, params }) {
    const data = await parent()
    const matches = await data.streamed.matches
    const match = matches.find((match: Match) => match.id === Number(params.id))

    if (!match) throw error(404)

    const matchRequest = new MatchAPIRequest(BSM_API_KEY)

    return {
        match: match,
        singleGameStats: matchRequest.getBoxscoreForGame(match.id)
    }
}