import {persisted} from "svelte-persisted-store";
import type {AppPreferences} from "$lib/types/AppPreferences";

export const preferences: Persisted<AppPreferences> = persisted('preferences', {
    selectedSeason: new Date().getFullYear()
})