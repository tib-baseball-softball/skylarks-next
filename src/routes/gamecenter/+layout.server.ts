import {MatchAPIRequest} from "../../../../bsm.js/src/service/MatchAPIRequest";
import {BSM_API_KEY} from "$env/static/private";
import {PUBLIC_CLUB_ID} from "$env/static/public";
import {Gameday} from "../../../../bsm.js/src/enum/Gameday";
import type {LayoutServerLoad} from "../../../.svelte-kit/types/src/routes/gamecenter/$types";

export const load = (async () => {
    const matchRequest = new MatchAPIRequest(BSM_API_KEY)

    return {
        streamed: {
            matches: matchRequest.loadGamesForClub(Number(PUBLIC_CLUB_ID), 2024, Gameday.any)
        }
    }
}) satisfies LayoutServerLoad