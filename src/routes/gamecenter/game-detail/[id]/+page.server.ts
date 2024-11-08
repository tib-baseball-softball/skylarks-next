import type {Match} from "bsm.js"
import {MatchAPIRequest} from "bsm.js";
import {error} from "@sveltejs/kit";
import {env} from "$env/dynamic/private";

export async function load({parent, params}) {
  const data = await parent()
  const matches = await data.streamed.matches
  let match = matches.find((match: Match) => match.id === Number(params.id))

  const matchRequest = new MatchAPIRequest(env.BSM_API_KEY)

  // if not already found in data, load it yourself
  if (!match) {
    try {
      match = await matchRequest.loadSingleMatch(Number(params.id))
    } catch (e) {
      if (e instanceof Error) {
        console.error(e.message)
      }
    }
  }

  if (!match) throw error(404, "Match couldn't be found.")

  return {
    match: match,
    singleGameStats: matchRequest.getBoxscoreForGame(match.id)
  }
}