import type { EventsResponse } from "$lib/model/pb-types";
import { watchSingleRecord } from "$lib/pocketbase/RecordOperations";

export const load = (async ({ fetch, params, depends }) => {
  const event = await watchSingleRecord<EventsResponse>("events", params.id, {
    expand: "participations_via_event.user",
    fetch: fetch
  })

  depends("event:list")

  return {
    event: event
  }
})
