import type {Match} from "../../../../bsm.js/src/model/Match";
import {GameWinner} from "$lib/enum/GameWinner";
import {MatchState} from "$lib/enum/MatchState";
import {PUBLIC_TEAM_NAME} from "$env/static/public";

export class MatchUtility {
    public static getWinnerForMatch(match: Match): GameWinner {
        if ((match.home_runs ?? 0) > (match.away_runs ?? 0)) {
            return GameWinner.home
        } else if ((match.away_runs ?? 0) > (match.home_runs ?? 0)) {
            return GameWinner.away
        } else {
            return GameWinner.none
        }
    }

    public static getMatchState(match: Match): MatchState {
        if (match.state === "planned") {
            return MatchState.notYetPlayed
        }

        if (match.state === "cancelled" || match.state === "canceled") {
            return MatchState.cancelled
        }

        if (match.home_team_name.includes(PUBLIC_TEAM_NAME) && match.away_team_name.includes(PUBLIC_TEAM_NAME)) {
            return MatchState.derby
        }

        const winner = MatchUtility.getWinnerForMatch(match)
        if (
            (winner === GameWinner.home && match.home_team_name.includes(PUBLIC_TEAM_NAME)) ||
            (winner === GameWinner.away && match.away_team_name.includes(PUBLIC_TEAM_NAME))
        ) {
            return MatchState.won
        } else if (
            (winner === GameWinner.away && match.home_team_name.includes(PUBLIC_TEAM_NAME)) ||
            (winner === GameWinner.home && match.away_team_name.includes(PUBLIC_TEAM_NAME))
        ) {
             return MatchState.lost
        } else {
            return MatchState.final
        }
    }
}