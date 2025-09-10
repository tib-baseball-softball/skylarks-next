/**
 * @see https://wuchale.dev/adapters/svelte/#sveltekit-ssrssg
 */
import {loadLocales, runWithLocale} from "wuchale/load-utils/server";
import type {Handle} from "@sveltejs/kit";
import {key, loadCatalog, loadIDs} from "virtual:wuchale/proxy";
import {locales} from "virtual:wuchale/locales";

// load at server startup
loadLocales(key, loadIDs, loadCatalog, locales);

export const handle: Handle = async ({event, resolve}) => {
  const locale = event.url.searchParams.get('locale') ?? 'en';
  return await runWithLocale(locale, () => resolve(event));
};