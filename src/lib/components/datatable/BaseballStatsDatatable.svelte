<script lang="ts">
  import Search from "$lib/components/datatable/Search.svelte";
  import RowsPerPage from "$lib/components/datatable/RowsPerPage.svelte";
  import RowCount from "$lib/components/datatable/RowCount.svelte";
  import Pagination from "$lib/components/datatable/Pagination.svelte";
  import {DataHandler} from "@vincjo/datatables";
  import type {StatisticsData, StatisticsSummary} from "bsm.js";
  import {StatsType} from "bsm.js";
  import StatsTableContent from "$lib/components/datatable/StatsTableContent.svelte";
  import {Segment} from "@skeletonlabs/skeleton-svelte";
  import type {StatsDataset} from "$lib/types/StatsDataset";
  import StatsContentRow from "$lib/components/datatable/StatsContentRow.svelte";
  import StatsBlock from "$lib/components/utility/StatsBlock.svelte";

  interface Props {
    data: StatsDataset;
    rowsPerPage?: number;
    tableType: "personal" | "seasonal";
  }

  let {data, rowsPerPage = 10, tableType}: Props = $props();

  let type: StatsType = $state(StatsType.batting);
  let summaryData:
      | StatisticsSummary<
      "BattingStatistics" | "PitchingStatistics" | "FieldingStatistics"
  >
      | undefined = $state(data.batting.summaries.at(0));

  const handler = new DataHandler<
      StatisticsData<
          "BattingStatistics" | "PitchingStatistics" | "FieldingStatistics"
      >
  >(data.batting.data, {rowsPerPage: rowsPerPage});

  function changeDataForHandler(type: StatsType) {
    switch (type) {
      case StatsType.batting:
        handler.setRows(data.batting.data);
        summaryData = data.batting.summaries.at(0);
        break;
      case StatsType.pitching:
        summaryData = data.pitching.summaries.at(0);
        handler.setRows(data.pitching.data);
        break;
      case StatsType.fielding:
        summaryData = data.fielding.summaries.at(0);
        handler.setRows(data.fielding.data);
        break;
    }
  }

  $effect.pre(() => {
    changeDataForHandler(type);
  });
</script>

<div class="overflow-x-auto space-y-4 table-wrap">
  <!-- Header -->
  <header class="flex justify-between gap-4">
    <Search {handler}/>

    <Segment name="stats-type" value={type} onValueChange={(e) => (type = e.value ?? StatsType.batting)}>
      <Segment.Item value={StatsType.batting} classes="flex-grow">Batting</Segment.Item>
      <Segment.Item value={StatsType.pitching} classes="flex-grow">Pitching</Segment.Item>
      <Segment.Item value={StatsType.fielding} classes="flex-grow">Fielding</Segment.Item>
    </Segment>

    <RowsPerPage {handler}/>
  </header>

  <div class="flex flex-col">
    <div
            class="stats stats-vertical sm:stats-horizontal preset-tonal-surface rounded-container"
    >
      <StatsBlock {type} row={summaryData}/>
    </div>
  </div>

  <!-- Table -->
  <table class="table  table-compact w-full">
    <StatsTableContent {handler} {tableType} {type}/>

    <tfoot>
    {#if summaryData}
      <tr>
        <th>Total</th>

        <StatsContentRow {type} {tableType} row={summaryData}/>
      </tr>
    {/if}
    </tfoot>
  </table>
  <!-- Footer -->
  <footer class="flex justify-between">
    <RowCount {handler}/>
    <Pagination {handler}/>
  </footer>
</div>

<style lang="postcss">
    .stats {
        margin-top: 1rem !important;
        margin-bottom: 0.75rem !important;
    }

    @media (min-width: 1024px) {
        .stats {
            margin-top: 1.5rem !important;
            margin-bottom: 1rem !important;
        }
    }

    /* ugly hack to prevent table overflow */
    @media (min-width: 1400px) and (max-width: 1800px) {
        .table-wrap {
            max-width: 90%;
        }
    }
</style>
