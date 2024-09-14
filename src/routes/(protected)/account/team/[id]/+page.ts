import { client } from "$lib/pocketbase";
import { error } from "@sveltejs/kit";
import type { EventsResponse, TeamsResponse } from "$lib/model/pb-types";
import { watchWithPagination } from "$lib/pocketbase/RecordOperations";

export const load = async ({ fetch, parent, params, depends }) => {
  const data = await parent();
  const teams: TeamsResponse[] = await data.teams;

  let team = teams.find((team) => team.id === params.id);

  if (!team) {
    team = await client
      .collection("teams")
      .getOne<TeamsResponse>(params.id, { expand: "club" });
  }
  if (!team) throw error(404, "Team not found");

  const events = await watchWithPagination<EventsResponse>(
    "events",
    {
      filter: `starttime >= @todayStart && team = "${team.id}"`,
      sort: "+starttime",
      expand: "participations_via_event.user",
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
