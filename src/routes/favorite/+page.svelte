<script lang="ts">
  import type {PageProps} from "./$types";
  import FavoriteTeamInfoCard from "$lib/components/favorite/FavoriteTeamInfoCard.svelte";
  import {preferences} from "$lib/globals.svelte.ts";
  import ClubTeamPicker from "$lib/components/utility/ClubTeamPicker.svelte";
  import {goto} from "$app/navigation";
  import MatchTeaserCard from "$lib/components/match/MatchTeaserCard.svelte";
  import StandingsTable from "$lib/components/table/StandingsTable.svelte";
  import LeagueInfoCard from "$lib/components/favorite/LeagueInfoCard.svelte";
  import PlayoffSheet from "$lib/components/favorite/PlayoffSheet.svelte";
  import StreakGraphSection from "$lib/components/favorite/StreakGraphSection.svelte";

  let {data}: PageProps = $props();
  const clubTeams = $derived(data.clubTeams ?? []);

  const favoriteTeam = $derived(clubTeams.find(clubTeam => clubTeam.team.id === preferences.current.favoriteTeamID));

  function reloadOnTeamChange() {
    let queryString = `?team=${preferences.current.favoriteTeamID}&season=${preferences.current.selectedSeason}`;
    goto(queryString);
  }
</script>

<div class="lg:flex justify-between items-center">
  <h1 class="h1 shrink-0">Favorite Team</h1>

  <div class="grow-0">
    <ClubTeamPicker clubTeams={clubTeams} onChange={reloadOnTeamChange}/>
  </div>
</div>

<section class="max-w-md">
  <h2 class="h2 my-3">Quick Details</h2>
  {#if preferences.current.favoriteTeamID !== 0 && favoriteTeam}
    <FavoriteTeamInfoCard clubTeam={favoriteTeam}/>
  {:else }
    <p>You have not selected a favorite team yet.</p>
  {/if}
</section>

{#if favoriteTeam?.team.league_entries?.length ?? 0 > 0}
  <section>
    <h2 class="h2 mt-5 mb-1">Team Stats Links</h2>
    {#each favoriteTeam?.team.league_entries ?? [] as entry}
      <a href="/league_entries/{entry.id}?team={favoriteTeam?.team.name}"
         class="btn preset-filled-primary-500 me-1 my-2"
         title="to stats page for league entry {entry.league.name}">
        {entry.league.name}
      </a>
    {/each}
  </section>
{/if}

{#if favoriteTeam}
  <section>
    <h2 class="h2 mt-5 mb-1">Data per League</h2>
    <p class="text-sm font-light">A single team can be associated to several leagues per season in BSM.</p>
  </section>

  {#await data.datasets}
    <p>Loading...</p>
  {:then datasets}
    {#each datasets as dataset}
      <section class="max-w-md">
        <h3 class="h3 mb-3">{dataset.league_group.name}</h3>

        {#if dataset.table_row}
          <LeagueInfoCard leagueGroup={dataset.league_group} tableRow={dataset.table_row}/>
        {:else}
          <p>The team "{favoriteTeam?.team.name}" does not have a league entry in this league.</p>
        {/if}
      </section>

      <section>
        <h4 class="h4 mb-3">Standings</h4>

        <StandingsTable table={dataset.table}/>
      </section>

      {#if dataset.playoff_series && (dataset.playoff_series?.series_games?.length ?? 0) > 0}
        <section>
          <h4 class="h4 mt-5 mb-1">Playoffs</h4>
          <PlayoffSheet playoffSeries={dataset.playoff_series}></PlayoffSheet>
        </section>
      {/if}

      {#if dataset.streak_data?.length ?? 0 > 0}
        <section>
          <h4 class="h4 mb-3">Graphs and Mood</h4>
          <StreakGraphSection {dataset}/>
        </section>
      {/if}

      <section>
        <h4 class="h4 mb-3">Games</h4>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-3 lg:gap-6 xl:gap-12">
          <div>
            <h5 class="h5 my-3">Previous Game</h5>
            {#if dataset.last_game}
              <MatchTeaserCard match={dataset.last_game}/>
            {:else}
              <p class="my-3">No previous game found.</p>
            {/if}
          </div>

          <div>
            <h5 class="h5 my-3">Next Game</h5>
            {#if dataset.next_game}
              <MatchTeaserCard match={dataset.next_game}/>
            {:else}
              <p class="my-3">No next game found.</p>
            {/if}
          </div>
        </div>
      </section>

    {:else }
      <p>No data found.</p>
    {/each}
  {:catch error}
    <p>Error: {error.message}</p>
  {/await}
{/if}