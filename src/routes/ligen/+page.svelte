<script lang="ts">
    import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
    import {ProgressBar} from "@skeletonlabs/skeleton";
    import ReloadUponPreferenceChange from "$lib/components/navigation/ReloadUponPreferenceChange.svelte";
    import {browser} from "$app/environment";
    import {goto} from "$app/navigation";
    import {preferences} from "$lib/stores.ts";
    import type {PageProps} from "./$types";

    const reload = () => {
    if (browser) {
      let queryString = `?season=${$preferences.selectedSeason}`;

      goto(queryString);
    }
  };

  let {data}: PageProps = $props();
</script>

<ReloadUponPreferenceChange callback={reload}/>

<div class="my-2 md:flex justify-between items-center">
  <h1 class="h1 mb-3">Ligen</h1>
  <div>
    <SeasonSelector/>
  </div>
</div>

{#await data.leagueGroups}
  <p>Lade Ligen...</p>
  <ProgressBar/>

{:then leagueGroups}

  <ul class="list mt-5 flex flex-col gap-3">
    {#each leagueGroups as leagueGroup}
      <li class="variant-soft-surface p-3 min-h-14">
        <a href="/ligen/{leagueGroup.id}">
          <span class="badge variant-ghost-primary w-20">{leagueGroup.acronym}</span>
          <span class="flex-auto ms-3">{leagueGroup.name}</span>
        </a>
      </li>
    {/each}
  </ul>

{:catch error}
  <p>error loading: {error.message}</p>
{/await}