<script lang="ts">
    import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";

    let {data} = $props()
</script>

<h1 class="h1">Mein Account</h1>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 mb-3">
    <div class="card lg:col-span-2">
        <header class="card-header">(header)</header>
        <section class="p-4">(content)</section>
        <footer class="card-footer">(footer)</footer>
    </div>
    <div class="card">
        <header class="card-header">(header)</header>
        <section class="p-4">(content)</section>
        <footer class="card-footer">(footer)</footer>
    </div>
    <div class="card">
        <header class="card-header">(header)</header>
        <section class="p-4">(content)</section>
        <footer class="card-footer">(footer)</footer>
    </div>
    <div class="card">
        <header class="card-header">(header)</header>
        <section class="p-4">(content)</section>
        <footer class="card-footer">(footer)</footer>
    </div>
    <div class="card">
        <header class="card-header">(header)</header>
        <section class="p-4">(content)</section>
        <footer class="card-footer">(footer)</footer>
    </div>
</div>

<h2 class="h2 mt-3">Meine Teams</h2>

{#await data.teams then teams}
    <div class="grid grid-cols-1 md:grid-cols-2 gap-3 mb-3">
        {#each teams as team}
            <a href="/account/team/{team.id}">
                <TeamTeaserCard {team} link={true}/>
            </a>
        {/each}
    </div>
{/await}