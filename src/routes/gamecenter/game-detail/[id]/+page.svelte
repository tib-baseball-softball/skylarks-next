<script lang="ts">
  import {ProgressRadial, Tab, TabGroup} from "@skeletonlabs/skeleton";
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
  let tabSet: number = $state(0);
</script>

<h1 class="h2 mt-3">Details zu Spiel {match.match_id}</h1>

<section class="my-3 lg:my-5">
  <MatchMainInfoSection {match}/>
</section>

<section class="mb-5">
  <TabGroup justify="justify-center">
    <Tab bind:group={tabSet} name="tabData" value={0}>Game Data</Tab>
    <Tab bind:group={tabSet} name="tabBoxscore" value={1}>Boxscore</Tab>
    <Tab bind:group={tabSet} name="tabGameReport" value={2}>Game Report</Tab>

    <!-- Tab Panels --->
    <svelte:fragment slot="panel">
      {#if tabSet === 0}

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 place-items-stretch gap-3 md:gap-4 lg:gap-5">
          <MatchDetailStatsCard {match}/>

          <MatchDetailLocationCard field={match?.field}/>

          <MatchDetailOfficialsCard {match}/>
        </div>

      {:else if tabSet === 1}

        {#await data.singleGameStats}
          <ProgressRadial/>
        {:then boxscore}
          {#if boxscore}
            <MatchBoxscoreSection {boxscore}/>
          {:else }
            <p>Kein Boxscore vorhanden.</p>
          {/if}
        {:catch error}
          <p>error loading boxscore: {error.message}</p>
        {/await}

      {:else if tabSet === 2}

        {#await data.gameReport}
          <ProgressRadial/>
        {:then gameReport}
          {#if gameReport}
            <GameReport classes="!my-2 md:max-w-[80%] 2xl:max-w-[70%]" report={gameReport}/>
          {:else }
            <ContentFilteredUnavailable text="No Game Report available."/>
          {/if}
        {:catch error}
          <p>error loading Game Report: {error.message}</p>
        {/await}


      {/if}
    </svelte:fragment>
  </TabGroup>
</section>

