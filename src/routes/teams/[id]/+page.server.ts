import {error} from "@sveltejs/kit";
import {
  type BattingStatisticsEntry,
  type ClubTeam,
  ClubTeamsAPIRequest,
  type FieldingStatisticsEntry,
  LeagueGroupAPIRequest,
  type PitchingStatisticsEntry,
  StatsAPIRequest,
  StatsType,
  TablesAPIRequest
} from "bsm.js";
import {env} from "$env/dynamic/private";
import {env as publicEnv} from "$env/dynamic/public";

export async function load({parent, params, url}) {
  const data = await parent();
  let clubTeams = await data.clubTeams;

  // this is kind of hacky because I could not find an API call to get a single team, so if the collection is empty
  // we have to load them all again...could also be a skill issue with SvelteKit
  if (clubTeams?.length === 0) {
    const clubTeamRequest = new ClubTeamsAPIRequest(env.BSM_API_KEY);
    const season = url.searchParams.get("season")?.toString() ?? clubTeamRequest.defaultSeason;
    clubTeams = await clubTeamRequest.getTeamsForClub(Number(publicEnv.PUBLIC_CLUB_ID), Number(season));
  }

  const clubTeam: ClubTeam | undefined = clubTeams?.find((clubTeam) => clubTeam.id === Number(params.id));

  if (!clubTeam) {
    throw error(404);
  }

  const statsRequest = new StatsAPIRequest(env.BSM_API_KEY);
  const leagueEntry = clubTeam.team.league_entries.at(0);

  if (!leagueEntry) {
    return {
      clubTeam: clubTeam
    };
  }

  const battingStats = statsRequest.getStatisticsForLeagueEntry<BattingStatisticsEntry>(leagueEntry.id, StatsType.batting);
  const pitchingStats = statsRequest.getStatisticsForLeagueEntry<PitchingStatisticsEntry>(leagueEntry.id, StatsType.pitching);
  const fieldingStats = statsRequest.getStatisticsForLeagueEntry<FieldingStatisticsEntry>(leagueEntry.id, StatsType.fielding);

  let leagueGroups = await data.leagueGroups;
  let leagueGroup = leagueGroups?.find((group) => group.league.id === leagueEntry.league.id);

  // Caution: LeagueGroups can still be unavailable even though ClubTeams are
  if (!leagueGroup) {
    const leagueGroupRequest = new LeagueGroupAPIRequest(env.BSM_API_KEY);
    try {
      leagueGroups = await leagueGroupRequest.getLeagueGroupsForClub(Number(url.searchParams.get("season")) ?? leagueGroupRequest.defaultSeason);
      leagueGroup = leagueGroups?.find((group) => group.league.id === leagueEntry.league.id);
    } catch (e) {
      console.error("Could not load league group", e);
    }
  }

  const tableRequest = new TablesAPIRequest(env.BSM_API_KEY);

  return {
    clubTeam: clubTeam,
    battingStats: battingStats,
    pitchingStats: pitchingStats,
    fieldingStats: fieldingStats,
    table: leagueGroup ? tableRequest.getSingleTable(leagueGroup.id) : null,
  };
}