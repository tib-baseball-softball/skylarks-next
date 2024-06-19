<script lang="ts">
    import type {PageServerData} from "../../../../.svelte-kit/types/src/routes/teams/[id]/$types";
    import type {ClubTeam} from "bsm.js/dist/model/ClubTeam";
    import {ProgressRadial} from "@skeletonlabs/skeleton";
    import BaseballStatsDatatable from "$lib/components/datatable/BaseballStatsDatatable.svelte";

    export let data: PageServerData
    $: clubTeam = data.clubTeam as ClubTeam
    $: battingStats = data.battingStats
</script>

<h1 class="h1 my-4">{clubTeam.team.name} (Saison {clubTeam.team.season})</h1>

<p>weitere Infos zum Team</p>

<h2 class="h2 my-4">Statistiken</h2>

<section class="my-2">
    {#await battingStats}
        <ProgressRadial/>
    {:then batting}
        {#if batting}
            <BaseballStatsDatatable data="{batting.data}" tableType="personal"/>
        {/if}
    {:catch error}
        <p>error loading: {error.message}</p>
    {/await}

</section>