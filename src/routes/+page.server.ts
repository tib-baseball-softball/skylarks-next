import {MatchAPIRequest} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";
import {PUBLIC_CLUB_ID} from "$env/static/public";
import {selectedSeason} from "../stores";
import {get} from "svelte/store";
import {Gameday} from "bsm.js";
import type {PageServerLoad} from "../../.svelte-kit/types/src/routes/$types";

export const load = (async ({ fetch }) => {
    const matchRequest = new MatchAPIRequest(BSM_API_KEY)

    const matchesCurrent = matchRequest.loadGamesForClub(Number(PUBLIC_CLUB_ID), get(selectedSeason), Gameday.current)
    const matchesPrevious = matchRequest.loadGamesForClub(Number(PUBLIC_CLUB_ID), get(selectedSeason), Gameday.previous)
    const matchesNext = matchRequest.loadGamesForClub(Number(PUBLIC_CLUB_ID), get(selectedSeason), Gameday.next)

    return {
        matchesCurrent: matchesCurrent,
        matchesPrevious: matchesPrevious,
        matchesNext: matchesNext,
    }
}) satisfies PageServerLoad