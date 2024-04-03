<script lang="ts">
    import Text from "$lib/components/contentObjects/Text.svelte";

    export let data;

    $: page = data.pageObject
    $: contentObjects = data.contentObjects ?? []
</script>

<article class="prose dark:prose-invert max-w-none">
    <h1 class="mt-4">{page.meta.title}</h1>

    {#each contentObjects as co}
        {#if co.type === "text"}
            <Text {co}></Text>
        {:else }
            <h2 class="h2">{co.content.header}</h2>
            <h3>{co.content.subheader}</h3>
            <p>{@html co.content.bodytext}</p>
        {/if}
    {/each}
</article>