<script lang="ts">
  import {ProgressRing} from "@skeletonlabs/skeleton-svelte";
  import StatsByTypePieChart from "$lib/components/diamondplanner/stats/StatsByTypePieChart.svelte";
  import AttendanceTotalStatsBlock from "$lib/components/diamondplanner/stats/AttendanceTotalStatsBlock.svelte";
  import {goto} from "$app/navigation";
  import {range} from "$lib/functions/range";
  import type {PersonalAttendanceStatsItem} from "$lib/model/PersonalAttendanceStats";

  let {data} = $props();

  const user = $derived(data.user);

  const currentYear = new Date().getFullYear();
  const seasonOptions = range(2023, currentYear);

  let selectedSeason = $state(currentYear);

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

<h1 class="h1 mt-4! mb-6!">Stats for {user?.first_name} {user?.last_name}</h1>

<div
        class="flex flex-wrap gap-4 lg:gap-8 preset-tonal-surface justify-between px-4 py-3 rounded-base"
>
  <label class="label flex items-center gap-2 grow justify-between md:grow-0">
    Season
    <select class="select" bind:value={selectedSeason}>
      {#each seasonOptions as option}
        <option value="{option}">{option}</option>
      {/each}
    </select>
  </label>
</div>

{#snippet statsSection(statsItem: PersonalAttendanceStatsItem)}
  <h2 class="h2 mb-2">{statsItem.teamName}</h2>
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
{/snippet}

{#await data.statsItem}
  <ProgressRing/>
{:then statsItem}

  {#if statsItem}
    {@render statsSection(statsItem)}
  {/if}

{:catch error}
  <p>error loading: {error.message}</p>
{/await}

{#await Promise.all(data.teamStatsItems)}
  <ProgressRing/>
{:then teamStatsItems}

  {#each teamStatsItems as teamStatsItem}
    <div class="mt-8! xl:mt-9!">
      {@render statsSection(teamStatsItem)}
    </div>
  {/each}

{:catch error}
  <p>error loading: {error.message}</p>
{/await}
