<script lang="ts">
  import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";
  import EventTeaser from "$lib/components/diamondplanner/event/EventTeaser.svelte";
  import Paginator from "$lib/pocketbase/Paginator.svelte";
  import {goto} from "$app/navigation";
  import type {CustomAuthModel, EventType} from "$lib/model/ExpandedResponse.js";
  import {Segment} from "@skeletonlabs/skeleton-svelte";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import TeamAdminSection from "$lib/components/diamondplanner/team/TeamAdminSection.svelte";
  import {Users} from "lucide-svelte";
  import AnnouncementSectionContent from "$lib/components/announcements/AnnouncementSectionContent.svelte";
  import type {PageProps} from "./$types";
  import AnnouncementForm from "$lib/components/forms/AnnouncementForm.svelte";

  let {data}: PageProps = $props();
  const events = $derived(data.events);
  let currentPage = $derived($events.page);
  let announcementStore = $derived(data.announcementStore);

  let showEvents = $state("next");
  let sorting: "asc" | "desc" | string = $state("asc");
  let showTypes: EventType | "any" | string = $state("any");

  const reloadWithQuery = () => {
    let queryString = `?timeframe=${showEvents}&page=${currentPage}&sort=${sorting}&type=${showTypes}`;

    goto(queryString, {
      noScroll: true,
    });
  };

  $effect.pre(() => {
    console.log(showEvents);
    reloadWithQuery();
  });

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canEdit = $derived(data.team.admins.includes(authRecord?.id) || data.team?.expand?.club?.admins.includes(authRecord?.id))
</script>

<h1 class="h1 my-3!">{data.team.name} ({data.team?.expand?.club.name})</h1>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 my-6!">
  <article class="card preset-tonal-surface lg:col-span-2">
    <header class="card-header">
      <h2 class="h4 font-medium">Team Description</h2>
    </header>
    <section class="p-4">{@html data.team.description}</section>
  </article>

  <TeamTeaserCard team={data.team} link={false}/>
</div>

<section class="my-8! space-y-4">
  <header>
    <h2 class="h2 mb-3">Announcements</h2>
  </header>

  <AnnouncementSectionContent store={announcementStore} />

  {#if canEdit}
    <AnnouncementForm
            announcement={null}
            team={data.team}
            club={null}
            buttonClasses="btn preset-filled-primary-500"
            showLabel={true}
    />
  {/if}
</section>

<h2 class="h2">Team Events</h2>

<div
        class="flex flex-wrap gap-4 xl:flex-nowrap preset-tonal-surface justify-between px-4 py-3 rounded-base text-sm lg:text-base"
>
  <label
          class="flex items-center gap-2 grow justify-between xl:justify-start md:grow-0"
  >
    Timeframe
    <Segment name="timeframe" value={showEvents} onValueChange={(e) => (showEvents = e.value ?? "next")} padding="p-1!"
             classes="event-segment-container">
      <Segment.Item value="next">
        Next
      </Segment.Item>
      <Segment.Item value="past">
        Past
      </Segment.Item>
    </Segment>
  </label>

  <label class="flex items-center gap-2 justify-between xl:justify-start grow md:grow-0">
    Sort
    <Segment name="sorting" value={sorting} onValueChange={(e) => (sorting = e.value ?? "asc")} padding="p-1!"
             classes="flex-wrap event-segment-container">
      <Segment.Item value="asc">Ascending</Segment.Item>
      <Segment.Item value="desc">Descending</Segment.Item>
    </Segment>
  </label>

  <label class="flex items-center gap-2 justify-between xl:justify-start grow md:grow-0">
    Type
    <Segment name="type" value={showTypes} onValueChange={(e) => (showTypes = e.value ?? "any")} gap="gap-1"
             padding="p-1!" classes="flex-wrap event-segment-container">
      <Segment.Item value="any">All</Segment.Item>
      <Segment.Item value="game">Game</Segment.Item>
      <Segment.Item value="practice">Practice</Segment.Item>
      <Segment.Item value="misc">Other</Segment.Item>
    </Segment>
  </label>
</div>

<div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-3 lg:gap-4 2xl:gap-5">
  {#each $events.items as event (event.id)}
    <div>
      <EventTeaser {event} link={true}/>
    </div>
  {:else}
    <p>No events available with the current filters.</p>
  {/each}
</div>

<Paginator store={events} showIfSinglePage={false}/>

<hr class="my-8!"/>
<section class="space-y-2 lg:space-y-4">
  <header>
    <h2 class="h3">Links</h2>
  </header>

  <div class="flex flex-wrap items-center gap-2 lg:gap-3">
    <a
            href="/account/team/{data.team.id}/members"
            class="btn preset-tonal-tertiary border border-tertiary-500"
    >
      <Users/>
      <span>Player List</span>
    </a>
  </div>
</section>

{#if canEdit}
  <hr class="my-2"/>

  <TeamAdminSection team={data.team} eventSeries={data.eventSeries}/>
{/if}

<style>
    :global {
        .event-segment-container {
            border: 1px solid var(--color-surface-600-400);

            .btn {
                padding: 0.25rem 0.6rem;
            }
        }
    }
</style>
