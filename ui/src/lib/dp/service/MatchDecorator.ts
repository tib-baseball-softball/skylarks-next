import type { Match } from "bsm.js";
import { GameWinner } from "$lib/dp/enum/GameWinner.ts";
import { MatchState } from "$lib/dp/enum/MatchState.ts";

/**
 * This class exists so that I can have methods on the BSM interface `Match` which is a POJO when deserialized from JSON.
 */
export class MatchDecorator {
  private readonly match: Match;

  constructor(match: Match) {
    this.match = match;
  }

  /**
   *
   * @param teamName name to compare - expects normalized format (no manipulation is done)
   */
  public isDerby(teamName: string): boolean {
    return (
      this.match.home_team_name.toLowerCase().includes(teamName) &&
      this.match.away_team_name.toLowerCase().includes(teamName)
    );
  }

  public getWinnerForMatch(): GameWinner {
    if (
      this.match?.home_runs === undefined ||
      this.match?.away_runs === undefined
    ) {
      return GameWinner.none;
    }

    if (this.match.home_runs > this.match.away_runs) {
      return GameWinner.home;
    } else if (this.match.away_runs > this.match.home_runs) {
      return GameWinner.away;
    } else {
      return GameWinner.none;
    }
  }

  public getMatchState(teamName: string): MatchState {
    const normalizedTeamName = teamName.toLowerCase();

    if (this.match.state === "planned") {
      return MatchState.notYetPlayed;
    }

    if (this.match.state === "cancelled" || this.match.state === "canceled") {
      return MatchState.cancelled;
    }

    if (this.isDerby(normalizedTeamName)) {
      return MatchState.derby;
    }

    const winner = this.getWinnerForMatch();

    if (
      (winner === GameWinner.home &&
        this.match.home_team_name.toLowerCase().includes(normalizedTeamName)) ||
      (winner === GameWinner.away &&
        this.match.away_team_name.toLowerCase().includes(normalizedTeamName))
    ) {
      return MatchState.won;
    } else if (
      (winner === GameWinner.away &&
        this.match.home_team_name.toLowerCase().includes(normalizedTeamName)) ||
      (winner === GameWinner.home &&
        this.match.away_team_name.toLowerCase().includes(normalizedTeamName))
    ) {
      return MatchState.lost;
    } else {
      return MatchState.final;
    }
  }

  public isPlayoffGame(): boolean {
    return this.match.match_id.includes("PO");
  }
}
