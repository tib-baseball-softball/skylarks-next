<script lang="ts">
  import type {PageProps} from "./$types";
  import type {StatsDataset} from "$lib/types/StatsDataset.ts";
  import BaseballStatsDatatable from "$lib/components/datatable/BaseballStatsDatatable.svelte";
  import {Progress} from "@skeletonlabs/skeleton-svelte";

  let {data}: PageProps = $props()

  async function getData(): Promise<StatsDataset> {
    return {
      batting: await data.battingStats!,
      pitching: await data.pitchingStats!,
      fielding: await data.fieldingStats!,
    };
  }
</script>

<h1 class="h1 mb-6!">Stats for {data.teamName ?? "Team"}</h1>
<section class="my-2 mb-4!">
  {#await getData()}
    <Progress/>
  {:then stats}
    {#if stats.batting && stats.pitching && stats.fielding}
      <BaseballStatsDatatable data={stats} tableType="personal"/>
    {/if}
  {:catch error}
    <p>error loading: {error.message}</p>
  {/await}
</section>