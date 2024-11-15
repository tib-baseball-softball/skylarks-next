<script lang="ts">
  import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";
  import EventTeaser from "$lib/components/diamondplanner/event/EventTeaser.svelte";
  import Paginator from "$lib/pocketbase/Paginator.svelte";
  import {UsersGroupOutline,} from "flowbite-svelte-icons";
  import {goto} from "$app/navigation";
  import type {CustomAuthModel, EventType} from "$lib/model/ExpandedResponse.js";
  import {RadioGroup, RadioItem,} from "@skeletonlabs/skeleton";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import TeamAdminSection from "$lib/components/diamondplanner/team/TeamAdminSection.svelte";

  let {data} = $props();
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

  $effect.pre(() => {
    console.log(showEvents);
    reloadWithQuery();
  });

  const authModel = authSettings.record as CustomAuthModel
</script>

<h1 class="h2">{data.team.name} ({data.team?.expand?.club.name})</h1>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 mb-3">
    <article class="card variant-soft-surface lg:col-span-2">
        <header class="card-header">
            <h2 class="h4 font-medium">Team Description</h2>
        </header>
        <section class="p-4">{@html data.team.description}</section>
    </article>

    <TeamTeaserCard team={data.team} link={false}/>
</div>

<h2 class="h3">Team Events</h2>

<div
        class="flex flex-wrap gap-4 lg:gap-8 variant-soft-surface justify-between px-4 py-3 rounded-token"
>
    <label
            class="flex items-center gap-2 flex-grow justify-between md:flex-grow-0"
    >
        Timeframe
        <RadioGroup>
            <RadioItem
                    checked
                    name="radio-next"
                    value="next"
                    bind:group={showEvents}
            >
                Next
            </RadioItem>
            <RadioItem name="radio-past" value="past" bind:group={showEvents}>
                Past
            </RadioItem>
        </RadioGroup>
    </label>

    <label class="label flex items-center gap-2 flex-grow justify-between md:flex-grow-0">
        Sort
        <RadioGroup>
            <RadioItem checked name="radio-asc" value="asc" bind:group={sorting}>Ascending</RadioItem>
            <RadioItem checked name="radio-desc" value="desc" bind:group={sorting}>Descending</RadioItem>
        </RadioGroup>
    </label>

    <label class="label flex items-center gap-2 flex-grow justify-between md:flex-grow-0">
        Type
        <RadioGroup padding="px-2 md:px-4 py-1">
            <RadioItem padding="px-4 py-1" checked name="radio-any" value="any" bind:group={showTypes}>All</RadioItem>
            <RadioItem checked name="radio-game" value="game" bind:group={showTypes}>Game</RadioItem>
            <RadioItem checked name="radio-practice" value="practice" bind:group={showTypes}>Practice</RadioItem>
            <RadioItem checked name="radio-misc" value="misc" bind:group={showTypes}>Other</RadioItem>
        </RadioGroup>
    </label>
</div>

<div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-3 lg:gap-4 2xl:gap-5">
    {#each $events.items as event}
        <EventTeaser {event} link={true}/>
    {:else}
        <p>No events available with the current filters.</p>
    {/each}
</div>

<Paginator store={events} showIfSinglePage={false}/>

<hr class="!my-8"/>
<section class="space-y-2 lg:space-y-4">
    <header>
        <h2 class="h3">Links</h2>
    </header>

    <div class="flex flex-wrap items-center gap-2 lg:gap-3">
        <a
                href="/account/team/{data.team.id}/members"
                class="btn variant-ghost-tertiary"
        >
            <UsersGroupOutline/>
            <span>Player List</span>
        </a>
    </div>
</section>

{#if data.team.admins.includes(authModel?.id)}
    <hr class="my-2"/>

    <TeamAdminSection team={data.team}/>
{/if}
