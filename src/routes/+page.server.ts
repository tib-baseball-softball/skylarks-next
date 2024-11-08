import {Gameday, MatchAPIRequest} from "bsm.js";
import {env} from "$env/dynamic/private"
import {env as publicEnv} from "$env/dynamic/public";
import {preferences} from "$lib/stores";
import {get} from "svelte/store";
import type {PageServerLoad} from "../../.svelte-kit/types/src/routes/$types";
import type {AppPreferences} from "$lib/types/AppPreferences";

export const load = (async ({fetch}) => {
  const matchRequest = new MatchAPIRequest(env.BSM_API_KEY)
  const appPreferences = get(preferences) as AppPreferences

  const matchesCurrent = matchRequest.loadGamesForClub(Number(publicEnv.PUBLIC_CLUB_ID), appPreferences.selectedSeason, Gameday.current)
  const matchesPrevious = matchRequest.loadGamesForClub(Number(publicEnv.PUBLIC_CLUB_ID), appPreferences.selectedSeason, Gameday.previous)
  const matchesNext = matchRequest.loadGamesForClub(Number(publicEnv.PUBLIC_CLUB_ID), appPreferences.selectedSeason, Gameday.next)

  return {
    matchesCurrent: matchesCurrent,
    matchesPrevious: matchesPrevious,
    matchesNext: matchesNext,
  }
}) satisfies PageServerLoad