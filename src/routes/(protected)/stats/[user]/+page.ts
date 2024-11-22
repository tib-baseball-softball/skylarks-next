import type {PageLoad} from "./$types";
import {client} from "$lib/pocketbase/index.svelte";
import type {PersonalAttendanceStatsItem} from "$lib/model/PersonalAttendanceStats";
import type {UserStatsQuery} from "$lib/types/UserStatsQuery";

export const load = (async ({fetch, params, url}) => {
  const user = params.user

  const queryParams: UserStatsQuery = {}

  const season = url.searchParams.get("season");
  if (season) {
    queryParams.season = season
  }

  const statsItem = client.send<PersonalAttendanceStatsItem>(`/api/stats/${user}`, {
    fetch: fetch,
    query: queryParams
  })

  return {
    statsItem: statsItem
  }
}) satisfies PageLoad