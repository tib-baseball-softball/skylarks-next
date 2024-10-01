import { watchSingleRecord } from "$lib/pocketbase/RecordOperations";
import type { ExpandedEvent } from "$lib/model/ExpandedResponse.js";
import type { PageLoad } from "./$types";

export const load = (async ({ fetch, params, depends }) => {
  const event = await watchSingleRecord<ExpandedEvent>("events", params.id, {
    expand: "participations_via_event.user, attire, team.admins",
    fetch: fetch
  })

  depends("event:list")

  return {
    event: event
  }
}) satisfies PageLoad
