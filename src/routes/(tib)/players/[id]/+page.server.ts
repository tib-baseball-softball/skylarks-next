import {
  type BattingStatisticsEntry,
  type FieldingStatisticsEntry,
  type PitchingStatisticsEntry,
  StatsAPIRequest,
  StatsType,
} from "bsm.js";
import {env} from "$env/dynamic/private";
import {PlayerClient} from "$lib/tib/service/PlayerClient.ts";
import type {PageServerLoad} from "../../../../../.svelte-kit/types/src/routes";

export const load: PageServerLoad = async ({params, fetch}) => {
  const statsRequest = new StatsAPIRequest(env.BSM_API_KEY);

  const playerBattingStats = statsRequest.getStatisticsForPerson<BattingStatisticsEntry>(
      Number(params.id),
      StatsType.batting
  );
  const playerPitchingStats = statsRequest.getStatisticsForPerson<PitchingStatisticsEntry>(
      Number(params.id),
      StatsType.pitching
  );
  const playerFieldingStats = statsRequest.getStatisticsForPerson<FieldingStatisticsEntry>(
      Number(params.id),
      StatsType.fielding
  );

  const playerClient = new PlayerClient(fetch, env.BSM_API_KEY);
  const player = playerClient.fetchSinglePlayer(params.id);

  return {
    battingStats: playerBattingStats,
    pitchingStats: playerPitchingStats,
    fieldingStats: playerFieldingStats,
    player: player,
  };
};
