import type {AppPreferences} from "$lib/types/AppPreferences";
import {Gameday} from "bsm.js";
import {persistedState} from "svelte-persisted-state";

/*
 * User Preferences, kept in sync via Local Storage.
 *
 * NB: Could be saved in the database for logged-in users
 *
 * Usage: https://github.com/oMaN-Rod/svelte-persisted-state
 */
export const preferences = persistedState<AppPreferences>('preferences', {
      gameday: Gameday.current,
      selectedSeason: new Date().getFullYear(),
      leagueGroupID: 0,
      favoriteTeamID: 0,
    },
    {
      storage: "local",
      syncTabs: true,
    }
);