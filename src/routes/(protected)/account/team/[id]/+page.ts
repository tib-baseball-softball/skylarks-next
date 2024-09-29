import { client } from "$lib/pocketbase";
import { error } from "@sveltejs/kit";
import { watchWithPagination } from "$lib/pocketbase/RecordOperations";
import type { ExpandedEvent, ExpandedTeam } from "$lib/model/ExpandedResponse.js";

export const load = async ({ fetch, parent, params, url, depends }) => {
  const data = await parent();
  const teams: ExpandedTeam[] = await data.teams;

  let team = teams.find((team) => team.id === params.id);

  if (!team) {
    team = await client
      .collection("teams")
      .getOne<ExpandedTeam>(params.id, { expand: "club" });
  }
  if (!team) throw error(404, "Team not found");

  let filter = url.searchParams.get("filter")

  if (!filter || filter === "next") {
    filter = `starttime >= @todayStart && team = "${team.id}"`
  } else if (filter === "past") {
    filter = `starttime <= @todayStart && team = "${team.id}"`
  }

  let sort = "+starttime"
  
  if (url.searchParams.get("sort") === "desc") {
    sort = "-starttime"
  }

  const events = await watchWithPagination<ExpandedEvent>(
    "events",
    {
      filter: filter,
      sort: sort,
      expand: "participations_via_event.user, attire",
      fetch: fetch
    },
    1,
    6,
  );

  depends("event:list")

  return {
    team: team,
    events: events,
  };
};
