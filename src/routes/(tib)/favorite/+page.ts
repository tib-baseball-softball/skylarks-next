import {browser} from "$app/environment";
import {client} from "$lib/dp/client.svelte.ts";
import {preferences} from "$lib/tib/globals.svelte.ts";
import type {HomeDataset} from "$lib/tib/types/HomeDataset.ts";
import type {PageLoad} from "../../../../.svelte-kit/types/src/routes";

export const load: PageLoad = async ({parent, url, fetch}) => {
  let teamId = Number(url.searchParams.get("team"));
  let season = Number(url.searchParams.get("season"));

  if (browser) {
    teamId = preferences.current.favoriteTeamID;
    season = preferences.current.selectedSeason;
  }

  let datasets: Promise<HomeDataset[]>;

  if (browser && Number(teamId) > 0) {
    try {
      datasets = client.send<HomeDataset[]>("/api/team/favorite", {
        fetch: fetch,
        query: {
          team: teamId,
          season: season,
        },
        requestKey: `favorite-${teamId}-${season}`,
      });
    } catch (e) {
      datasets = new Promise(() => []);
    }
  } else {
    datasets = new Promise(() => []);
  }

  const data = await parent();

  return {
    clubTeams: await data.clubTeams,
    leagueGroups: await data.leagueGroups,
    datasets: datasets,
    clubs: data.clubs,
    teams: data.teams,
  };
};
