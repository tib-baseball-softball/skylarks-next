import {
  StatsAPIRequest,
  StatsType,
  type BattingStatisticsEntry,
  type PitchingStatisticsEntry,
  type FieldingStatisticsEntry,
} from "bsm.js";
import type {PageLoad} from "./$types";
import {client} from "$lib/dp/client.svelte.ts";
import {RELAY_URL} from "$lib/tib/globals.svelte.ts";
import {error} from "@sveltejs/kit";

export const load: PageLoad = async ({ params, url, fetch }) => {
  const leagueEntry = params.entry;
  const teamName = url.searchParams.get("team") ?? undefined;

  if (!leagueEntry) {
    throw error(400, { message: "No league entry provided." });
  }

  const statsRequest = new StatsAPIRequest("");
  statsRequest.setAPIURL(RELAY_URL);

  // TODO: Verify the exact BSM endpoint paths for statistics by league entry
  const battingURL = statsRequest.buildRequestURL(`league_entries/${leagueEntry}/statistics/${StatsType.batting}.json`, []);
  const pitchingURL = statsRequest.buildRequestURL(`league_entries/${leagueEntry}/statistics/${StatsType.pitching}.json`, []);
  const fieldingURL = statsRequest.buildRequestURL(`league_entries/${leagueEntry}/statistics/${StatsType.fielding}.json`, []);

  const battingStats = client.send<BattingStatisticsEntry[]>(battingURL.pathname + battingURL.search, {
    fetch,
    requestKey: `leagueEntry-batting-${leagueEntry}`,
  });
  const pitchingStats = client.send<PitchingStatisticsEntry[]>(pitchingURL.pathname + pitchingURL.search, {
    fetch,
    requestKey: `leagueEntry-pitching-${leagueEntry}`,
  });
  const fieldingStats = client.send<FieldingStatisticsEntry[]>(fieldingURL.pathname + fieldingURL.search, {
    fetch,
    requestKey: `leagueEntry-fielding-${leagueEntry}`,
  });

  return {
    teamName,
    battingStats,
    pitchingStats,
    fieldingStats,
  };
};
