import {MatchAPIRequest} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";
import {PUBLIC_CLUB_ID} from "$env/static/public";
import {Gameday} from "bsm.js";
import type {LayoutServerLoad} from "../../../.svelte-kit/types/src/routes/gamecenter/$types";
import {selectedSeason} from "../../stores";
import {get} from "svelte/store";

export const load = (async () => {
    const matchRequest = new MatchAPIRequest(BSM_API_KEY)

    return {
        streamed: {
            matches: matchRequest.loadGamesForClub(Number(PUBLIC_CLUB_ID), get(selectedSeason), Gameday.any)
        }
    }
}) satisfies LayoutServerLoad