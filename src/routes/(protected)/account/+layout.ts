import {client} from "$lib/pocketbase";
import type {LayoutLoad} from "../../../../.svelte-kit/types/src/routes/(protected)/account/$types";
import type {TeamsResponse} from "$lib/model/pb-types";

export const load = (async ({ }) => {
    const teams = client.collection("teams").getFullList<TeamsResponse>({
        expand: "club"
    })

    return {
        teams: teams
    }
}) satisfies LayoutLoad