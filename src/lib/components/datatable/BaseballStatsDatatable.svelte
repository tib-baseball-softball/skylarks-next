<script lang="ts">
    import ThSort from '$lib/components/datatable/ThSort.svelte';
    import ThFilter from '$lib/components/datatable/ThFilter.svelte';
    import Search from '$lib/components/datatable/Search.svelte';
    import RowsPerPage from '$lib/components/datatable/RowsPerPage.svelte';
    import RowCount from '$lib/components/datatable/RowCount.svelte';
    import Pagination from '$lib/components/datatable/Pagination.svelte';
    import { DataHandler } from '@vincjo/datatables';
    import {StatsType} from "bsm.js";

    //TODO: handle more than just team stats
    export let data: any[]
    export let rowsPerPage: number = 20
    export let type: StatsType

    const handler = new DataHandler(data, { rowsPerPage: rowsPerPage });
    const rows = handler.getRows();
</script>

<div class=" overflow-x-auto space-y-4 table-container">
    <!-- Header -->
    <header class="flex justify-between gap-4">
        <Search {handler} />
        <RowsPerPage {handler} />
    </header>
    <!-- Table -->
    <table class="table table-hover table-compact w-full table-auto">
        <thead>
            <tr>
                <ThSort {handler} orderBy="games">G</ThSort>
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
            <tr>
                <ThFilter {handler} filterBy="games" />
                <ThFilter {handler} filterBy="games_started" />
                <ThFilter {handler} filterBy="plate_appearances" />
                <ThFilter {handler} filterBy="at_bats" />
                <ThFilter {handler} filterBy="runs" />
                <ThFilter {handler} filterBy="runs_batted_in" />
                <ThFilter {handler} filterBy="hits" />
                <ThFilter {handler} filterBy="doubles" />
                <ThFilter {handler} filterBy="triples" />
                <ThFilter {handler} filterBy="homeruns" />
                <ThFilter {handler} filterBy="strikeouts" />
                <ThFilter {handler} filterBy="base_on_balls" />
                <ThFilter {handler} filterBy="intentional_base_on_balls" />
                <ThFilter {handler} filterBy="hit_by_pitches" />
                <ThFilter {handler} filterBy="stolen_bases" />
                <ThFilter {handler} filterBy="caught_stealings" />
                <ThFilter {handler} filterBy="sacrifice_hits" />
                <ThFilter {handler} filterBy="sacrifice_flys" />
                <ThFilter {handler} filterBy="interferences" />
                <ThFilter {handler} filterBy="batting_average" />
                <ThFilter {handler} filterBy="on_base_percentage" />
                <ThFilter {handler} filterBy="slugging_percentage" />
                <ThFilter {handler} filterBy="on_base_plus_slugging" />

            </tr>
        </thead>
        <tbody>
            {#each $rows as row}
                <tr>
                    <td>{row.games}</td>
                    <td>{row.games_started}</td>
                    <td>{row.plate_appearances}</td>
                    <td>{row.at_bats}</td>
                    <td>{row.runs}</td>
                    <td>{row.runs_batted_in}</td>
                    <td>{row.hits}</td>
                    <td>{row.doubles}</td>
                    <td>{row.triples}</td>
                    <td>{row.homeruns}</td>
                    <td>{row.strikeouts}</td>
                    <td>{row.base_on_balls}</td>
                    <td>{row.intentional_base_on_balls}</td>
                    <td>{row.hit_by_pitches}</td>
                    <td>{row.stolen_bases}</td>
                    <td>{row.caught_stealings}</td>
                    <td>{row.sacrifice_hits}</td>
                    <td>{row.sacrifice_flys}</td>
                    <td>{row.interferences}</td>
                    <td>{row.batting_average}</td>
                    <td>{row.on_base_percentage}</td>
                    <td>{row.slugging_percentage}</td>
                    <td>{row.on_base_plus_slugging}</td>
                </tr>
            {/each}
        </tbody>
    </table>
    <!-- Footer -->
    <footer class="flex justify-between">
        <RowCount {handler} />
        <Pagination {handler} />
    </footer>
</div>