<script lang="ts">
  // @ts-ignore
  import {Tabs} from "bits-ui";
  import MatchMainInfoSection from "$lib/tib/components/match/MatchMainInfoSection.svelte";
  import MatchBoxscoreSection from "$lib/tib/components/boxscore/MatchBoxscoreSection.svelte";
  import MatchDetailStatsCard from "$lib/tib/components/match/MatchDetailStatsCard.svelte";
  import MatchDetailLocationCard from "$lib/tib/components/match/MatchDetailLocationCard.svelte";
  import MatchDetailOfficialsCard from "$lib/tib/components/match/MatchDetailOfficialsCard.svelte";
  import type {PageProps} from "./$types";
  import ContentFilteredUnavailable from "$lib/tib/components/match/ContentFilteredUnavailable.svelte";
  import GameReport from "$lib/tib/components/gameReport/GameReport.svelte";
  import ShareButton from "$lib/tib/components/match/ShareButton.svelte";
  import ProgressRing from "$lib/dp/components/utils/ProgressRing.svelte";

  let {data}: PageProps = $props();

  let match = $derived(data.match);
  let tabSet: "gameData" | "boxscore" | "gameReport" | string = $state("gameData");
</script>

<h1 class="h2 mt-3">Details zu Spiel {match.match_id}</h1>

<section class="my-3 lg:my-5">
  <MatchMainInfoSection {match}/>
</section>

<section class="mb-8">

  <Tabs.Root
    bind:value={tabSet}
  >
    <Tabs.List
      class="tabs-list preset-tonal-surface mt-6"
    >
      <Tabs.Trigger
        class="tabs-trigger btn"
        value="gameData"
      >Game Data
      </Tabs.Trigger>
      <Tabs.Trigger
        class="tabs-trigger btn"
        value="boxscore"
      >Box Score
      </Tabs.Trigger>
      <Tabs.Trigger
        class="tabs-trigger btn"
        value="gameReport"
      >Game Report
      </Tabs.Trigger>
    </Tabs.List>

    <Tabs.Content class="pt-4" value="gameData">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 place-items-stretch gap-5">
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
          <ContentFilteredUnavailable text="Box Score for this game could not be loaded or is unavailable."/>
        {/if}
      {:catch error}
        <p>error loading box score: {error.message}</p>
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

<section class="my-5!">
  <ShareButton {match}/>
  <details class="mt-4">
    <summary class="font-light">Caveat</summary>
    <p class="text-sm font-light">Sharing games does not work on Firefox.</p>
  </details>
</section>
