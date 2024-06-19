<script lang="ts">
    import Search from '$lib/components/datatable/Search.svelte';
    import RowsPerPage from '$lib/components/datatable/RowsPerPage.svelte';
    import RowCount from '$lib/components/datatable/RowCount.svelte';
    import Pagination from '$lib/components/datatable/Pagination.svelte';
    import {DataHandler} from '@vincjo/datatables';
    import {StatsType} from "bsm.js";
    import StatsTableContent from "$lib/components/datatable/StatsTableContent.svelte";
    import type {StatisticsData} from "bsm.js/dist/model/AbstractStatisticsEntry";

    export let data: StatisticsData<"BattingStatistics" | "PitchingStatistics" | "FieldingStatistics">[]
    export let rowsPerPage: number = 25
    export let tableType: "personal" | "seasonal"

    let type: StatsType = StatsType.batting

    const handler = new DataHandler(data, { rowsPerPage: rowsPerPage });
</script>

<div class=" overflow-x-auto space-y-4 table-container">
    <!-- Header -->
    <header class="flex justify-between gap-4">
        <Search {handler} />
        <RowsPerPage {handler} />
    </header>
    <!-- Table -->
    <table class="table table-hover table-compact w-full table-auto">

        <StatsTableContent {handler} tableType="{tableType}" {type}/>

    </table>
    <!-- Footer -->
    <footer class="flex justify-between">
        <RowCount {handler} />
        <Pagination {handler} />
    </footer>
</div>