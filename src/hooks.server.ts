import {loadLocales, runWithLocale} from "wuchale/load-utils/server";
import type {Handle} from "@sveltejs/kit";
import {key, loadCatalog, loadIDs} from './locales/loader.ssr.svelte.js';
import {locales} from "virtual:wuchale/locales";

/**
 * @see https://wuchale.dev/adapters/svelte/#sveltekit-ssrssg
 */

// load at server startup
loadLocales(key, loadIDs, loadCatalog, locales);

// TODO: server will always load English for now
export const handle: Handle = async ({event, resolve}) => {
  const locale = event.url.searchParams.get('locale') ?? 'en';
  return await runWithLocale(locale, () => resolve(event));
};