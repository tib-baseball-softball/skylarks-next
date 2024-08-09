<script lang="ts">
    import {Tab, TabGroup} from "@skeletonlabs/skeleton";
    import type {Match} from "bsm.js";
    import CurrentMatchBlock from "$lib/components/match/CurrentMatchBlock.svelte";

    let tabSet: number = 1;

    export let matchesCurrent: Promise<Match[]>
    export let matchesPrevious: Promise<Match[]>
    export let matchesNext: Promise<Match[]>
</script>

<div class="card overview-card variant-soft-surface dark:border dark:border-tertiary-500-400-token">
    <header class="card-header">
        <h2 class="h3">Aktuelle Spiele</h2>
    </header>
    <section class="p-4">
        <TabGroup justify="justify-center" flex="flex-auto">
            <Tab bind:group={tabSet} name="tab1" value={0}>Vorheriger Spieltag</Tab>
            <Tab bind:group={tabSet} name="tab2" value={1}>Aktueller Spieltag</Tab>
            <Tab bind:group={tabSet} name="tab3" value={2}>Nächster Spieltag</Tab>
            <!-- Tab Panels --->
            <svelte:fragment slot="panel">
                {#if tabSet === 0}
                    <CurrentMatchBlock matches="{matchesPrevious}"/>
                {:else if tabSet === 1}
                    <CurrentMatchBlock matches="{matchesCurrent}"/>
                {:else if tabSet === 2}
                    <CurrentMatchBlock matches="{matchesNext}"/>
                {/if}
            </svelte:fragment>
        </TabGroup>
    </section>
    <footer class="card-footer flex justify-end">
        <a href="/gamecenter" class="btn variant-filled-primary dark:variant-ghost-primary px-10">Gamecenter</a>
    </footer>
</div>