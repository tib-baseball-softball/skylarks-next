import {TYPO3Client} from "$lib/service/TYPO3Client.ts";
import type {Player} from "$lib/model/Player.ts";

export class PlayerClient extends TYPO3Client {
  /**
   * Fetches a single player by their unique identifier.
   *
   * @param {string} bsmID - The unique identifier of the player to fetch.
   * @return {Promise<Player | undefined>} A promise that resolves to the player object if found, or undefined if not found.
   */
  async fetchSinglePlayer(bsmID: string): Promise<Player | undefined> {
    const url = this.buildRequestURL("/api/v2/players", {
      bsmid: bsmID
    });

    const response: Response = await this.fetch(url, {
      headers: this.AUTH_HEADERS
    });

    let player: Player | undefined;
    const players: Player[] = await response.json();

    if (players && players.length > 0) {
      player = players.at(0);
    }

    return player;
  }

  async fetchPlayersByFilters(team?: number): Promise<Player[] | undefined> {
    let query: PlayerQuery = {};
    if (team) {
      query.team = String(team);
    }
    const url = this.buildRequestURL("/api/v2/players", query);

    const response: Response = await this.fetch(url, {
      headers: this.AUTH_HEADERS
    });

    return await response.json();
  }
}

type PlayerQuery = {
  id?: string,
  bsmid?: string,
  team?: string,
}