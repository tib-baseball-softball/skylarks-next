import type { PageStore } from "$lib/dp/records/PageStore.ts";
import { watchWithPagination } from "$lib/dp/records/RecordOperations.ts";
import type { ExpandedEvent } from "$lib/dp/types/ExpandedResponse.ts";
import { Collection } from "$lib/dp/enum/Collection.ts";

type BaseOptions = {
  url: URL;
  fetch: typeof window.fetch;
};
export type EventStoreOptions =
  | (BaseOptions & {
      mode: "club";
      clubID: string;
    })
  | (BaseOptions & {
      mode: "team";
      teamID: string;
    });

export class EventService {
  private DEFAULT_PER_PAGE = 6;
  /**
   * Takes a team ID and an URL, construct PocketBase filter string from query parameters
   * and loads events with realtime.
   */
  public async loadEventStore(
    options: EventStoreOptions,
  ): Promise<PageStore<ExpandedEvent>> {
    let filter = "";
    switch (options.mode) {
      case "team":
        filter = `(team = "${options.teamID}" || additional_teams.id ?= "${options.teamID}")`;
        break;
      case "club":
        filter = `(club = "${options.clubID}")`;
        break;
    }

    const timeframe = options.url.searchParams.get("timeframe");

    if (!timeframe || timeframe === "next") {
      filter = filter.concat(`&& starttime >= @todayStart`);
    } else if (timeframe === "past") {
      filter = filter.concat(`&& starttime <= @todayStart`);
    }

    const showTypes = options.url.searchParams.get("type") ?? "any";

    if (showTypes !== "any") {
      filter = filter.concat(`&& type = "${showTypes}"`);
    }

    // add sort parameter

    let sort = "+starttime";

    if (options.url.searchParams.get("sort") === "desc") {
      sort = "-starttime";
    }

    // check pagination info

    const pageNumber = Number(options.url.searchParams.get("page")) ?? 1;

    return await watchWithPagination<ExpandedEvent>(
      Collection.Events,
      {
        filter: filter,
        sort: sort,
        expand:
          "participations_via_event.user, attire, location, club, team, additional_teams",
        fetch: fetch,
        requestKey: `${options.mode === "club" ? options.clubID : options.teamID}-events`,
      },
      pageNumber,
      this.DEFAULT_PER_PAGE,
    );
  }
}
