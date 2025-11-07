import { TYPO3Client } from "$lib/service/TYPO3Client.ts"
import type { TiBTeam } from "$lib/model/TiBTeam.ts"

export class TeamClient extends TYPO3Client {
  async fetchTeamByFilter(query: TeamQuery): Promise<TiBTeam[] | undefined> {
    const url = this.buildRequestURL("/api/v2/teams", query)

    const response: Response = await this.fetch(url, {
      headers: this.AUTH_HEADERS,
    })

    return await response.json()
  }
}

type TeamQuery = {
  id?: string
  league?: string
}
