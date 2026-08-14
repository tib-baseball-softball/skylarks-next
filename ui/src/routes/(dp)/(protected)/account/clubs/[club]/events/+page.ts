import { EventService } from "$lib/dp/service/EventService";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params, depends, url }) => {
  const eventService = new EventService();
  const events = await eventService.loadEventStore({
    url: url,
    fetch: fetch,
    mode: "club",
    clubID: params.club,
  });

  depends("event:list");

  return {
    eventStore: events,
  };
};
