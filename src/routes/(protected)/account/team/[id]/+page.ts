import {client, watch} from "$lib/pocketbase";
import {error} from "@sveltejs/kit";
import type {EventsResponse, TeamsResponse} from "$lib/model/pb-types";

export const load = (async ({ parent, params }) => {
    const data = await parent()
    const teams = await data.teams
    
    let team = teams.find((team) => team.id === params.id)
    
    if (!team) {
        team = await client.collection("teams").getOne<TeamsResponse>(params.id, {expand: "club"})
    }
    if (!team) throw error(404, "Team nicht gefunden")
    
    const events = await watch<EventsResponse>("events", {
        filter: `starttime >= @todayStart && team = "${team.id}"`,
        sort: '+starttime',
    })
    
    return {
        team: team,
        events: events,
    }
})