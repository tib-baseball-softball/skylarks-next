import {LeagueGroupAPIRequest} from "bsm.js";
import {env} from "$env/dynamic/private";

export async function load({parent, url}) {
  const data = await parent()

  let leagueGroups = data.leagueGroups

  const leagueGroupRequest = new LeagueGroupAPIRequest(env.BSM_API_KEY)
  const season = url.searchParams.get("season")?.toString()

  if (season) {
    leagueGroups = leagueGroupRequest.getLeagueGroupsForClub(Number(season))
  }

  return {
    leagueGroups: leagueGroups
  }
}