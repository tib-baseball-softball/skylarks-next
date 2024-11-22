<script lang="ts">
  import {ProgressRadial} from "@skeletonlabs/skeleton";
  import StatsByTypePieChart from "$lib/components/diamondplanner/stats/StatsByTypePieChart.svelte";
  import AttendanceTotalStatsBlock from "$lib/components/diamondplanner/stats/AttendanceTotalStatsBlock.svelte";
  import {goto} from "$app/navigation";
  import {range} from "$lib/functions/range";

  let {data} = $props()

  const currentYear = new Date().getFullYear()
  const seasonOptions = range(2023, currentYear)

  let selectedSeason = $state(currentYear)

  const reloadWithQuery = () => {
    let queryString = `?season=${selectedSeason}`;

    goto(queryString, {
      noScroll: true,
    });
  };

  $effect.pre(() => {
    console.log(selectedSeason);
    reloadWithQuery();
  });
</script>

<h1 class="h1 lg:mt-4">My Stats</h1>

<div
        class="flex flex-wrap gap-4 lg:gap-8 variant-soft-surface justify-between px-4 py-3 rounded-token"
>
    <label class="label flex items-center gap-2 flex-grow justify-between md:flex-grow-0">
        Season
        <select class="select" bind:value={selectedSeason}>
            {#each seasonOptions as option}
                <option value="{option}">{option}</option>
            {/each}
        </select>
    </label>
</div>

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
