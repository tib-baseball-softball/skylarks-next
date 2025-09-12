import type {LayoutLoad} from "../../.svelte-kit/types/src/routes/$types";
import type {CustomAuthModel, ExpandedClub, ExpandedTeam,} from "$lib/model/ExpandedResponse";
import {browser} from "$app/environment";
import {authSettings, client} from "$lib/pocketbase/index.svelte";
import '../locales/loader.svelte.js';
// @ts-ignore
import {locales} from "virtual:wuchale/locales";
import {loadLocale} from "wuchale/load-utils";
import {preferences} from "$lib/globals.svelte.ts";

/**
 * This loads club and team data for the extended DiamondPlanner nav bar.
 *
 * Should be executed only client-side and only if logged in.
 */
export const load = (async ({data, fetch, depends}) => {
  let clubs: ExpandedClub[] = [];
  let teams: ExpandedTeam[] = [];

  if (browser) {
    let locale: string;
    if (preferences.current.locale) {
      // locale has been set before
      locale = preferences.current.locale;
    } else {
      // locale is unset so far, try to read from browser settings
      const browserPreferredLocale = navigator.language.slice(0, 2); // Safari has `de-DE`, Chrome and Firefox use `de`

      if (locales.includes(browserPreferredLocale)) {
        locale = browserPreferredLocale;
      } else {
        locale = "en";
      }
    }
    await loadLocale(locale);
  }

  if (browser && client.authStore.isValid) {
    /**
     * setting requestKey to `null` means "do not autocancel" => nav is always up to date
     */
    teams = await client.collection("teams").getFullList<ExpandedTeam>({
      expand: "club,admins",
      fetch: fetch,
      requestKey: null,
      sort: "+name",
    });

    const authRecord = authSettings.record as CustomAuthModel;
    clubs = await client.collection("clubs").getFullList<ExpandedClub>({
      filter: `"${authRecord?.club}" ?~ id`,
      fetch: fetch,
      requestKey: null,
      expand: "admins",
    });
  }
  depends("nav:load");

  return {
    clubTeams: data.clubTeams,
    leagueGroups: data.leagueGroups,
    clubs: clubs,
    teams: teams,
  };
}) satisfies LayoutLoad;
