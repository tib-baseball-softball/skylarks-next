<script lang="ts">
    import type {PageServerData} from "../../../../.svelte-kit/types/src/routes/players/[id]/$types";
    import {ProgressBar} from "@skeletonlabs/skeleton";
    import {StatsType} from "bsm.js";
    import BaseballStatsDatatable from "$lib/components/datatable/BaseballStatsDatatable.svelte";

    export let data: PageServerData
    $: playerBattingStats = data.playerBattingStats
</script>

<h1 class="h1 my-4">Spieler*innenprofil</h1>

<h2 class="h2 my-4">Statistiken</h2>

{#await playerBattingStats}
        <ProgressBar/>
    {:then batting}
        <BaseballStatsDatatable type="{StatsType.batting}" data="{batting.data}" tableType="seasonal"/>
    {:catch error}
    <p>error loading: {error.message}</p>
{/await}