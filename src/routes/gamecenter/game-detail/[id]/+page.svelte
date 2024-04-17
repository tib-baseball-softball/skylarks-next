<script lang="ts">
    import {ProgressRadial, Tab, TabGroup} from "@skeletonlabs/skeleton";

    export let data

    let tabSet: number = 0
</script>

<p>This is the detail for page with ID {data.match.id}</p>

<section>
    <TabGroup>
        <Tab bind:group={tabSet} name="tab1" value={0}>Spieldaten</Tab>
        <Tab bind:group={tabSet} name="tab2" value={1}>Boxscore</Tab>
        <!-- Tab Panels --->
        <svelte:fragment slot="panel">
            {#if tabSet === 0}
                (tab panel 1 contents)
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

