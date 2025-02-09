import {type BattingStatisticsEntry, StatsAPIRequest, StatsType} from "bsm.js";
import {env} from "$env/dynamic/private";
import {env as publicEnv} from "$env/dynamic/public";

export async function load({params, fetch}) {
  const statsRequest = new StatsAPIRequest(env.BSM_API_KEY);

  const playerBattingStats = statsRequest.getStatisticsForPerson<BattingStatisticsEntry>(Number(params.id), StatsType.batting);
  const playerPitchingStats = statsRequest.getStatisticsForPerson<BattingStatisticsEntry>(Number(params.id), StatsType.pitching);
  const playerFieldingStats = statsRequest.getStatisticsForPerson<BattingStatisticsEntry>(Number(params.id), StatsType.fielding);

  const response = await fetch(`${publicEnv.PUBLIC_TYPO3_URL}/api/v2/players?bsmid=${params.id}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      "X-Authorization": env.SKYLARKS_API_AUTH_HEADER
    }
  });
  const player = response.json();

  return {
    battingStats: playerBattingStats,
    pitchingStats: playerPitchingStats,
    fieldingStats: playerFieldingStats,
    player: player
  };
}