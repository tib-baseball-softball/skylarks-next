<script lang="ts">
  import {getDrawerStore} from "@skeletonlabs/skeleton";
  import type {EventseriesResponse} from "$lib/model/pb-types.ts";
  import type {EventSeriesCreationData, ExpandedTeam} from "$lib/model/ExpandedResponse.ts";
  import EventSeriesForm from "$lib/components/forms/EventSeriesForm.svelte";
  import {slide} from "svelte/transition";
  import EventSeriesListItem from "$lib/components/diamondplanner/eventseries/EventSeriesListItem.svelte";
  import DeleteButton from "$lib/components/utility/DeleteButton.svelte";
  import {client} from "$lib/pocketbase/index.svelte.ts";
  import {invalidateAll} from "$app/navigation";
  import {DateTimeUtility} from "$lib/service/DateTimeUtility.ts";
  import {X} from "lucide-svelte";

  const drawerStore = getDrawerStore();

  const eventSeries: EventseriesResponse[] = $derived($drawerStore.meta.eventSeries ?? []);
  const team: ExpandedTeam = $derived($drawerStore.meta.team);

  let showForm = $state(false);
  let selectedEventSeries: EventSeriesCreationData | null = $state(null);

  function setupAndShowForm(eventSeries: EventSeriesCreationData | null) {
    selectedEventSeries = eventSeries;

    // this is pretty ugly, but the form needs to get local times instead of UTC values
    if (eventSeries && selectedEventSeries && eventSeries.endtime) {
      selectedEventSeries.starttime = DateTimeUtility.convertTimeFromUTC(eventSeries.starttime);
      selectedEventSeries.endtime = DateTimeUtility.convertTimeFromUTC(eventSeries.endtime);
    }
    showForm = true;

    // wait for Svelte transition to finish
    setTimeout(() => {
      const scrollTarget = document.getElementById("eventSeriesFormContainer");

      scrollTarget?.scrollIntoView({
        behavior: "smooth",
        block: "start",
        inline: "nearest",
      });
    }, 230)
  }

  function deleteEventSeries(id: string) {
    drawerStore.close();
    client.collection("eventseries").delete(id);
    invalidateAll();
  }
</script>

<article class="p-3 sm:p-4 lg:p-6 space-y-3">
  <div class="flex items-center gap-5">
    <button
            aria-label="cancel and close"
            class="btn variant-ghost-surface"
            onclick={drawerStore.close}
    >
      <X/>
    </button>
    <header class="text-xl font-semibold">
      <h2 class="h3">Manage Event Series for {team?.name}</h2>
    </header>
  </div>

  <h3 class="h3 !mt-6">Active Event Series</h3>

  {#each eventSeries as series}
    <EventSeriesListItem eventSeries={series}/>

    <div class="flex gap-2">
      <button class="btn variant-ghost" onclick={() => setupAndShowForm(series)}>Edit</button>

      <DeleteButton
              id={series.id}
              modelName="Event Series"
              action={deleteEventSeries}
              classes="variant-ghost-error"
              buttonText="Delete"
      />
    </div>
  {/each}

  {#if eventSeries.length === 0}
    <section class="font-light space-y-4">
      <div>No event series have been set up for this team yet.</div>
    </section>
  {/if}
  <button class="btn variant-ghost-primary !mt-6" onclick={() => setupAndShowForm(null)}>
    Add new Event Series
  </button>

  <hr class="!my-6">
  <div id="eventSeriesFormContainer">
    {#if showForm}
      <div transition:slide={{duration: 230}}>
        <EventSeriesForm team={team} eventSeries={selectedEventSeries} bind:showForm={showForm}/>
      </div>
    {/if}
  </div>
</article>
