<script lang="ts">
  import { Progress } from "@skeletonlabs/skeleton-svelte";
  import BaseballStatsDatatable from "$lib/components/datatable/BaseballStatsDatatable.svelte";
  import type {StatsDataset} from "$lib/types/StatsDataset";
  import PlayerHeaderSection from "$lib/components/player/PlayerHeaderSection.svelte";
  import PlayerDataCard from "$lib/components/player/PlayerDataCard.svelte";
  import type {PageProps} from "./$types";

  let {data}: PageProps = $props();

  async function getData(): Promise<StatsDataset> {
    return {
      batting: await data.battingStats!,
      pitching: await data.pitchingStats!,
      fielding: await data.fieldingStats!,
    };
  }
</script>

{#await data.battingStats then batting}
  <h1 class="h1 my-4">
    Player Profile for {batting.person?.first_name}
    {batting.person?.last_name}
  </h1>
{/await}

{#await data.player}
  <Progress/>
{:then player}
  {#if player}
    <div class="my-8!">
      <h2 class="h2 my-4">Profile Data</h2>
      <section class="grid grid-cols-1 md:grid-cols-2 gap-3 md:gap-4">
        <PlayerHeaderSection {player}/>
        <PlayerDataCard {player}/>
      </section>
    </div>
  {/if}
{:catch error}
  <p>An error occurred fetching player data: {error.message}.<br>
    The data source might be temporarily unavailable, please try again later.</p>
{/await}

<h2 class="h2 my-6 lg:mt-10 mb-8">Stats</h2>

{#await getData()}
  <Progress/>
{:then stats}
  <BaseballStatsDatatable data={stats} tableType="seasonal"/>
{:catch error}
  <p>error loading: {error.message}</p>
{/await}
