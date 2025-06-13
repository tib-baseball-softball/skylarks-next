import type {PageLoad} from './$types';
import {client} from "$lib/pocketbase/index.svelte.ts";
import type {HomeDataset} from "$lib/types/HomeDataset.ts";
import {browser} from "$app/environment";
import {preferences} from "$lib/globals.svelte.ts";

export const load: PageLoad = async ({parent, url, fetch}) => {
  const data = await parent();
  let teamId = Number(url.searchParams.get("team"));
  let season = Number(url.searchParams.get("season"));

  if (browser) {
    teamId = preferences.current.favoriteTeamID
    season = preferences.current.selectedSeason
  }

  let datasets: Promise<HomeDataset[]> = new Promise(() => []);

  if (browser && Number(teamId) > 0) {
    datasets = client.send<HomeDataset[]>("/api/team/favorite", {
      fetch: fetch,
      query: {
        team: teamId,
        season: season,
      },
      requestKey: `favorite-${teamId}-${season}`
    });
  }

  return {
    clubTeams: await data.clubTeams,
    leagueGroups: await data.leagueGroups,
    datasets: datasets,
    clubs: data.clubs,
    teams: data.teams,
  };
};