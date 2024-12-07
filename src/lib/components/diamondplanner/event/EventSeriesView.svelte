<script lang="ts">
  import {getDrawerStore} from "@skeletonlabs/skeleton";
  import {CloseOutline} from "flowbite-svelte-icons";
  import type {EventseriesResponse} from "$lib/model/pb-types.ts";
  import type {EventSeriesCreationData, ExpandedTeam} from "$lib/model/ExpandedResponse.ts";
  import EventSeriesForm from "$lib/components/forms/EventSeriesForm.svelte";
  import {slide} from "svelte/transition";

  const drawerStore = getDrawerStore();

  const eventSeries: EventseriesResponse[] = $derived($drawerStore.meta.eventSeries ?? [])
  const team: ExpandedTeam = $derived($drawerStore.meta.team)

  let showForm = $state(false)
  let selectedEventSeries: EventSeriesCreationData | null = $state(null)

  function setupAndShowForm(eventSeries: EventSeriesCreationData | null) {
    selectedEventSeries = eventSeries
    showForm = true
  }
</script>

<article class="p-6 space-y-3">
    <div class="flex items-center gap-5">
        <button
                aria-label="cancel and close"
                class="btn variant-ghost-surface"
                onclick={drawerStore.close}
        >
            <CloseOutline/>
        </button>
        <header class="text-xl font-semibold">
            <h2 class="h3">Manage Event Series for {team?.name}</h2>
        </header>
    </div>

    <h3 class="h3 !mt-6">Active event series</h3>

    {#each eventSeries as series}
        <div>{series.title}</div>
    {/each}

    {#if eventSeries.length === 0}
        <section class="font-light space-y-4">
            <div>No event series have been set up for this team yet.</div>

        </section>
    {/if}
    <button class="btn variant-ghost-primary" onclick={() => setupAndShowForm(null)}>
        Add event series
    </button>

    <hr class="!my-6">
    {#if showForm}
        <div transition:slide={{duration: 230}}>
            <EventSeriesForm team={team} eventSeries={selectedEventSeries} bind:showForm={showForm}/>
        </div>
    {/if}
</article>
