import {type BattingStatisticsEntry, StatsAPIRequest, StatsType} from "bsm.js";
import {BSM_API_KEY, SKYLARKS_API_AUTH_HEADER} from "$env/static/private";
import {PUBLIC_BACKEND_URL} from "$env/static/public";

export async  function load({params, fetch}) {
    const statsRequest = new StatsAPIRequest(BSM_API_KEY)

    const playerBattingStats = statsRequest.getStatisticsForPerson<BattingStatisticsEntry>(Number(params.id), StatsType.batting)
    const playerPitchingStats = statsRequest.getStatisticsForPerson<BattingStatisticsEntry>(Number(params.id), StatsType.pitching)
    const playerFieldingStats = statsRequest.getStatisticsForPerson<BattingStatisticsEntry>(Number(params.id), StatsType.fielding)

    const response = await fetch(`${PUBLIC_BACKEND_URL}/api/v2/players?bsmid=${params.id}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "X-Authorization": SKYLARKS_API_AUTH_HEADER
        }
    })
    const player = response.json()

    return {
        battingStats: playerBattingStats,
        pitchingStats: playerPitchingStats,
        fieldingStats: playerFieldingStats,
        player: player
    }
}