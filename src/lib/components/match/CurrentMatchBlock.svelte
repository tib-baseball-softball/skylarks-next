<script lang="ts">
    import type {Match} from "bsm.js";
    import {ProgressRadial} from "@skeletonlabs/skeleton";
    import CurrentMatchRow from "$lib/components/match/CurrentMatchRow.svelte";

    interface Props {
        matches: Promise<Match[]>;
    }

    let { matches }: Props = $props();
</script>

{#await matches}
    <p>Loading matches...</p>
    <ProgressRadial/>
{:then results}
    <section class="grid md:grid-flow-col grid-rows-4 grid-cols-1 md:grid-cols-2 gap-y-2">
        {#each results as match}
            <CurrentMatchRow {match}/>
        {/each}

        {#if results.length === 0}
            <p>Keine Spiele in diesem Zeitraum.</p>
        {/if}
    </section>

{:catch error}
    <p>error loading matches: {error.message}</p>
{/await}