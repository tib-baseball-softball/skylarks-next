<script lang="ts">
    import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
    import {ProgressBar} from "@skeletonlabs/skeleton";
    import type {PageServerData} from "../../../.svelte-kit/types/src/routes/ligen/$types";

    export let data: PageServerData
</script>

<SeasonSelector/>
{#await data.leagueGroups}
    <p>Lade Ligen...</p>
    <ProgressBar />

{:then leagueGroups}

    {#each leagueGroups as leagueGroup}
        <p>{leagueGroup.name}</p>
    {/each}

{:catch error}
    <p>error loading: {error.message}</p>
{/await}