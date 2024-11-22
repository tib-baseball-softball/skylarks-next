<script lang="ts">
  import {ProgressRadial} from "@skeletonlabs/skeleton";
  import StatsByTypePieChart from "$lib/components/diamondplanner/stats/StatsByTypePieChart.svelte";
  import AttendanceTotalStatsBlock from "$lib/components/diamondplanner/stats/AttendanceTotalStatsBlock.svelte";

  let {data} = $props()
</script>

<h1 class="h1 lg:mt-4">My Stats</h1>

{#await data.statsItem}
    <ProgressRadial/>
{:then statsItem}
    {#if statsItem}
        <section class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
            <StatsByTypePieChart statsItem={statsItem}/>
        </section>

        <section class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
            {#each statsItem.attendanceTotals as attendance}
                <AttendanceTotalStatsBlock
                        season={statsItem.season}
                        attendance={attendance}
                />
            {/each}
        </section>
    {/if}
{:catch error}
    <p>error loading: {error.message}</p>
{/await}
