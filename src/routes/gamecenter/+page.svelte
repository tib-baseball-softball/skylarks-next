<script lang="ts">
  import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
  import {ProgressBar, SlideToggle, Tab, TabGroup,} from "@skeletonlabs/skeleton";
  import {Gameday} from "bsm.js";
  import {preferences} from "$lib/stores";
  import LeagueFilter from "$lib/components/utility/LeagueFilter.svelte";
  import {goto} from "$app/navigation";
  import {browser} from "$app/environment";
  import GamecenterMatchSection from "$lib/components/match/GamecenterMatchSection.svelte";
  import type {PageProps} from "./$types";
  import ReloadUponPreferenceChange from "$lib/components/navigation/ReloadUponPreferenceChange.svelte";

  const DEFAULT_LEAGUE_GROUP_ID = 0;

  const reloadGameData = () => {
    if (browser) {
      let queryString = `?gameday=${$preferences.gameday}&season=${$preferences.selectedSeason}`;

      if ($preferences.leagueGroupID !== DEFAULT_LEAGUE_GROUP_ID) {
        queryString =
            queryString + `&leagueGroup=${$preferences.leagueGroupID}`;
      }
      goto(queryString);
    }
  };

  let {data}: PageProps = $props();
  let leagueGroups = $derived(data.leagueGroups);

  let showExternal = $state(false);
</script>

<ReloadUponPreferenceChange callback={reloadGameData}/>

<div class="my-2 md:flex justify-between items-start">
  <h1 class="h1">Gamecenter</h1>

  <div>
    <div class="flex gap-2">
      {#await leagueGroups then groups}
        <LeagueFilter leagueGroups={groups}/>
      {/await}
      <SeasonSelector/>
    </div>

    <div class="flex gap-2 my-4 justify-end">
      <SlideToggle
              size="sm"
              name="slide"
              active="bg-surface-900 dark:bg-tertiary-700"
              bind:checked={showExternal}
      />
      <p>Zeige externe Spiele</p>
    </div>
  </div>
</div>

<section class="mb-5 mt-3">
  <label id="gameday_label">
    Spieltag
    <TabGroup justify="justify-center" labelledby="gameday_label">
      <Tab
              bind:group={$preferences.gameday}
              name="tabPrevious"
              value={Gameday.previous}>Voriger
      </Tab
      >
      <Tab
              bind:group={$preferences.gameday}
              name="tabCurrent"
              value={Gameday.current}>Aktueller
      </Tab
      >
      <Tab
              bind:group={$preferences.gameday}
              name="tabNext"
              value={Gameday.next}>Nächster
      </Tab
      >
      <Tab
              bind:group={$preferences.gameday}
              name="tabAny"
              value={Gameday.any}>Alle
      </Tab
      >
      <!-- Tab Panels --->
      <svelte:fragment slot="panel"></svelte:fragment>
    </TabGroup>
  </label>
</section>

{#await data.streamed.matches}
  <p>Loading matches...</p>
  <ProgressBar/>
{:then matches}
  <GamecenterMatchSection {matches} {showExternal}/>
{:catch error}
  <p>error loading matches: {error.message}</p>
{/await}
