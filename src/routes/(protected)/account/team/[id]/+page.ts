import {client} from "$lib/pocketbase/index.svelte";
import {error} from "@sveltejs/kit";
import type {ExpandedTeam} from "$lib/model/ExpandedResponse.js";
import type {PageLoad} from "./$types";
import {EventService} from "$lib/service/EventService";
import type {EventseriesResponse} from "$lib/model/pb-types.ts";

export const load = (async ({fetch, parent, params, url, depends}) => {
  const data = await parent();
  const teams: ExpandedTeam[] = await data.teams;

  let team = teams.find((team) => team.id === params.id);

  if (!team) {
    team = await client
        .collection("teams")
        .getOne<ExpandedTeam>(params.id, {
          expand: "club",
          fetch: fetch
        });
  }
  if (!team) throw error(404, "Team not found");

  const eventService = new EventService()
  const events = await eventService.loadEventStore(team.id, url, fetch)
  const eventSeries = await client.collection("eventseries").getFullList<EventseriesResponse>({
    filter: `team = "${team.id}"`
  });

  depends("event:list")

  return {
    team: team,
    events: events,
    eventSeries: eventSeries,
  };
}) satisfies PageLoad
