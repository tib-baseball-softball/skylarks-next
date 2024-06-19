<script lang="ts">
    import type {DataHandler} from "@vincjo/datatables";
    import {StatsType} from "bsm.js";
    import BattingStatsHeaderRow from "$lib/components/datatable/BattingStatsHeaderRow.svelte";
    import BattingStatsContentRow from "$lib/components/datatable/BattingStatsContentRow.svelte";

    export let handler: DataHandler<any>
    export let tableType: "personal" | "seasonal"
    export let type: StatsType

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

    {#if type === StatsType.batting}
        <BattingStatsHeaderRow {handler}/>
    {/if}
</tr>
</thead>
<tbody>
{#each $rows as row}
    <tr>
        {#if tableType === "personal"}
            <td><a class="anchor" href="/players/{row.person.id}">{row.person.first_name} {row.person.last_name}</a></td>
        {/if}

        {#if tableType === "seasonal"}
            <td>{row.league.season} ({row.league.acronym})</td>
        {/if}

        {#if type === StatsType.batting}
            <BattingStatsContentRow {row}/>
        {/if}
    </tr>
{/each}
</tbody>