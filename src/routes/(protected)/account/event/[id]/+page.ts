import type { EventsResponse } from "$lib/model/pb-types";
import { client } from "$lib/pocketbase";

export const load = (async ({ fetch, params, depends }) => {
  const event = client.collection("events").getOne<EventsResponse>(params.id, {
    expand: "participations_via_event.user",
    fetch: fetch
  })

  depends("event:list")

  return {
    event: event
  }
})
