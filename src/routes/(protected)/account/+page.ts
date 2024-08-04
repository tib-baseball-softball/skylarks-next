import type { PageLoad } from "./$types"
import {client} from "$lib/pocketbase";
import {authModel} from "$lib/pocketbase";
import {get} from "svelte/store";
import type {ClubsResponse} from "$lib/model/pb-types";

export const load = (async ({ }) => {
    const modelValue = get(authModel)
    
    console.log(modelValue)
    const ids: string[] = modelValue?.expand?.club.map((club: ClubsResponse) => club.id)
    
    const clubs = client.collection("clubs").getFullList(
       // {filter: `id ?= ${ids}`}
    )

    return {
        clubs: clubs
    }
}) satisfies PageLoad