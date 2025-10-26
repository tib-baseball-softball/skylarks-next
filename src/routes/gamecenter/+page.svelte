<script lang="ts">
  import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
  import Switch from "$lib/components/utility/Switch.svelte";
  // @ts-ignore
  import {Tabs} from "bits-ui";
  import {Gameday} from "bsm.js";
  import {preferences} from "$lib/globals.svelte.ts";
  import LeagueFilter from "$lib/components/utility/LeagueFilter.svelte";
  import {goto} from "$app/navigation";
  import {browser} from "$app/environment";
  import GamecenterMatchSection from "$lib/components/match/GamecenterMatchSection.svelte";
  import type {PageProps} from "./$types";
  import ReloadUponPreferenceChange from "$lib/components/navigation/ReloadUponPreferenceChange.svelte";
  import ProgressRing from "$lib/components/utility/ProgressRing.svelte";

  const DEFAULT_LEAGUE_GROUP_ID = 0;

  const reloadGameData = () => {
    if (browser) {
      let queryString = `?gameday=${preferences.current.gameday}&season=${preferences.current.selectedSeason}`;

      if (preferences.current.leagueGroupID !== DEFAULT_LEAGUE_GROUP_ID) {
        queryString =
            queryString + `&leagueGroup=${preferences.current.leagueGroupID}`;
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
  function onGamedayChange(value: string) {
    //@ts-expect-error
    preferences.current.gameday = value;
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
              controlActive="bg-surface-900 dark:bg-tertiary-700"
              name="slide"
              onCheckedChange={(e) => (showExternal = e.checked)}
      />
      <p>Show external games</p>
    </div>
  </div>
</div>

<section class="mb-5 mt-3">
  <label class="label" id="gameday_label">
    Gameday
    <Tabs.Root
            class=""
            onValueChange={onGamedayChange}
            value={preferences.current.gameday}
    >
      <Tabs.List
              class="tabs-list preset-tonal-surface"
      >
        <Tabs.Trigger
                class="tabs-trigger btn"
                value={Gameday.previous}
        >Previous
        </Tabs.Trigger>
        <Tabs.Trigger
                class="tabs-trigger btn"
                value={Gameday.current}
        >Current
        </Tabs.Trigger>
        <Tabs.Trigger
                class="tabs-trigger btn"
                value={Gameday.next}
        >Next
        </Tabs.Trigger>
        <Tabs.Trigger
                class="tabs-trigger btn"
                value={Gameday.any}
        >All
        </Tabs.Trigger>
      </Tabs.List>

      <Tabs.Content class="pt-3" value={preferences.current.gameday}>
        {#await data.streamed.matches}
          <p>Loading matches...</p>
          <ProgressRing/>
        {:then matches}
          <GamecenterMatchSection {matches} {showExternal}/>
        {:catch error}
          <p>error loading matches: {error.message}</p>
        {/await}
      </Tabs.Content>
    </Tabs.Root>

  </label>
</section>
