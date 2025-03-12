import type { ExpandedEvent } from "$lib/model/ExpandedResponse";
import type { PageStore } from "$lib/pocketbase/PageStore";
import { watchWithPagination } from "$lib/pocketbase/RecordOperations";
import type { Fetch } from "$lib/types/Fetch.ts";

export class EventService {
  /**
   * Takes a team ID and an URL, construct PocketBase filter string from query parameters
   * and loads events with realtime.
   */
  public async loadEventStore(teamID: string, url: URL, fetch: Fetch): Promise<PageStore<ExpandedEvent>> {
    let filter = `team = "${teamID}"`;

    let timeframe = url.searchParams.get("timeframe");

    if (!timeframe || timeframe === "next") {
      filter = filter.concat(`&& starttime >= @todayStart`);
    } else if (timeframe === "past") {
      filter = filter.concat(`&& starttime <= @todayStart`);
    }

    const showTypes = url.searchParams.get("type") ?? "any";

    if (showTypes !== "any") {
      filter = filter.concat(`&& type = "${showTypes}"`);
    }

    // add sort parameter

    let sort = "+starttime";

    if (url.searchParams.get("sort") === "desc") {
      sort = "-starttime";
    }

    // check pagination info

    const pageNumber = Number(url.searchParams.get("page")) ?? 1;

    return await watchWithPagination<ExpandedEvent>(
      "events",
      {
        filter: filter,
        sort: sort,
        expand: "participations_via_event.user, attire, location",
        fetch: fetch
      },
      pageNumber,
      6,
    );
  }
}
