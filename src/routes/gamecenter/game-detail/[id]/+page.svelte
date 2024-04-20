<script lang="ts">
    import {ProgressRadial, Tab, TabGroup} from "@skeletonlabs/skeleton";
    import MatchMainInfoSection from "$lib/components/match/MatchMainInfoSection.svelte";

    export let data

    $: match = data.match
    let tabSet: number = 0
</script>

<h1 class="h2 mt-3">Details zu Spiel {match.match_id}</h1>
<section class="my-3 lg:my-5">
    <MatchMainInfoSection {match}/>
</section>

<section>
    <TabGroup justify="justify-center">
        <Tab bind:group={tabSet} name="tab1" value={0}>weitere Spieldaten</Tab>
        <Tab bind:group={tabSet} name="tab2" value={1}>Boxscore</Tab>
        <!-- Tab Panels --->
        <svelte:fragment slot="panel">
            {#if tabSet === 0}
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 place-items-stretch gap-3 md:gap-4 lg:gap-5">
                    <div class="card variant-ghost-surface p-3">
                        <p class="mb-2 font-medium">Info</p>
                        <p>(content)</p>
                    </div>
                    <div class="card variant-ghost-surface p-3">
                        <p class="mb-2 font-medium">Info</p>
                        <p>(content)</p>
                    </div>
                    <div class="card variant-ghost-surface p-3">
                        <p class="mb-2 font-medium">Info</p>
                        <p>(content)</p>
                    </div>
                </div>
            {:else if tabSet === 1}
                {#await data.singleGameStats}
                    <ProgressRadial />
                {:then boxscore}
                    {@html boxscore.full_boxscore_html}
                {:catch error}
                    <p>error loading boxscore: {error.message}</p>
                {/await}
            {/if}
        </svelte:fragment>
    </TabGroup>
</section>

