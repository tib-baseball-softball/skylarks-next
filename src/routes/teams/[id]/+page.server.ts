import {error} from "@sveltejs/kit";
import type {ClubTeam} from "bsm.js/dist/model/ClubTeam";
import {
    type BattingStatisticsEntry,
    type FieldingStatisticsEntry,
    type PitchingStatisticsEntry,
    StatsAPIRequest,
    StatsType
} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";

export async function load({ parent, params}) {
    const data = await parent()
    const clubTeams = await data.clubTeams ?? []
    const clubTeam: ClubTeam | undefined = clubTeams.find((clubTeam) => clubTeam.id === Number(params.id))

    if (!clubTeam) {
        throw error(404)
    }

    const statsRequest = new StatsAPIRequest(BSM_API_KEY)
    const leagueEntry = clubTeam.team.league_entries.at(0)

    if (!leagueEntry) {
        return {
            clubTeam: clubTeam
        }
    }

    const battingStats = statsRequest.getStatisticsForLeagueEntry<BattingStatisticsEntry>(leagueEntry.id, StatsType.batting)
    const pitchingStats = statsRequest.getStatisticsForLeagueEntry<PitchingStatisticsEntry>(leagueEntry.id, StatsType.pitching)
    const fieldingStats = statsRequest.getStatisticsForLeagueEntry<FieldingStatisticsEntry>(leagueEntry.id, StatsType.fielding)

    return {
        clubTeam: clubTeam,
        battingStats: battingStats,
        pitchingStats: pitchingStats,
        fieldingStats: fieldingStats,
    }
}