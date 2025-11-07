import { ClubTeamsAPIRequest, LeagueGroupAPIRequest } from "bsm.js"
import { env } from "$env/dynamic/private"
import { preferences } from "$lib/globals.svelte.ts"
import { env as publicEnv } from "$env/dynamic/public"
import type { LayoutServerLoad } from "./$types"

export const load: LayoutServerLoad = async () => {
  const leagueGroupRequest = new LeagueGroupAPIRequest(env.BSM_API_KEY)
  const clubTeamRequest = new ClubTeamsAPIRequest(env.BSM_API_KEY)
  const appPreferences = preferences.current

  const clubTeams = clubTeamRequest.getTeamsForClub(
    Number(publicEnv.PUBLIC_CLUB_ID),
    appPreferences.selectedSeason
  )

  const leagueGroups = leagueGroupRequest.getLeagueGroupsForClub(appPreferences.selectedSeason)

  return {
    clubTeams: clubTeams,
    leagueGroups: leagueGroups,
  }
}
