<script lang="ts">
  import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
  import {Progress, Switch} from "@skeletonlabs/skeleton-svelte";
  // @ts-ignore
  import {Tabs} from "bits-ui";
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
  function onGamedayChange(value: string) {
    //@ts-expect-error
    $preferences.gameday = value;
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
    <Tabs.Root
            value={$preferences.gameday}
            onValueChange={onGamedayChange}
            class=""
    >
      <Tabs.List
              class="tabs-list preset-tonal-surface"
      >
        <Tabs.Trigger
                value={Gameday.previous}
                class="tabs-trigger"
        >Previous
        </Tabs.Trigger>
        <Tabs.Trigger
                value={Gameday.current}
                class="tabs-trigger"
        >Current
        </Tabs.Trigger>
        <Tabs.Trigger
                value={Gameday.next}
                class="tabs-trigger"
        >Next
        </Tabs.Trigger>
        <Tabs.Trigger
                value={Gameday.any}
                class="tabs-trigger"
        >All
        </Tabs.Trigger>
      </Tabs.List>

      <Tabs.Content value={$preferences.gameday} class="pt-3">
        {@render tabsContent()}
      </Tabs.Content>
    </Tabs.Root>

    {#snippet tabsContent()}
      {#await data.streamed.matches}
        <p>Loading matches...</p>
        <Progress/>
      {:then matches}
        <GamecenterMatchSection {matches} {showExternal}/>
      {:catch error}
        <p>error loading matches: {error.message}</p>
      {/await}
    {/snippet}
  </label>
</section>
