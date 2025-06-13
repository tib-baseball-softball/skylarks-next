<script lang="ts">
  import type {PageProps} from "./$types";
  import FavoriteTeamInfoCard from "$lib/components/favorite/FavoriteTeamInfoCard.svelte";
  import {preferences} from "$lib/globals.svelte.ts";
  import ClubTeamPicker from "$lib/components/utility/ClubTeamPicker.svelte";
  import {goto} from "$app/navigation";
  import MatchTeaserCard from "$lib/components/match/MatchTeaserCard.svelte";
  import StandingsTable from "$lib/components/table/StandingsTable.svelte";
  import LeagueInfoCard from "$lib/components/favorite/LeagueInfoCard.svelte";

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

  <ClubTeamPicker clubTeams={clubTeams} onChange={reloadOnTeamChange}/>
</div>

<section>
  <h2 class="h2 my-3">Quick Details</h2>
  {#if preferences.current.favoriteTeamID !== 0 && favoriteTeam}
    <FavoriteTeamInfoCard clubTeam={favoriteTeam}/>
  {:else }
    <p>You have not selected a favorite team yet.</p>
  {/if}
</section>

<section>
  <h2 class="h2 mt-5 mb-1">Team Stats</h2>
  NONE
</section>

<section>
  <h2 class="h2 mt-5 mb-1">Data per League</h2>
  <p class="text-sm font-light">A single team can be associated to several leagues per season in BSM.</p>
</section>

{#await data.datasets}
  <p>Loading...</p>
{:then datasets}
  {#each datasets as dataset}
    <section>
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

    <section>
      <h4 class="h4 mb-3">Games</h4>

      {#if dataset.last_game}
        <h5 class="h5 my-3">Previous Game</h5>
        <MatchTeaserCard match={dataset.last_game}/>
      {:else}
        <p>No previous game found.</p>
      {/if}

      {#if dataset.next_game}
        <h5 class="h5 my-3">Next Game</h5>
        <MatchTeaserCard match={dataset.next_game}/>
      {:else}
        <p>No next game found.</p>
      {/if}
    </section>

  {:else }
    <p>No data found.</p>
  {/each}
{:catch error}
  <p>Error: {error.message}</p>
{/await}
