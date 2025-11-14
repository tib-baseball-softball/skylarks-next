import {error} from "@sveltejs/kit";
import type {LeagueGroup} from "bsm.js";
import {env} from "$env/dynamic/public";
import {client} from "$lib/dp/client.svelte.js";
import type {LeaderboardSummary} from "$lib/tib/types/LeaderboardData.ts";
import type {PageLoad} from "./$types";

/**
 * Gets statistics for league leaders in all relevant categories.
 */
export const load: PageLoad = async ({parent, params, url, fetch}) => {
  const data = await parent();
  const leagueGroups = await data.leagueGroups;
  let leagueGroup: LeagueGroup | undefined = leagueGroups.find(
      (leagueGroup) => leagueGroup.id === Number(params.id)
  );

  if (!leagueGroup) {
    leagueGroup = await client.send<LeagueGroup>("/api/bsm/relay", {
      fetch: fetch,
      query: {
        url: `https://bsm.baseball-softball.de/league_groups/${params.id}.json`,
        club: env.PUBLIC_CLUB_ID,
      },
      requestKey: "leagueGroupStatsRoute",
    });
  }

  if (!leagueGroup) {
    throw error(404, "leagueGroup not found");
  }

  const leaderboardData = client.send<LeaderboardSummary>(`api/bsm/relay/top10/${leagueGroup.id}`, {
    fetch: fetch,
    query: {
      statsType: url.searchParams.get("statsType") ?? "",
    },
  });

  return {
    leagueGroup: leagueGroup,
    leaderboardData: leaderboardData,
  };
};
