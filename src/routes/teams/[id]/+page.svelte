<script lang="ts">
  import type {ClubTeam} from "bsm.js";
  import {ProgressRing} from "@skeletonlabs/skeleton-svelte";
  import BaseballStatsDatatable from "$lib/components/datatable/BaseballStatsDatatable.svelte";
  import type {StatsDataset} from "$lib/types/StatsDataset";
  import TeamDetailInfoCard from "$lib/components/team/TeamDetailInfoCard.svelte";
  import StandingsTable from "$lib/components/table/StandingsTable.svelte";

  interface Props {
    data: any;
  }

  let {data}: Props = $props();
  let clubTeam = $derived(data.clubTeam as ClubTeam);

  async function getData(): Promise<StatsDataset> {
    return {
      batting: await data.battingStats!,
      pitching: await data.pitchingStats!,
      fielding: await data.fieldingStats!,
    };
  }
</script>

<h1 class="h1 my-4">{clubTeam.team.name} (Saison {clubTeam.team.season})</h1>

<section class="my-5 lg:max-w-[50%]">
  <h2 class="h2">Information</h2>

  <TeamDetailInfoCard {clubTeam}/>
</section>

<section>
  <h2 class="h2">Standings</h2>

  {#await data?.table}
    <ProgressRing/>
  {:then table}
    {#if table}
      <div class="col-span-2 standings-container">
        <StandingsTable {table}/>
      </div>
    {:else }
      <p class="col-span-2">No Standings available.</p>
    {/if}
  {:catch error}
    <p>error loading: {error.message}</p>
  {/await}
</section>

<section class="my-2 mb-4!">
  <h2 class="h2">Stats</h2>
  {#await getData()}
    <ProgressRing/>
  {:then stats}
    {#if stats.batting && stats.pitching && stats.fielding}
      <BaseballStatsDatatable data={stats} tableType="personal"/>
    {/if}
  {:catch error}
    <p>error loading: {error.message}</p>
  {/await}
</section>

<style lang="postcss">
    h2 {
        margin-bottom: calc(var(--spacing) * 3)
    }

    /* ugly hack to prevent table overflow */
    @media (min-width: 1400px) and (max-width: 1800px) {
        .standings-container {
            max-width: 90%;
        }
    }
</style>
