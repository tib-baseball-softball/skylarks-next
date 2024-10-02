<script lang="ts">
    import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";
    import EventTeaser from "$lib/components/diamondplanner/event/EventTeaser.svelte";
    import Paginator from "$lib/pocketbase/Paginator.svelte";
    import {
        CalendarPlusOutline,
        CogOutline,
        UsersGroupOutline,
    } from "flowbite-svelte-icons";
    import { goto } from "$app/navigation";
    import type { EventType } from "$lib/model/ExpandedResponse.js";
    import {
        getDrawerStore,
        type DrawerSettings,
    } from "@skeletonlabs/skeleton";

    let { data } = $props();
    const events = $derived(data.events);
    let currentPage = $derived($events.page);

    let showEvents = $state("next");
    let sorting: "asc" | "desc" = $state("asc");
    let showTypes: EventType | "any" = $state("any");

    const reloadWithQuery = () => {
        let queryString = `?timeframe=${showEvents}&page=${currentPage}&sort=${sorting}&type=${showTypes}`;

        goto(queryString, {
            noScroll: true,
        });
    };

    const drawerStore = getDrawerStore();
    const settings: DrawerSettings = $derived({
        id: "event-form",
        position: "right",
        width: "w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%]",
        meta: {
            event: null,
            club: data.team?.club,
            team: data.team,
        },
    });

    $effect.pre(() => {
        console.log(showEvents);
        reloadWithQuery();
    });
</script>

<h1 class="h2">{data.team.name} ({data.team?.expand?.club.name})</h1>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 mb-3">
    <article class="card variant-soft-surface lg:col-span-2">
        <header class="card-header">
            <h2 class="h4 font-medium">Team Description</h2>
        </header>
        <section class="p-4">{@html data.team.description}</section>
    </article>

    <TeamTeaserCard team={data.team} link={false} />
</div>

<h2 class="h3">Team Events</h2>

<div class="flex flex-wrap gap-4 variant-soft-surface px-4 py-3 rounded-token">
    <label class="flex items-center space-x-2">
        <input
            class="radio"
            type="radio"
            checked
            name="radio-next"
            value="next"
            bind:group={showEvents}
        />
        <p>Next</p>
    </label>

    <label class="flex items-center space-x-2">
        <input
            class="radio"
            type="radio"
            name="radio-past"
            value="past"
            bind:group={showEvents}
        />
        <p>Past</p>
    </label>

    <span class="divider-vertical h-10 mx-3"></span>

    <label class="label flex items-center gap-2">
        Sort
        <select class="select" bind:value={sorting} onchange={reloadWithQuery}>
            <option value="asc">Ascending</option>
            <option value="desc">Descending</option>
        </select>
    </label>

    <span class="divider-vertical h-10 mx-3"></span>

    <label class="label flex items-center gap-2">
        Show
        <select
            class="select"
            bind:value={showTypes}
            onchange={reloadWithQuery}
        >
            <option value="any">All</option>
            <option value="game">Games</option>
            <option value="practice">Practices</option>
            <option value="misc">Other</option>
        </select>
    </label>
</div>

<div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-3">
    {#each $events.items as event}
        <EventTeaser {event} link={true} />
    {:else}
        <p>No events available with the current filters.</p>
    {/each}
</div>

<Paginator store={events} showIfSinglePage={false} />

<hr class="!my-8" />
<section class="space-y-2 lg:space-y-4">
    <header>
        <h2 class="h3">Actions and Links</h2>
    </header>

    <div class="flex flex-wrap items-center gap-2 lg:gap-3">
        <button
            class="btn variant-ghost-primary"
            onclick={() => drawerStore.open(settings)}
        >
            <CalendarPlusOutline />
            <span>New Event</span>
        </button>

        <a
            href="/account/team/{data.team.id}/members"
            class="btn variant-ghost-tertiary"
        >
            <UsersGroupOutline />
            <span>Player List</span>
        </a>

        <a href="#" class="btn variant-ghost-surface">
            <CogOutline />
            <span>Settings</span>
        </a>
    </div>
</section>
