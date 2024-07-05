<script lang="ts">
    import type {Match} from "bsm.js";
    import {MatchDecorator} from "$lib/service/MatchDecorator";
    import {MatchState} from "$lib/enum/MatchState";

    export let match: Match

    const matchDecorator = new MatchDecorator(match)
    const matchState = matchDecorator.getMatchState()
</script>

<div class="text-xl font-semibold">
    {#if matchState === MatchState.won}
        <span class="text-success-700 dark:text-success-500">W</span>
    {:else if matchState === MatchState.lost}
        <span class="text-error-600 dark:text-error-400">L</span>
    {:else if matchState === MatchState.derby}
        <span class="text-primary-500-400-token text-2xl">♥</span>
    {:else if matchState === MatchState.final}
        <span>F</span>
    {:else if matchState === MatchState.notYetPlayed}
        <span class="text-lg">TBD</span>
    {:else if matchState === MatchState.cancelled}
        <span class="text-lg">PPD</span>
    {/if}
</div>