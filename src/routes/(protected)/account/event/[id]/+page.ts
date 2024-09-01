import type { EventsResponse } from "$lib/model/pb-types";
import { client } from "$lib/pocketbase";

export const load = (async ({ fetch, params }) => {
  const event = client.collection("events").getOne<EventsResponse>(params.id, {
    expand: "participations_via_event.user",
    fetch: fetch
  })

  return {
    event: event
  }
})
