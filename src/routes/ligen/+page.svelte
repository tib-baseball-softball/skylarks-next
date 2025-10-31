<script lang="ts">
  import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
  import ReloadUponPreferenceChange from "$lib/components/navigation/ReloadUponPreferenceChange.svelte";
  import {browser} from "$app/environment";
  import {goto} from "$app/navigation";
  import {preferences} from "$lib/globals.svelte.ts";
  import type {PageProps} from "./$types";
  import ProgressRing from "$lib/components/utility/ProgressRing.svelte";

  const reload = () => {
    if (browser) {
      let queryString = `?season=${preferences.current.selectedSeason}`;

      goto(queryString);
    }
  };

  let {data}: PageProps = $props();
</script>

<ReloadUponPreferenceChange callback={reload}/>

<div class="my-2 md:flex justify-between items-center">
  <h1 class="h1 mb-3">Leagues</h1>
  <div>
    <SeasonSelector/>
  </div>
</div>

{#await data.leagueGroups}
  <p>Loading Leagues...</p>
  <ProgressRing/>

{:then leagueGroups}

  <ul class="list mt-5 flex flex-col gap-3">
    {#each leagueGroups as leagueGroup}
      <li class="preset-tonal-surface dark:preset-filled-surface-300-700 p-3 min-h-14 rounded-base">
        <a href="/ligen/{leagueGroup.id}">
          <span class="badge preset-tonal-primary border border-primary-500 w-20">{leagueGroup.acronym}</span>
          <span class="flex-auto ms-3">{leagueGroup.name}</span>
        </a>
      </li>
    {/each}
  </ul>

{:catch error}
  <p>error loading: {error.message}</p>
{/await}