<script lang="ts">
    import type {Match} from "bsm.js";
    import {MatchDecorator} from "$lib/service/MatchDecorator";
    import {MatchState} from "$lib/enum/MatchState";

    interface Props {
        match: Match;
    }

    let { match }: Props = $props();

    const matchDecorator = new MatchDecorator(match)
    const matchState = matchDecorator.getMatchState()
</script>

<div class="text-lg font-semibold">
    {#if matchState === MatchState.won}
        <span class="text-success-700 dark:text-success-500">W</span>
    {:else if matchState === MatchState.lost}
        <span class="text-error-600 dark:text-error-400">L</span>
    {:else if matchState === MatchState.derby}
        <span class="text-primary-600-400 text-2xl">♥</span>
    {:else if matchState === MatchState.final}
        <span>F</span>
    {:else if matchState === MatchState.notYetPlayed}
        <span class="text-md">TBD</span>
    {:else if matchState === MatchState.cancelled}
        <span class="text-md">PPD</span>
    {/if}
</div>