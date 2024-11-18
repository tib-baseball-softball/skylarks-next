<script lang="ts">
  import {ConicGradient, type ConicStop, ProgressRadial} from "@skeletonlabs/skeleton";
  import type {SingleStatElement} from "$lib/types/SingleStatElement";
  import StatsBlockContent from "$lib/components/utility/StatsBlockContent.svelte";

  let {data} = $props()

  const conicStops: ConicStop[] = [
    {
      label: "In",
      color: "rgb(var(--color-success-500))",
      start: 0,
      end: 40,
    },
    {
      label: "Maybe",
      color: "rgb(var(--color-warning-500))",
      start: 40,
      end: 66,
    },
    {
      label: "Out",
      color: "rgb(var(--color-error-500))",
      start: 66,
      end: 100,
    },
  ];

  const block: SingleStatElement = {
    title: "Games",
    value: "10 / 16",
    desc: "All possible games for 2024 season",
  };

  const block2: SingleStatElement = {
    title: "Practices",
    value: "23 / 45",
    desc: "All possible practices for your teams (2024 season)",
  };
</script>

<h1 class="h1 lg:mt-4">My Stats</h1>

<div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    <div class="card variant-glass-surface shadow-lg">
        <header class="card-header">
            <h2 class="h4 font-semibold">Attendance Stats</h2>
        </header>

        <section class="p-4">
            <ConicGradient legend stops={conicStops}></ConicGradient>
        </section>

        <footer class="mt-1 card-footer font-light text-sm">
            Events where you did not select anything are counted as "out".
        </footer>
    </div>

    <div class="card variant-glass-surface shadow-lg">
        <header class="card-header">
            <h2 class="h4 font-semibold">Totals - Games</h2>
        </header>

        <section class="p-4 flex justify-center">
            <StatsBlockContent {block} classes="place-items-center gap-3"/>
        </section>
    </div>

    <div class="card variant-glass-surface shadow-lg">
        <header class="card-header">
            <h2 class="h4 font-semibold">Totals- Practices</h2>
        </header>

        <section class="p-4 flex justify-center">
            <StatsBlockContent
                    block={block2}
                    classes="place-items-center gap-3"
            />
        </section>
    </div>
</div>

{#await data.statsItem}
    <ProgressRadial/>
{:then statsItem}
    {#if statsItem}
        <p class="code">{JSON.stringify(statsItem)}</p>
    {/if}
{:catch error}
    <p>error loading: {error.message}</p>
{/await}