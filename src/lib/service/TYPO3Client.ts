import { AbstractAPIRequest } from "bsm.js"
import type { Fetch } from "$lib/types/Fetch.ts"
import { env as publicEnv } from "$env/dynamic/public"
import { env } from "$env/dynamic/private"

export class TYPO3Client extends AbstractAPIRequest {
  protected readonly fetch: Fetch
  protected readonly API_URL = publicEnv.PUBLIC_TYPO3_URL

  protected readonly AUTH_HEADERS = {
    "Content-Type": "application/json",
    "X-Authorization": env.SKYLARKS_API_AUTH_HEADER,
  }

  constructor(fetch: Fetch, apiKey: string) {
    super(apiKey)
    this.fetch = fetch
  }
}
