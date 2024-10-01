import type { CustomAuthModel, ExpandedTeam } from '$lib/model/ExpandedResponse';
import { client } from '$lib/pocketbase';
import type { PageLoad } from './$types';
import { watchWithPagination } from '$lib/pocketbase/RecordOperations';
import { error } from '@sveltejs/kit';

export const load = (async ({ fetch, parent, params, depends }) => {
  const data = await parent();
  const teams: ExpandedTeam[] = await data.teams;

  let team = teams.find((team) => team.id === params.id);

  if (!team) {
    team = await client
      .collection("teams")
      .getOne<ExpandedTeam>(params.id, { expand: "club" });
  }
  if (!team) throw error(404, "Team not found");

  const players = await watchWithPagination<CustomAuthModel>(
    "users",
    {
      filter: `teams ?~ "${params.id}"`,
      fetch: fetch,
    },
    1,
    99,
  )

  depends("members:list")

  return {
    players: players,
    team: team,
  };
}) satisfies PageLoad;
