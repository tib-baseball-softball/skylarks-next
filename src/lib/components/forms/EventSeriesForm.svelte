<script lang="ts">
  import {invalidate} from "$app/navigation";
  import type {EventSeriesCreationData, ExpandedTeam} from "$lib/model/ExpandedResponse.ts";
  import {client} from "$lib/pocketbase/index.svelte.js";
  import {DateTimeUtility} from "$lib/service/DateTimeUtility.js";
  import Flatpickr from "$lib/components/utility/Flatpickr.svelte";
  import {Weekday} from "$lib/types/Weekday.ts";
  import type {LocationsResponse} from "$lib/model/pb-types.ts";
  import type {Toast} from "$lib/types/Toast.ts";
  import {toastController} from "$lib/service/ToastController.svelte.ts";

  interface Props {
    eventSeries: EventSeriesCreationData | null,
    team: ExpandedTeam,
    showForm: boolean
  }

  let {eventSeries, team, showForm = $bindable()}: Props = $props();

  const toastSettingsSuccess: Toast = {
    message: "Event Series saved successfully.",
    background: "preset-filled-success-500",
  };

  const toastSettingsError: Toast = {
    message: "An error occurred while saving the event series.",
    background: "preset-filled-error-500",
  };

  const formDefault = $state({
    id: "",
    title: "",
    starttime: "",
    endtime: "",
    interval: 7,
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

    try {
      form.starttime = DateTimeUtility.convertTimeToUTC(form.starttime);

      if (form.endtime) {
        form.endtime = DateTimeUtility.convertTimeToUTC(form.endtime);
      }
    } catch (error) {
      console.error("Invalid format for starttime or endtime in form.");
      toastController.trigger(toastSettingsError);
      return;
    }

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
      toastController.trigger(toastSettingsError);
    }

    if (result) {
      toastController.trigger(toastSettingsSuccess);
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
  <div class="grid grid-cols-2 gap-2 lg:gap-3 xl:gap-4 2xl:gap-5">
    <input
            name="id"
            autocomplete="off"
            class="input"
            type="hidden"
            readonly
            bind:value={form.id}
    />

    <label class="label">
      Create Events after
      <Flatpickr
              bind:value={form.series_start}
              options={Object.assign(DateTimeUtility.datePickerOptionsNoTime, {
                  static: true, // render the picker as a child element to the form to work in a sheet portal context
              })}
      />
    </label>

    <label class="label">
      Create Events before
      <Flatpickr
              bind:value={form.series_end}
              options={Object.assign(DateTimeUtility.datePickerOptionsNoTime, {
                  static: true,
              })}
      />
    </label>

    <span class="col-span-2 text-sm font-light block">
      This selects boundaries between which events will be created.
      The start date does not have to be the actual training day.
    </span>

    <label class="label col-span-2 md:col-span-1 mt-3 md:mt-0">
      Title
      <input
              name="title"
              class="input"
              required
              type="text"
              bind:value={form.title}
      />
    </label>

    {#snippet timeFormatHint()}
      <span class="text-sm font-light block">Format: HH:mm</span>
    {/snippet}

    <label class="label">
      Event start time
      <input
              name="starttime"
              type="time"
              class="input"
              required
              bind:value={form.starttime}
      >
      {@render timeFormatHint()}
    </label>

    <label class="label">
      Event end time
      <input
              name="endtime"
              type="time"
              class="input"
              bind:value={form.endtime}
      >
      {@render timeFormatHint()}
    </label>

    <label class="label">
      Day of Week
      <select
              class="select"
              name="weekday"
              bind:value={form.weekday}
      >
        <option value={Weekday.Monday}>Monday</option>
        <option value={Weekday.Tuesday}>Tuesday</option>
        <option value={Weekday.Wednesday}>Wednesday</option>
        <option value={Weekday.Thursday}>Thursday</option>
        <option value={Weekday.Friday}>Friday</option>
        <option value={Weekday.Saturday}>Saturday</option>
        <option value={Weekday.Sunday}>Sunday</option>
      </select>
    </label>

    <label class="label">
      Recurring Interval
      <input
              name="interval"
              type="number"
              class="input"
              bind:value={form.interval}
      >
      <span class="text-sm font-light block">Intervals are saved in days.</span>
    </label>

    <label class="label col-span-2">
      Description
      <textarea name="desc" class="textarea" bind:value={form.desc}
      ></textarea>
    </label>

    <label class="label md:col-span-2">
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
