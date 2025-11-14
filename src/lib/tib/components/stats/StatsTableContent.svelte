<script lang="ts">
  import type {DataHandler} from "@vincjo/datatables";
  import type {StatisticsData} from "bsm.js";
  import {StatsType} from "bsm.js";
  import StatsHeaderRow from "$lib/tib/components/stats/StatsHeaderRow.svelte";
  import StatsContentRow from "$lib/tib/components/stats/StatsContentRow.svelte";

  interface Props {
    handler: DataHandler<
        StatisticsData<"BattingStatistics" | "PitchingStatistics" | "FieldingStatistics">
    >;
    tableType: "personal" | "seasonal";
    type: StatsType;
  }

  let {handler, tableType, type}: Props = $props();

  const rows = handler.getRows();
</script>

<thead>
<tr class="sticky">
  {#if tableType === "personal"}
    <th>Name</th>
  {/if}

  {#if tableType === "seasonal"}
    <th>Saison</th>
  {/if}

  <StatsHeaderRow {handler} {tableType} {type}/>
</tr>
</thead>
<tbody>
{#each $rows as row}
  <tr>
    {#if tableType === "personal"}
      <td><a class="anchor"
             href="/(tib)/players/{row?.person?.id}">{row?.person?.first_name} {row?.person?.last_name}</a>
      </td>
    {/if}

    {#if tableType === "seasonal"}
      <td>{row?.league?.season} ({row?.league?.acronym})</td>
    {/if}

    <StatsContentRow {row} {type} {tableType}/>
  </tr>
{/each}

{#if $rows.length === 0}
  <tr>
    <th class="py-4" colspan="8">No data available for this category.</th>
  </tr>
{/if}

</tbody>