<script lang="ts">
    import MatchTeaserCard from "$lib/components/match/MatchTeaserCard.svelte";
    import type {Match} from "bsm.js";
    import {PUBLIC_TEAM_NAME} from "$env/static/public";
    import ContentFilteredUnavailable from "$lib/components/match/ContentFilteredUnavailable.svelte";

    export let matches: Match[] = []
    export let showExternal = false
    let skylarksGames = matches.filter(match => match.away_team_name.includes(PUBLIC_TEAM_NAME) || match.home_team_name.includes(PUBLIC_TEAM_NAME))
</script>

<section>
    <header>
        <h2 class="h2 my-3">Spiele</h2>
    </header>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 lg:gap-6 mb-6">
        {#if showExternal}

            {#each matches as match}
                <MatchTeaserCard {match}/>
            {/each}

            {#if matches.length === 0}
                <ContentFilteredUnavailable/>
            {/if}

        {:else}

            {#each skylarksGames as match}
                <MatchTeaserCard {match}/>
            {/each}

            {#if skylarksGames.length === 0}
                <ContentFilteredUnavailable/>
            {/if}

        {/if}
    </div>
</section>