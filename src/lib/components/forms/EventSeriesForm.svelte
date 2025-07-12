<script lang="ts">
  import {invalidate} from "$app/navigation";
  import type {EventSeriesCreationData, ExpandedTeam} from "$lib/model/ExpandedResponse.ts";
  import {client} from "$lib/pocketbase/index.svelte.js";
  import {DateTimeUtility} from "$lib/service/DateTimeUtility.js";
  import Flatpickr from "$lib/components/utility/Flatpickr.svelte";
  import type {LocationsResponse} from "$lib/model/pb-types.ts";
  import {toastController} from "$lib/service/ToastController.svelte.ts";

  interface Props {
    eventSeries: EventSeriesCreationData | null,
    team: ExpandedTeam,
    showForm: boolean
  }

  let {eventSeries, team, showForm = $bindable()}: Props = $props();


  const formDefault = $state({
    id: "",
    title: "",
    interval: 7,
    duration: 0,
    series_start: "",
    series_end: "",
    desc: "",
    location: "",
    team: team?.id,
  });

  const form: EventSeriesCreationData = $derived(
      eventSeries ?? formDefault,
  );

  const locationOptions = client.collection("locations").getFullList<LocationsResponse>({
    filter: `club = "${team.club}"`,
    requestKey: `locations-for-eventSeries-${team.id}`,
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: EventSeriesCreationData | null = null;

    try {
      if (form.id) {
        result = await client
            .collection("eventseries")
            .update<EventSeriesCreationData>(form.id, form);
      } else {
        result = await client
            .collection("eventseries")
            .create<EventSeriesCreationData>(form);
      }
    } catch {
      toastController.triggerGenericFormErrorMessage("Event Series");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Event Series");
      showForm = false;
      await invalidate("event:list");
    }
  }
</script>

<h3 class="h3">
  {#if eventSeries}
    Edit Event Series "{eventSeries.title}"
  {:else}
    Create new Event Series for {team.name}
  {/if}
</h3>
<form onsubmit={submitForm} class="mt-4 space-y-3">
  <div class="grid grid-cols-1 md:grid-cols-2 gap-2 lg:gap-3 xl:gap-4 2xl:gap-5">
    <input
            name="id"
            autocomplete="off"
            class="input"
            type="hidden"
            readonly
            bind:value={form.id}
    />

    <label class="label md:col-span-2">
      Series First Occurrence
      <Flatpickr
              bind:value={form.series_start}
              options={Object.assign(DateTimeUtility.datePickerOptions, {
                  static: true, // render the picker as a child element to the form to work in a sheet portal context
              })}
      />
    </label>

    <span class="md:col-span-2 text-sm font-light block preset-outlined-primary-500 p-2 my-1 rounded-container">
      This is the first event in your new series. It will have the exact start time you enter here.<br>
      The next events are then created <em>RECURRING INTERVAL</em> days after the previous one and all will have
      the end time set to <em>DURATION</em> minutes after their start time.
    </span>

    <label class="label md:col-span-2">
      Series End
      <Flatpickr
              bind:value={form.series_end}
              options={Object.assign(DateTimeUtility.datePickerOptionsNoTime, {
                  static: true,
              })}
      />
      <span class="text-sm font-light block">No events will be created after this date.</span>
    </label>

    <label class="label md:col-span-2 mt-3 md:mt-0">
      Title
      <input
              name="title"
              class="input"
              required
              type="text"
              bind:value={form.title}
      />
    </label>

    <label class="label">
      Recurring Interval
      <input
              name="interval"
              type="number"
              class="input"
              bind:value={form.interval}
      >
      <span class="text-sm font-light block">Interval between each series occurrence (in days).</span>
    </label>

    <label class="label">
      Duration
      <input
              name="duration"
              type="number"
              class="input"
              bind:value={form.duration}
      >
      <span class="text-sm font-light block">Duration of each event to determine end time (in minutes).</span>
    </label>

    <label class="label">
      Description
      <textarea name="desc" class="textarea" bind:value={form.desc}
      ></textarea>
    </label>

    <label class="label">
      Location
      <select
              class="select"
              bind:value={form.location}
      >
        {#await locationOptions then options}
          <option value="">None</option>
          {#each options as option}
            <option value={option.id}>{option?.address_addon ? option.address_addon : "No additional name"}
              ({option.name}), {option.street}, {option.postal_code} {option.city}, {option.country}</option>
          {/each}
        {/await}
      </select>
    </label>

  </div>

  <hr class="my-5!"/>

  <div class="flex justify-center gap-3">
    <button type="button" class="btn preset-tonal border border-surface-500" onclick={() => {if (showForm) showForm = false}}>
      Cancel
    </button>

    <button type="submit" class="btn preset-filled-primary-500">
      Submit
    </button>
  </div>
</form>
