<script lang="ts">
    import type {PageServerData} from "../../../../.svelte-kit/types/src/routes/ligen/[id]/$types";
    import {ProgressRadial} from "@skeletonlabs/skeleton";
    import StandingsTable from "$lib/components/table/StandingsTable.svelte";
    import LeagueDetailInfoCard from "$lib/components/league/LeagueDetailInfoCard.svelte";

    interface Props {
        data: PageServerData;
    }

    let { data }: Props = $props();
    let expectedTable = $derived(data.table)
    let leagueGroup = $derived(data.leagueGroup)
</script>

<h1 class="h1 my-4">{leagueGroup.name} ({leagueGroup.season})</h1>

<section class="my-5">
    <h2 class="h2">Informationen</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3">

        <LeagueDetailInfoCard {leagueGroup}/>

    </div>
</section>

<section class="my-3">

    <h2 class="h2">Tabelle</h2>
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

<style lang="postcss">
    h2 {
        @apply mb-3
    }
</style>