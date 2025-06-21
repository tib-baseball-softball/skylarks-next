<script lang="ts">
  import type {HomeDataset} from "$lib/types/HomeDataset.ts";
  // @ts-ignore
  import {Tabs} from "bits-ui";
  import LeagueChartGraph from "$lib/components/favorite/LeagueChartGraph.svelte";

  interface Props {
    dataset: HomeDataset;
  }

  let {dataset}: Props = $props();

  let tabSet: "graph" | "percentage" | "series" | string = $state("percentage");
</script>

<Tabs.Root
        bind:value={tabSet}
        class=""
>
  <Tabs.List
          class="tabs-list border mb-1 preset-tonal-surface justify-around!"
  >
    <Tabs.Trigger
            value="graph"
            class="tabs-trigger"
    >
      Graph</Tabs.Trigger>
    <Tabs.Trigger
            value="percentage"
            class="tabs-trigger"
    >Percentage</Tabs.Trigger>
    <Tabs.Trigger
            value="series"
            class="tabs-trigger"
    >Series</Tabs.Trigger>
  </Tabs.List>

  <Tabs.Content value="graph" class="pt-4">
    <LeagueChartGraph table={dataset.table}/>
  </Tabs.Content>

  <Tabs.Content value="percentage" class="pt-4">

  </Tabs.Content>

  <Tabs.Content value="series" class="pt-4">

  </Tabs.Content>
</Tabs.Root>