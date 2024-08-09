import {client} from "$lib/pocketbase";

export const load = (async ({params}) => {
    const event = client.collection("events").getOne(params.id, { expand: "participants" })
    
    return {
        event: event
    }
})