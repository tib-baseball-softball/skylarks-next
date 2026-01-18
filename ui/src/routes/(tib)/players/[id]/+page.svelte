<script lang="ts">
  import BaseballStatsDatatable from "$lib/tib/components/stats/BaseballStatsDatatable.svelte";
  import PlayerDataCard from "$lib/dp/components/player/PlayerDataCard.svelte";
  import PlayerHeaderSection from "$lib/dp/components/player/PlayerHeaderSection.svelte";
  import ProgressRing from "$lib/dp/components/utils/ProgressRing.svelte";
  import type {StatsDataset} from "$lib/tib/types/StatsDataset.ts";
  import type {PageProps} from "../../../../../.svelte-kit/types/src/routes";

  const {data}: PageProps = $props();

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
  <ProgressRing/>
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
  <ProgressRing/>
{:then stats}
  <BaseballStatsDatatable data={stats} tableType="seasonal"/>
{:catch error}
  <p>error loading: {error.message}</p>
{/await}
