import type {PageLoad} from "./$types";
import {authSettings, client} from "$lib/pocketbase/index.svelte";
import type {PersonalAttendanceStatsItem} from "$lib/model/PersonalAttendanceStats";
import type {UserStatsQuery} from "$lib/types/UserStatsQuery";
import type {CustomAuthModel} from "$lib/model/ExpandedResponse";

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

  let userRecord: CustomAuthModel

  // in case the currently logged-in user wants to see their own stats page, skip loading record
  if (authSettings?.record?.id === user) {
    userRecord = authSettings.record as CustomAuthModel
  } else {
    userRecord = await client.collection("users").getOne<CustomAuthModel>(user)
  }

  let teamStatsItems: Promise<PersonalAttendanceStatsItem>[] = []
  for (const team of userRecord.teams) {
    queryParams.team = team
    teamStatsItems.push(client.send<PersonalAttendanceStatsItem>(`/api/stats/${user}`, {
      fetch: fetch,
      query: queryParams,
      requestKey: team,
    }))
  }

  return {
    user: userRecord,
    statsItem: statsItem,
    teamStatsItems: teamStatsItems
  }
}) satisfies PageLoad