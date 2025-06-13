<script lang="ts">
  import type {PageProps} from "./$types";
  import FavoriteTeamInfoCard from "$lib/components/favorite/FavoriteTeamInfoCard.svelte";
  import {preferences} from "$lib/globals.svelte.ts";
  import ClubTeamPicker from "$lib/components/utility/ClubTeamPicker.svelte";
  import {goto} from "$app/navigation";

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
  <h2 class="h2 mt-5 mb-1">Leagues</h2>
  <p class="text-sm font-light">A single team can be associated to several leagues per season in BSM.</p>

  {#await data.datasets}
    <p>Loading...</p>
  {:then datasets}
    {#each datasets as dataset}

    {:else }
      <p>No data found.</p>
    {/each}
  {:catch error}
    <p>Error: {error.message}</p>
  {/await}
</section>