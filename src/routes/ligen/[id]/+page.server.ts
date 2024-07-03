import {type LeagueGroup, LeagueGroupAPIRequest, TablesAPIRequest} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";
import {error} from "@sveltejs/kit";

export async function load({ parent, params }) {
    const data = await parent()
    const leagueGroups = await data.leagueGroups
    let leagueGroup: LeagueGroup | undefined = leagueGroups.find((leagueGroup) => leagueGroup.id === Number(params.id))

    if (!leagueGroup) {
        try {
            const leagueGroupRequest = new LeagueGroupAPIRequest(BSM_API_KEY)
            leagueGroup = await leagueGroupRequest.getSingleLeagueGroup(Number(params.id))
        } catch (error) {
            console.error(error)
        }
    }

    if (!leagueGroup) {
        throw error(404)
    }

    const tableRequest = new TablesAPIRequest(BSM_API_KEY)
    const table = tableRequest.getSingleTable(Number(params.id))

    return {
        table: table,
        leagueGroup: leagueGroup
    }
}