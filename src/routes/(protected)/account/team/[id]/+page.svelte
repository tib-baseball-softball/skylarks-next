<script lang="ts">
    import {ProgressRadial} from "@skeletonlabs/skeleton";
    import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";

    let {data} = $props()
</script>

<h1 class="h2">{data.team.name} ({data.team?.expand?.club.name})</h1>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 mb-3">

    <article class="card variant-ghost-surface lg:col-span-2">
        <header class="card-header"><h2 class="h3">Teambeschreibung</h2></header>
        <section class="p-4">{@html data.team.description}</section>
    </article>

    <TeamTeaserCard team={data.team}/>
</div>

<h2 class="h3">Aktuelle Events</h2>
{#await data.events}
    <ProgressRadial/>
{:then events}
    <ul>
        {#each events.items as event}
        <li>
            <a class="anchor" href="/account/event/{event.id}">{event.title}</a>
        </li>
        {/each}
    </ul>
{:catch error}
    <p>Fehler beim Laden: {error.message}</p>
{/await}