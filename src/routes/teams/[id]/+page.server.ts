import {error} from "@sveltejs/kit";
import type {ClubTeam} from "bsm.js";
import {
    type BattingStatisticsEntry,
    type FieldingStatisticsEntry,
    type PitchingStatisticsEntry,
    StatsAPIRequest,
    ClubTeamsAPIRequest,
    StatsType, TablesAPIRequest
} from "bsm.js";
import {BSM_API_KEY} from "$env/static/private";
import {PUBLIC_CLUB_ID} from "$env/static/public";

export async function load({ parent, params, url }) {
    const data = await parent()
    let clubTeams = await data.clubTeams

    // this is kind of hacky because I could not find an API call to get a single team, so if the collection is empty
    // we have to load them all again...could also be a skill issue with SvelteKit
    if (clubTeams?.length === 0) {
        const clubTeamRequest = new  ClubTeamsAPIRequest(BSM_API_KEY)
        const season = url.searchParams.get("season")?.toString() ?? clubTeamRequest.defaultSeason
        clubTeams = await clubTeamRequest.getTeamsForClub(Number(PUBLIC_CLUB_ID), Number(season))
    }

    const clubTeam: ClubTeam | undefined = clubTeams?.find((clubTeam) => clubTeam.id === Number(params.id))

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

    const leagueGroups = await data.leagueGroups
    const leagueGroup = leagueGroups?.find((group) => group.league.id === leagueEntry.league.id)

    const tableRequest = new TablesAPIRequest(BSM_API_KEY)

    return {
        clubTeam: clubTeam,
        battingStats: battingStats,
        pitchingStats: pitchingStats,
        fieldingStats: fieldingStats,
        table: leagueGroup ? tableRequest.getSingleTable(leagueGroup.id) : null,
    }
}