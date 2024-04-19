import type {Match} from "../../../../bsm.js/src/model/Match";
import {GameWinner} from "$lib/enum/GameWinner";

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
}