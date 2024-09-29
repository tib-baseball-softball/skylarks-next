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

  let { data } = $props();
  const events = $derived(data.events);

  let showEvents = $state("next");

  const reloadWithQuery = () => {
    let queryString = `?filter=${showEvents}`;

    goto(queryString, {
        noScroll: true,
    });
  };

  $effect.pre(() => {
    console.log(showEvents)
    reloadWithQuery()
  })
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

<div class="flex gap-4">
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
</div>

<div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-3">
  {#each $events.items as event}
    <EventTeaser {event} link={true} />
  {:else}
    <p>Keine Events verfügbar.</p>
  {/each}
</div>

<Paginator store={events} showIfSinglePage={false} />

<hr class="!my-8" />
<section class="space-y-2 lg:space-y-4">
  <header>
    <h2 class="h3">Actions and Links</h2>
  </header>

  <div class="flex flex-wrap items-center gap-2 lg:gap-3">
    <button class="btn variant-ghost-primary">
      <CalendarPlusOutline />
      <span>New Event</span>
    </button>

    <a href="#" class="btn variant-ghost-tertiary">
      <UsersGroupOutline />
      <span>Player List</span>
    </a>

    <a href="#" class="btn variant-ghost-surface">
      <CogOutline />
      <span>Settings</span>
    </a>
  </div>
</section>
