<script lang="ts">
  import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
  import {Progress} from "@skeletonlabs/skeleton-svelte";
  import ReloadUponPreferenceChange from "$lib/components/navigation/ReloadUponPreferenceChange.svelte";
  import {preferences} from "$lib/globals.svelte.ts";
  import {browser} from "$app/environment";
  import {goto} from "$app/navigation";
  import type {PageProps} from "./$types";

  const reload = () => {
    if (browser) {
      let queryString = `?season=${preferences.current.selectedSeason}`;

      goto(queryString);
    }
  };

  let {data}: PageProps = $props();
</script>

<div class="my-2 md:flex justify-between items-center">
  <h1 class="h1 mb-3">Teams</h1>
  <div>
    <SeasonSelector/>
  </div>
</div>

<ReloadUponPreferenceChange callback={reload}/>

{#await data.clubTeams}
  <p>Loading Teams...</p>
  <Progress/>

{:then clubTeams}

  <ul class="list mt-5 flex flex-col gap-3">
    {#each clubTeams ?? [] as clubTeam}
      <li class="preset-tonal-surface dark:preset-filled-surface-300-700 p-3 min-h-14 rounded-base">
        <a href="/teams/{clubTeam.id}?season={preferences.current.selectedSeason}">
          <span class="badge preset-tonal-primary border border-primary-500 w-20">{clubTeam.team.league_entries.at(0)?.league.acronym}</span>
          <span class="flex-auto ms-3">{clubTeam.team.name}</span>
        </a>
      </li>
    {/each}
  </ul>

{:catch error}
  <p>error loading: {error.message}</p>
{/await}