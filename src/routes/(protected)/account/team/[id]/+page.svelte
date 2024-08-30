<script lang="ts">
    import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";
    import EventTeaser from "$lib/components/diamondplanner/event/EventTeaser.svelte";
    import Paginator from "$lib/pocketbase/Paginator.svelte";

    let {data} = $props()
    const events = $derived(data.events);
</script>

<h1 class="h2">{data.team.name} ({data.team?.expand?.club.name})</h1>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 mb-3">

    <article class="card variant-ghost-surface lg:col-span-2">
        <header class="card-header"><h2 class="h4 font-medium">Teambeschreibung</h2></header>
        <section class="p-4">{@html data.team.description}</section>
    </article>

    <TeamTeaserCard team={data.team} link={false}/>
</div>

<h2 class="h3">Aktuelle Events</h2>

<div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-3">
    {#each $events.items as event}
        <EventTeaser {event} link={true}/>
    {:else }
        <p>Keine Events verfügbar.</p>
    {/each}
</div>

<Paginator store={events} showIfSinglePage={true}/>