<script lang="ts">
    import SeasonSelector from "$lib/components/utility/SeasonSelector.svelte";
    import {ProgressBar} from "@skeletonlabs/skeleton";
    import ReloadUponSeasonChange from "$lib/components/navigation/ReloadUponSeasonChange.svelte";
    import {preferences} from "$lib/stores";

    export let data
</script>

<div class="my-2 md:flex justify-between items-center">
    <h1 class="h1 mb-3">Teams</h1>
    <div>
        <SeasonSelector/>
    </div>
</div>

<ReloadUponSeasonChange/>

{#await data.clubTeams}
    <p>Lade Teams...</p>
    <ProgressBar/>

{:then clubTeams}

    <ul class="list mt-5 flex flex-col gap-3">
        {#each clubTeams ?? [] as clubTeam}
            <li class="variant-soft-surface p-3 min-h-14">
                <a href="/teams/{clubTeam.id}?season={$preferences.selectedSeason}">
                    <span class="badge variant-ghost-primary w-20">{clubTeam.team.league_entries[0].league.acronym}</span>
                    <span class="flex-auto ms-3">{clubTeam.team.name}</span>
                </a>
            </li>
        {/each}
    </ul>

{:catch error}
    <p>error loading: {error.message}</p>
{/await}