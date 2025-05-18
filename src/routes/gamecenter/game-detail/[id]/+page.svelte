<script lang="ts">
  import {ProgressRing} from "@skeletonlabs/skeleton-svelte";
  // @ts-ignore
  import {Tabs} from "bits-ui";
  import MatchMainInfoSection from "$lib/components/match/MatchMainInfoSection.svelte";
  import MatchBoxscoreSection from "$lib/components/boxscore/MatchBoxscoreSection.svelte";
  import MatchDetailStatsCard from "$lib/components/match/MatchDetailStatsCard.svelte";
  import MatchDetailLocationCard from "$lib/components/match/MatchDetailLocationCard.svelte";
  import MatchDetailOfficialsCard from "$lib/components/match/MatchDetailOfficialsCard.svelte";
  import type {PageProps} from "./$types";
  import ContentFilteredUnavailable from "$lib/components/match/ContentFilteredUnavailable.svelte";
  import GameReport from "$lib/components/gameReport/GameReport.svelte";

  let {data}: PageProps = $props();

  let match = $derived(data.match);
  let tabSet: "gameData" | "boxscore" | "gameReport" | string = $state("gameData");
</script>

<h1 class="h2 mt-3">Details zu Spiel {match.match_id}</h1>

<section class="my-3 lg:my-5">
  <MatchMainInfoSection {match}/>
</section>

<section class="mb-5">

  <Tabs.Root
          bind:value={tabSet}
          class=""
  >
    <Tabs.List
            class="tabs-list preset-tonal-surface"
    >
      <Tabs.Trigger
              value="gameData"
              class="tabs-trigger"
      >Game Data</Tabs.Trigger>
      <Tabs.Trigger
              value="boxscore"
              class="tabs-trigger"
      >Box Score</Tabs.Trigger>
      <Tabs.Trigger
              value="gameReport"
              class="tabs-trigger"
      >Game Report</Tabs.Trigger>
    </Tabs.List>

    <Tabs.Content value="gameData" class="pt-3">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 place-items-stretch gap-3 md:gap-4 lg:gap-5">
        <MatchDetailStatsCard {match}/>
        <MatchDetailLocationCard field={match?.field}/>
        <MatchDetailOfficialsCard {match}/>
      </div>
    </Tabs.Content>

    <Tabs.Content value="boxscore">
      {#await data.singleGameStats}
        <ProgressRing/>
      {:then boxscore}
        {#if boxscore}
          <MatchBoxscoreSection {boxscore}/>
        {:else }
          <ContentFilteredUnavailable text="No Box Score available."/>
        {/if}
      {:catch error}
        <p>error loading boxscore: {error.message}</p>
      {/await}
    </Tabs.Content>

    <Tabs.Content value="gameReport">
      {#await data.gameReport}
        <ProgressRing/>
      {:then gameReport}
        {#if gameReport}
          <GameReport classes="my-2! md:max-w-[80%] 2xl:max-w-[70%]" report={gameReport}/>
        {:else }
          <ContentFilteredUnavailable text="No Game Report available."/>
        {/if}
      {:catch error}
        <p>error loading Game Report: {error.message}</p>
      {/await}
    </Tabs.Content>
  </Tabs.Root>
</section>