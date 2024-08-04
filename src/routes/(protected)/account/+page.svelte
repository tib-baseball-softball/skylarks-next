<script lang="ts">
    let {data} = $props()
</script>

<p>Accountseite</p>

{#await data.clubs then clubs}
    {#each clubs as club}
        <p>{club.name}</p>
    {/each}
{/await}