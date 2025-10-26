<script lang="ts">
  import type {HomeDataset} from "$lib/types/HomeDataset.ts";
  // @ts-ignore
  import {Tabs} from "bits-ui";
  import LeagueChartGraph from "$lib/components/favorite/LeagueChartGraph.svelte";
  import WinningPercentageContainer from "$lib/components/favorite/WinningPercentageContainer.svelte";
  import StreakContainer from "$lib/components/favorite/StreakContainer.svelte";

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
            class="tabs-trigger btn"
            value="graph"
    >
      Graph
    </Tabs.Trigger>
    <Tabs.Trigger
            class="tabs-trigger btn"
            value="percentage"
    >Percentage
    </Tabs.Trigger>
    <Tabs.Trigger
            class="tabs-trigger btn"
            value="series"
    >Series
    </Tabs.Trigger>
  </Tabs.List>

  <Tabs.Content class="pt-4" value="graph">
    <LeagueChartGraph table={dataset.table}/>
  </Tabs.Content>

  <Tabs.Content class="pt-4" value="percentage">
    <WinningPercentageContainer tableRow={dataset.table_row}/>
  </Tabs.Content>

  <Tabs.Content class="pt-4" value="series">
    <StreakContainer dataset={dataset}/>
  </Tabs.Content>
</Tabs.Root>