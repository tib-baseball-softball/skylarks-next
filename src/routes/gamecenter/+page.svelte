<script lang="ts">
  import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
  import {Progress, Switch, Tabs} from "@skeletonlabs/skeleton-svelte";
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

  /**
   *  enum <=> string conversion necessary
   */
  function onGamedayChange(e: { value: string }) {
    //@ts-expect-error
    $preferences.gameday = e.value;
  }
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
      <Switch
              name="slide"
              controlActive="bg-surface-900 dark:bg-tertiary-700"
              onCheckedChange={(e) => (showExternal = e.checked)}
      />
      <p>Show external games</p>
    </div>
  </div>
</div>

<section class="mb-5 mt-3">
  <label id="gameday_label" class="label">
    Gameday
    <Tabs listJustify="justify-center" onValueChange={onGamedayChange}>

      {#snippet list()}
        <Tabs.Control value={Gameday.previous}>Previous</Tabs.Control>
        <Tabs.Control value={Gameday.current}>Current</Tabs.Control>
        <Tabs.Control value={Gameday.next}>Next</Tabs.Control>
        <Tabs.Control value={Gameday.any}>All</Tabs.Control>
      {/snippet}

      {#snippet content()}
        {#await data.streamed.matches}
          <p>Loading matches...</p>
          <Progress/>
        {:then matches}
          <GamecenterMatchSection {matches} {showExternal}/>
        {:catch error}
          <p>error loading matches: {error.message}</p>
        {/await}
      {/snippet}
    </Tabs>
  </label>
</section>
