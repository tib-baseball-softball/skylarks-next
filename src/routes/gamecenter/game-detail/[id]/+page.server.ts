import type { Match } from "bsm.js"
import {error} from "@sveltejs/kit";
import {BSM_API_KEY} from "$env/static/private";
import {MatchAPIRequest} from "bsm.js";

export async function load({ parent, params }) {
    const data = await parent()
    const matches = await data.streamed.matches
    let match = matches.find((match: Match) => match.id === Number(params.id))

    const matchRequest = new MatchAPIRequest(BSM_API_KEY)

    // if not already found in data, load it yourself
    if (!match) {
        try {
            match = await matchRequest.loadSingleMatch(Number(params.id))
        }
        catch (e) {
            if (e instanceof Error) {
                console.error(e.message)
            }
        }
    }

    if (!match) throw error(404, "Spiel konnte nicht gefunden werden.")

    return {
        match: match,
        singleGameStats: matchRequest.getBoxscoreForGame(match.id)
    }
}