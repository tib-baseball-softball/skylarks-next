import {ClubTeamsAPIRequest} from "bsm.js/dist/service/ClubTeamsAPIRequest";
import {BSM_API_KEY} from "$env/static/private";
import {PUBLIC_CLUB_ID} from "$env/static/public";

export const load = (async ({ parent, url }) => {
    const data = await parent()

    let clubTeams = data.clubTeams

    const clubTeamRequest = new ClubTeamsAPIRequest(BSM_API_KEY)
    const season = url.searchParams.get("season")

    if (season) {
        clubTeams = clubTeamRequest.getTeamsForClub(Number(PUBLIC_CLUB_ID), Number(season))
    }

    return {
        clubTeams: clubTeams
    }
})