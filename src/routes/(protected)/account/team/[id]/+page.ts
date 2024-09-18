import { client } from "$lib/pocketbase";
import { error } from "@sveltejs/kit";
import { watchWithPagination } from "$lib/pocketbase/RecordOperations";
import type { ExpandedEvent, ExpandedTeam } from "$lib/model/ExpandedResponse.js";

export const load = async ({ fetch, parent, params, depends }) => {
  const data = await parent();
  const teams: ExpandedTeam[] = await data.teams;

  let team = teams.find((team) => team.id === params.id);

  if (!team) {
    team = await client
      .collection("teams")
      .getOne<ExpandedTeam>(params.id, { expand: "club" });
  }
  if (!team) throw error(404, "Team not found");

  const events = await watchWithPagination<ExpandedEvent>(
    "events",
    {
      filter: `starttime >= @todayStart && team = "${team.id}"`,
      sort: "+starttime",
      expand: "participations_via_event.user, attire",
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
