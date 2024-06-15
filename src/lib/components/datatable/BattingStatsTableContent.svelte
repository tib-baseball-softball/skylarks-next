<script lang="ts">
    import ThSort from "$lib/components/datatable/ThSort.svelte";
    import type {DataHandler} from "@vincjo/datatables";
    import type {StatsType} from "bsm.js";

    export let handler: DataHandler<any>[]
    export let type: "personal" | "seasonal"
    const rows = handler.getRows();
</script>

<thead>
<tr class="sticky">
    {#if type === "personal"}
        <th>Name</th>
    {/if}
    <ThSort {handler} orderBy="values.games">G</ThSort>
    <ThSort {handler} orderBy="games_started">GS</ThSort>
    <ThSort {handler} orderBy="plate_appearances">PA</ThSort>
    <ThSort {handler} orderBy="at_bats">AB</ThSort>
    <ThSort {handler} orderBy="runs">R</ThSort>
    <ThSort {handler} orderBy="runs_batted_in">RBI</ThSort>
    <ThSort {handler} orderBy="hits">H</ThSort>
    <ThSort {handler} orderBy="doubles">2B</ThSort>
    <ThSort {handler} orderBy="triples">3B</ThSort>
    <ThSort {handler} orderBy="homeruns">HR</ThSort>
    <ThSort {handler} orderBy="strikeouts">K</ThSort>
    <ThSort {handler} orderBy="base_on_balls">BB</ThSort>
    <ThSort {handler} orderBy="intentional_base_on_balls">IBB</ThSort>
    <ThSort {handler} orderBy="hit_by_pitches">HP</ThSort>
    <ThSort {handler} orderBy="stolen_bases">SB</ThSort>
    <ThSort {handler} orderBy="caught_stealings">CS</ThSort>
    <ThSort {handler} orderBy="sacrifice_hits">SH</ThSort>
    <ThSort {handler} orderBy="sacrifice_flys">SF</ThSort>
    <ThSort {handler} orderBy="interferences">I</ThSort>
    <ThSort {handler} orderBy="batting_average">AVG</ThSort>
    <ThSort {handler} orderBy="on_base_percentage">OBP</ThSort>
    <ThSort {handler} orderBy="slugging_percentage">SLG</ThSort>
    <ThSort {handler} orderBy="on_base_plus_slugging">OPS</ThSort>
</tr>
</thead>
<tbody>
{#each $rows as row}
    <tr>
        {#if type === "personal"}
            <td>{row.person.first_name} {row.person.last_name}</td>
        {/if}
        <td>{row.values.games}</td>
        <td>{row.values.games_started}</td>
        <td>{row.values.plate_appearances}</td>
        <td>{row.values.at_bats}</td>
        <td>{row.values.runs}</td>
        <td>{row.values.runs_batted_in}</td>
        <td>{row.values.hits}</td>
        <td>{row.values.doubles}</td>
        <td>{row.values.triples}</td>
        <td>{row.values.homeruns}</td>
        <td>{row.values.strikeouts}</td>
        <td>{row.values.base_on_balls}</td>
        <td>{row.values.intentional_base_on_balls}</td>
        <td>{row.values.hit_by_pitches}</td>
        <td>{row.values.stolen_bases}</td>
        <td>{row.values.caught_stealings}</td>
        <td>{row.values.sacrifice_hits}</td>
        <td>{row.values.sacrifice_flys}</td>
        <td>{row.values.interferences}</td>
        <td>{row.values.batting_average}</td>
        <td>{row.values.on_base_percentage}</td>
        <td>{row.values.slugging_percentage}</td>
        <td>{row.values.on_base_plus_slugging}</td>
    </tr>
{/each}
</tbody>