<script lang="ts">
    import type {PageServerData} from "../../../../.svelte-kit/types/src/routes/teams/[id]/$types";
    import type {ClubTeam} from "bsm.js/dist/model/ClubTeam";
    import {ProgressRadial} from "@skeletonlabs/skeleton";
    import BaseballStatsDatatable from "$lib/components/datatable/BaseballStatsDatatable.svelte";
    import type {StatsDataset} from "$lib/types/StatsDataset";

    export let data: PageServerData
    $: clubTeam = data.clubTeam as ClubTeam

    async function getData(): Promise<StatsDataset> {
        return {
            batting: await data.battingStats!,
            pitching: await data.pitchingStats!,
            fielding: await data.fieldingStats!
        }
    }
</script>

<h1 class="h1 my-4">{clubTeam.team.name} (Saison {clubTeam.team.season})</h1>

<p>weitere Infos zum Team</p>

<h2 class="h2 my-4">Statistiken</h2>

<section class="my-2">
    {#await getData()}
        <ProgressRadial/>
    {:then stats}
        {#if stats.batting && stats.pitching && stats.fielding}
            <BaseballStatsDatatable data="{stats}" tableType="personal"/>
        {/if}
    {:catch error}
        <p>error loading: {error.message}</p>
    {/await}

</section>