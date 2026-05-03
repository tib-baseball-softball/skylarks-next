import type {LayoutLoad} from "./$types";
import {authSettings, client} from "$lib/dp/client.svelte.ts";
import {browser} from "$app/environment";
import {loadLocale} from "wuchale/load-utils";
import type {CustomAuthModel, ExpandedClub, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
import {locales} from "../locales/data";
import {appLocale} from "$lib/dp/locale.svelte.ts";
import {Collection} from "$lib/dp/enum/Collection.ts";

export const ssr = false;

export const load: LayoutLoad = async ({fetch, depends}) => {
  let clubs: ExpandedClub[] = [];
  let teams: ExpandedTeam[] = [];

  if (browser) {
    let locale: string;
    if (appLocale.current) {
      // locale has been set before
      locale = appLocale.current;
    } else {
      // locale is unset so far, try to read from browser settings
      const browserPreferredLocale = navigator.language.slice(0, 2); // Safari has `de-DE`, Chrome and Firefox use `de`

      // @ts-ignore - locales is wrongly inferred to be Intl.Locale instead of simple Wuchale union type
      if (locales.includes(browserPreferredLocale)) {
        locale = browserPreferredLocale;
      } else {
        locale = "en";
      }
    }
    await loadLocale(locale);
  }

  /**
   * setting requestKey to `null` means "do not autocancel" => nav is always up to date
   */
  if (browser && client.authStore.isValid) {
    const authRecord = authSettings.record as CustomAuthModel;

    // both calls are intentionally not filtered to let auth rules do the filtering
    teams = await client.collection(Collection.Teams).getFullList<ExpandedTeam>({
      expand: "club,admins",
      fetch: fetch,
      sort: "+name",
      requestKey: null,
    });

    clubs = await client.collection(Collection.Clubs).getFullList<ExpandedClub>({
      filter: `"${authRecord?.club}" ?~ id`,
      fetch: fetch,
      expand: "admins",
      requestKey: null,
    });
  }
  depends("nav:load");

  return {
    clubs: clubs,
    teams: teams,
  };
};
