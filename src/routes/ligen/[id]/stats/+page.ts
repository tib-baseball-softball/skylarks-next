import {type LeagueGroup} from "bsm.js";
import {error} from "@sveltejs/kit";
import {client} from "$lib/pocketbase/index.svelte";
import type {LeaderboardSummary} from "$lib/model/LeaderboardData";
import {env} from "$env/dynamic/public";

/**
 * Gets statistics for league leaders in all relevant categories.
 */
export async function load({parent, params, url, fetch}) {
  const data = await parent();
  const leagueGroups = await data.leagueGroups;
  let leagueGroup: LeagueGroup | undefined = leagueGroups.find((leagueGroup) => leagueGroup.id === Number(params.id));

  if (!leagueGroup) {
    leagueGroup = await client.send<LeagueGroup>("/api/bsm/relay", {
      fetch: fetch,
      query: {
        url: `https://bsm.baseball-softball.de/league_groups/${params.id}.json`,
        club: env.PUBLIC_CLUB_ID,
      },
      requestKey: "leagueGroupStatsRoute"
    });
  }

  if (!leagueGroup) {
    throw error(404, "leagueGroup not found");
  }

  const leaderboardData = client.send<LeaderboardSummary>(`api/bsm/relay/top10/${leagueGroup.id}`, {
    query: {
      statsType: url.searchParams.get("statsType") ?? "",
    }
  });

  return {
    leagueGroup: leagueGroup,
    leaderboardData: leaderboardData,
  };
}