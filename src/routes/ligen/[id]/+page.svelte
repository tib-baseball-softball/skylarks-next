<script lang="ts">
    import type {PageServerData} from "../../../../.svelte-kit/types/src/routes/ligen/[id]/$types";
    import {ProgressRadial} from "@skeletonlabs/skeleton";
    import StandingsTable from "$lib/components/table/StandingsTable.svelte";

    export let data: PageServerData
    $: expectedTable = data.table
    $: leagueGroup = data.leagueGroup
</script>

<h1 class="h1 my-4">{leagueGroup.name} ({leagueGroup.season})</h1>

<section class="my-3">
    {#await expectedTable}
        <ProgressRadial/>
    {:then table}
        {#if table}
            <StandingsTable {table}/>
        {/if}
    {:catch error}
        <p>error loading: {error.message}</p>
    {/await}

</section>
