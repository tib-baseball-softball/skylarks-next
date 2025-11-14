<script lang="ts">
  import {browser} from "$app/environment";
  import {goto} from "$app/navigation";
  import ReloadUponPreferenceChange from "$lib/tib/components/utils/ReloadUponPreferenceChange.svelte";
  import ProgressRing from "$lib/dp/components/utils/ProgressRing.svelte";
  import SeasonSelector from "$lib/tib/components/utils/SeasonSelector.svelte";
  import {preferences} from "$lib/tib/globals.svelte.ts";
  import type {PageProps} from "../../../../.svelte-kit/types/src/routes";

  const reload = () => {
    if (browser) {
      const queryString = `?season=${preferences.current.selectedSeason}`;

      goto(queryString);
    }
  };

  const {data}: PageProps = $props();
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
        <a href="/(tib)/ligen/{leagueGroup.id}">
          <span class="badge preset-tonal-primary border border-primary-500 w-20">{leagueGroup.acronym}</span>
          <span class="flex-auto ms-3">{leagueGroup.name}</span>
        </a>
      </li>
    {/each}
  </ul>

{:catch error}
  <p>error loading: {error.message}</p>
{/await}