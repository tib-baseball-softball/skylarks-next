import { client } from "$lib/pocketbase";
import { error } from "@sveltejs/kit";
import type { ExpandedTeam } from "$lib/model/ExpandedResponse.js";
import type { PageLoad } from "./$types";
import { EventService } from "$lib/service/EventService";

export const load = (async ({ fetch, parent, params, url, depends }) => {
  const data = await parent();
  const teams: ExpandedTeam[] = await data.teams;

  let team = teams.find((team) => team.id === params.id);

  if (!team) {
    team = await client
      .collection("teams")
      .getOne<ExpandedTeam>(params.id, { expand: "club" });
  }
  if (!team) throw error(404, "Team not found");

  const eventService = new EventService()
  const events = await eventService.loadEventStore(team.id, url, fetch)

  depends("event:list")

  return {
    team: team,
    events: events,
  };
}) satisfies PageLoad
