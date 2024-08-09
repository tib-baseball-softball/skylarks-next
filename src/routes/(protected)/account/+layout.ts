import {client} from "$lib/pocketbase";
import type {LayoutLoad} from "../../../../.svelte-kit/types/src/routes/(protected)/account/$types";

export const load = (async ({ }) => {
    const teams = client.collection("teams").getFullList()

    return {
        teams: teams
    }
}) satisfies LayoutLoad