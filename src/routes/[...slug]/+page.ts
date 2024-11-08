import type {PageLoad} from "../../../.svelte-kit/types/src/routes/[...slug]/$types";
import type {Typo3Page} from "$lib/model/Typo3Page";
import {env} from "$env/dynamic/public"

/**
 * This is the dynamically generated page load for all TYPO3 page for which no static
 * SvelteKit route exists. Page contents are requested from TYPO3 backend,
 * if nothing is found, display 404 page.
 */
export const load = (async ({fetch, params}) => {
  const slugParam = params.slug
  const url = `${env.PUBLIC_BACKEND_URL}/${slugParam}`

  const response = await fetch(url, {
    headers: {
      "Accept": "application/json",
    },
  })
  const responseData: Typo3Page = await response.json()

  return {
    pageObject: responseData,
    contentObjects: responseData.content.colPos0,
  }
}) satisfies PageLoad