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

  //Construct PocketBase filter string from query parameters

  let filter = `team = "${team.id}"`

  let timeframe = url.searchParams.get("timeframe")

  if (!timeframe || timeframe === "next") {
    filter = filter.concat(`&& starttime >= @todayStart`)
  } else if (timeframe === "past") {
    filter = filter.concat(`&& starttime <= @todayStart`)
  }

  const showTypes = url.searchParams.get("type") ?? "any"

  if (showTypes !== "any") {
    filter = filter.concat(`&& type = "${showTypes}"`)
  }

  // add sort parameter

  let sort = "+starttime"

  if (url.searchParams.get("sort") === "desc") {
    sort = "-starttime"
  }

  // check pagination info

  const pageNumber = Number(url.searchParams.get("page")) ?? 1;

  const events = await watchWithPagination<ExpandedEvent>(
    "events",
    {
      filter: filter,
      sort: sort,
      expand: "participations_via_event.user, attire",
      fetch: fetch
    },
    pageNumber,
    6,
  );

  depends("event:list")

  return {
    team: team,
    events: events,
  };
};
