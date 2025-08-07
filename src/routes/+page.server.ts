import {Gameday, MatchAPIRequest} from "bsm.js";
import {env} from "$env/dynamic/private";
import {env as publicEnv} from "$env/dynamic/public";
import {preferences} from "$lib/globals.svelte.ts";
import type {PageServerLoad} from "../../.svelte-kit/types/src/routes/$types";

export const load = (async ({fetch}) => {
  const matchRequest = new MatchAPIRequest(env.BSM_API_KEY)
  const appPreferences = preferences.current

  const matchesCurrent = matchRequest.loadGamesForClub(Number(publicEnv.PUBLIC_CLUB_ID), appPreferences.selectedSeason, Gameday.current, true)
  const matchesPrevious = matchRequest.loadGamesForClub(Number(publicEnv.PUBLIC_CLUB_ID), appPreferences.selectedSeason, Gameday.previous, true)
  const matchesNext = matchRequest.loadGamesForClub(Number(publicEnv.PUBLIC_CLUB_ID), appPreferences.selectedSeason, Gameday.next, true)

  return {
    matchesCurrent: matchesCurrent,
    matchesPrevious: matchesPrevious,
    matchesNext: matchesNext,
  }
}) satisfies PageServerLoad