<script lang="ts">
    import type {ClubTeam} from "bsm.js/dist/model/ClubTeam";
    import {ProgressRadial} from "@skeletonlabs/skeleton";
    import BaseballStatsDatatable from "$lib/components/datatable/BaseballStatsDatatable.svelte";
    import type {StatsDataset} from "$lib/types/StatsDataset";
    import TeamDetailInfoCard from "$lib/components/team/TeamDetailInfoCard.svelte";

    export let data
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

<section class="my-5">
    <h2 class="h2">Informationen</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3">

        <TeamDetailInfoCard {clubTeam}/>

    </div>
</section>


<section class="my-2">
    <h2 class="h2">Statistiken</h2>
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

<style lang="postcss">
    h2 {
        @apply mb-3
    }
</style>