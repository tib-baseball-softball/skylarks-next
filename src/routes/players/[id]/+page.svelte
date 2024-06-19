<script lang="ts">
    import type {PageServerData} from "../../../../.svelte-kit/types/src/routes/players/[id]/$types";
    import {ProgressBar} from "@skeletonlabs/skeleton";
    import BaseballStatsDatatable from "$lib/components/datatable/BaseballStatsDatatable.svelte";
    import type {StatsDataset} from "$lib/types/StatsDataset";

    export let data: PageServerData

    async function getData(): Promise<StatsDataset> {
        return {
            batting: await data.battingStats!,
            pitching: await data.pitchingStats!,
            fielding: await data.fieldingStats!
        }
    }
</script>

<h1 class="h1 my-4">Spieler*innenprofil</h1>

<h2 class="h2 my-4">Statistiken</h2>

{#await getData()}
        <ProgressBar/>
    {:then stats}
        <BaseballStatsDatatable data="{stats}" tableType="seasonal"/>
    {:catch error}
    <p>error loading: {error.message}</p>
{/await}