<script lang="ts">
    import Search from '$lib/components/datatable/Search.svelte';
    import RowsPerPage from '$lib/components/datatable/RowsPerPage.svelte';
    import RowCount from '$lib/components/datatable/RowCount.svelte';
    import Pagination from '$lib/components/datatable/Pagination.svelte';
    import {DataHandler} from '@vincjo/datatables';
    import {StatsType} from "bsm.js";
    import StatsTableContent from "$lib/components/datatable/StatsTableContent.svelte";
    import {RadioGroup, RadioItem} from "@skeletonlabs/skeleton";
    import type {StatisticsData} from "bsm.js/dist/model/AbstractStatisticsEntry";
    import type {StatsDataset} from "$lib/types/StatsDataset";

    export let data: StatsDataset
    export let rowsPerPage: number = 25
    export let tableType: "personal" | "seasonal"

    let type: StatsType = StatsType.batting

    const handler = new DataHandler<StatisticsData<"BattingStatistics" | "PitchingStatistics" | "FieldingStatistics">>(data.batting.data, { rowsPerPage: rowsPerPage })

    function changeDataForHandler(type: StatsType) {
        switch (type) {
            case StatsType.batting:
                handler.setRows(data.batting.data)
                break
            case StatsType.pitching:
                handler.setRows(data.pitching.data)
                break
            case StatsType.fielding:
                handler.setRows(data.fielding.data)
                break
        }
    }

    $: changeDataForHandler(type)
</script>

<div class=" overflow-x-auto space-y-4 table-container">
    <!-- Header -->
    <header class="flex justify-between gap-4">
        <Search {handler} />

        <RadioGroup>
            <RadioItem bind:group={type} name="batting" value={StatsType.batting}>Batting</RadioItem>
            <RadioItem bind:group={type} name="pitching" value={StatsType.pitching}>Pitching</RadioItem>
            <RadioItem bind:group={type} name="fielding" value={StatsType.fielding}>Fielding</RadioItem>
        </RadioGroup>

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