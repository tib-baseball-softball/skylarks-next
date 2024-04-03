import {PUBLIC_BACKEND_URL} from "$env/static/public";
import type {Typo3Page} from "$lib/model/Typo3Page";
import type {PageLoad} from "../../.svelte-kit/types/src/routes/[...slug]/$types";

export const load = (async ({ fetch }) => {
    const url = `${PUBLIC_BACKEND_URL}`

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