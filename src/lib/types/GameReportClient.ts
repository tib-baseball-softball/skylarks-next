import type {GameReport} from "$lib/model/GameReport.ts";
import {TYPO3Client} from "$lib/service/TYPO3Client.ts";
import {env} from "$env/dynamic/public";

export class GameReportClient extends TYPO3Client {
  /**
   * Loads single
   * @param bsmMatchID
   * @throws Error
   */
  public async loadSingleGameReportForBSMMatchID(bsmMatchID: string): Promise<GameReport> {
    const url = this.buildRequestURL("api/v2/gamereports", [
      ["bsmMatchID", bsmMatchID]
    ]);

    if (env.PUBLIC_APPLICATION_CONTEXT === "Development") {
      console.log(url);
    }

    const response = await this.fetch(url, {
      headers: this.AUTH_HEADERS
    });
    const gameReports: GameReport[] = await response.json();
    const gameReport = gameReports.at(0);

    if (gameReport) {
      return gameReport;
    } else {
      throw new Error("Game report could not be parsed from response.");
    }
  }
}