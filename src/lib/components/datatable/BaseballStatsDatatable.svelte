<script lang="ts">
    import Search from '$lib/components/datatable/Search.svelte';
    import RowsPerPage from '$lib/components/datatable/RowsPerPage.svelte';
    import RowCount from '$lib/components/datatable/RowCount.svelte';
    import Pagination from '$lib/components/datatable/Pagination.svelte';
    import {DataHandler} from '@vincjo/datatables';
    import {StatsType} from "bsm.js";
    import StatsTableContent from "$lib/components/datatable/StatsTableContent.svelte";
    import {RadioGroup, RadioItem} from "@skeletonlabs/skeleton";
    import type {StatisticsData, StatisticsSummary} from "bsm.js/dist/model/AbstractStatisticsEntry";
    import type {StatsDataset} from "$lib/types/StatsDataset";
    import StatsContentRow from "$lib/components/datatable/StatsContentRow.svelte";

    export let data: StatsDataset
    export let rowsPerPage: number = 25
    export let tableType: "personal" | "seasonal"

    let type: StatsType = StatsType.batting
    let summaryData: StatisticsSummary<"BattingStatistics" | "PitchingStatistics" | "FieldingStatistics"> | undefined = data.batting.summaries.at(0)

    const handler = new DataHandler<StatisticsData<"BattingStatistics" | "PitchingStatistics" | "FieldingStatistics">>(data.batting.data, {rowsPerPage: rowsPerPage})

    function changeDataForHandler(type: StatsType) {
        switch (type) {
            case StatsType.batting:
                handler.setRows(data.batting.data)
                summaryData = data.batting.summaries.at(0)
                break
            case StatsType.pitching:
                summaryData = data.pitching.summaries.at(0)
                handler.setRows(data.pitching.data)
                break
            case StatsType.fielding:
                summaryData = data.fielding.summaries.at(0)
                handler.setRows(data.fielding.data)
                break
        }
    }

    $: changeDataForHandler(type)
</script>

<div class=" overflow-x-auto space-y-4 table-container">
    <!-- Header -->
    <header class="flex justify-between gap-4">
        <Search {handler}/>

        <RadioGroup>
            <RadioItem bind:group={type} name="batting" value={StatsType.batting}>Batting</RadioItem>
            <RadioItem bind:group={type} name="pitching" value={StatsType.pitching}>Pitching</RadioItem>
            <RadioItem bind:group={type} name="fielding" value={StatsType.fielding}>Fielding</RadioItem>
        </RadioGroup>

        <RowsPerPage {handler}/>
    </header>
    <!-- Table -->
    <table class="table table-hover table-compact w-full table-auto">

        <StatsTableContent {handler} {tableType} {type}/>

        <tfoot>
        {#if summaryData}
            <tr>
                <th>Gesamt</th>

                <StatsContentRow {type} {tableType} row="{summaryData}"/>
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