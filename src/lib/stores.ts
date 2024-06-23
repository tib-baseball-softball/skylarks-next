import {persisted} from "svelte-persisted-store";
import type {AppPreferences} from "$lib/types/AppPreferences";
import {Gameday} from "bsm.js";

export const preferences: Persisted<AppPreferences> = persisted('preferences', {
    gameday: Gameday.current,
    selectedSeason: new Date().getFullYear(),
    leagueGroupID: 0
})