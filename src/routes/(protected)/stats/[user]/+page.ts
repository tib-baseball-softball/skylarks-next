import type {PageLoad} from "./$types";
import {client} from "$lib/pocketbase/index.svelte";
import type {PersonalAttendanceStatsItem} from "$lib/model/PersonalAttendanceStats";

export const load = (async ({fetch, params}) => {
  const user = params.user

  const statsItem = client.send<PersonalAttendanceStatsItem>(`/api/stats/${user}`, {
    fetch: fetch
  })

  return {
    statsItem: statsItem
  }
}) satisfies PageLoad