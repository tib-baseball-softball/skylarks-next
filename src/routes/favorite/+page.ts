import type {PageLoad} from './$types';

export const load: PageLoad = async ({ parent }) => {
  const data = await parent()
  return {
    clubTeams: await data.clubTeams,
    leagueGroups: await data.leagueGroups,
    clubs: data.clubs,
    teams: data.teams,
  }
}