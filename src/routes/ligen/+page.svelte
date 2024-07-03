<script lang="ts">
    import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
    import {ProgressBar} from "@skeletonlabs/skeleton";
    import type {PageServerData} from "../../../.svelte-kit/types/src/routes/ligen/$types";
    import ReloadUponSeasonChange from "$lib/components/navigation/ReloadUponSeasonChange.svelte";

    export let data: PageServerData
</script>

<ReloadUponSeasonChange/>

<h1 class="h1 mb-3">Ligen</h1>
<SeasonSelector/>
{#await data.leagueGroups}
    <p>Lade Ligen...</p>
    <ProgressBar/>

{:then leagueGroups}

    <ul class="list mt-5 flex flex-col gap-3">
        {#each leagueGroups as leagueGroup}
            <li class="variant-soft-surface p-3 min-h-14">
                <a href="/ligen/{leagueGroup.id}">
                    <span class="badge variant-ghost-primary w-20">{leagueGroup.acronym}</span>
                    <span class="flex-auto ms-3">{leagueGroup.name}</span>
                </a>
            </li>
        {/each}
    </ul>

{:catch error}
    <p>error loading: {error.message}</p>
{/await}