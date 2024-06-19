import {type BattingStatisticsEntry, StatsAPIRequest, StatsType} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";

export async  function load({params}) {
    const statsRequest = new StatsAPIRequest(BSM_API_KEY)

    const playerBattingStats = statsRequest.getStatisticsForPerson<BattingStatisticsEntry>(Number(params.id), StatsType.batting)
    const playerPitchingStats = statsRequest.getStatisticsForPerson<BattingStatisticsEntry>(Number(params.id), StatsType.pitching)
    const playerFieldingStats = statsRequest.getStatisticsForPerson<BattingStatisticsEntry>(Number(params.id), StatsType.fielding)

    return {
        battingStats: playerBattingStats,
        pitchingStats: playerPitchingStats,
        fieldingStats: playerFieldingStats,
    }
}