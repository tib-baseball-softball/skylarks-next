import type {PageServerLoad} from "./$types";
import {
  type BattingStatisticsEntry,
  type FieldingStatisticsEntry,
  type PitchingStatisticsEntry,
  StatsAPIRequest,
  StatsType
} from "bsm.js";
import {env} from "$env/dynamic/private";
import {error} from "@sveltejs/kit";

export const load: PageServerLoad = async ({params, url}) => {
  const leagueEntry = params.entry;
  const teamName = url.searchParams.get("team")

  if (!leagueEntry) {
    throw error(400, {message: "No league entry provided."});
  }

  const statsRequest = new StatsAPIRequest(env.BSM_API_KEY);

  const battingStats = statsRequest.getStatisticsForLeagueEntry<BattingStatisticsEntry>(Number(leagueEntry), StatsType.batting);
  const pitchingStats = statsRequest.getStatisticsForLeagueEntry<PitchingStatisticsEntry>(Number(leagueEntry), StatsType.pitching);
  const fieldingStats = statsRequest.getStatisticsForLeagueEntry<FieldingStatisticsEntry>(Number(leagueEntry), StatsType.fielding);

  return {
    teamName: teamName,
    battingStats: battingStats,
    pitchingStats: pitchingStats,
    fieldingStats: fieldingStats,
  };
};