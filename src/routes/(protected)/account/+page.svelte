<script lang="ts">
    let {data} = $props()
</script>

<h1 class="h1">Mein Account</h1>

<p>...</p>

<h2 class="h2">Meine Teams</h2>

{#await data.teams then teams}
    {#each teams as team}
        <a class="anchor" href="/account/team/{team.id}">{team.name}</a>
    {/each}
{/await}