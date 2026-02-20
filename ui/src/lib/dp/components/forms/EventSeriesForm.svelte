<script lang="ts">
  import {invalidate} from "$app/navigation";
  import Flatpickr from "$lib/dp/components/formElements/Flatpickr.svelte";
  import {client} from "$lib/dp/client.svelte.js";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {EventSeriesCreationData, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import type {LocationsResponse} from "$lib/dp/types/pb-types.ts";

  interface Props {
    eventSeries: EventSeriesCreationData | null;
    team: ExpandedTeam;
    showForm: boolean;
  }

  let {eventSeries, team, showForm = $bindable()}: Props = $props();

  function formFromProps(data: EventSeriesCreationData | null): EventSeriesCreationData {
    return data ?? {
      id: "",
      title: "",
      interval: 7,
      duration: 0,
      series_start: "",
      series_end: "",
      desc: "",
      location: "",
      team: team?.id,
    };
  }

  let form: EventSeriesCreationData = $derived.by(() => {
    const formData = $state(formFromProps(eventSeries));
    return formData;
  });

  const isNewSeries = $derived(form.id === "");

  const locationOptions = $derived(client.collection("locations").getFullList<LocationsResponse>({
    filter: `club = "${team.club}"`,
    requestKey: `locations-for-eventSeries-${team.id}`,
  }));

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: EventSeriesCreationData | null = null;

    try {
      if (form.id) {
        result = await client.collection("eventseries").update<EventSeriesCreationData>(form.id, form);
      } else {
        result = await client.collection("eventseries").create<EventSeriesCreationData>(form);
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
<form onsubmit={submitForm}>
  <div class="edit-form-grid">
    {#if !isNewSeries}
      <span class="field-wide text-sm font-light block">
        When editing an event series, only select fields can be changed because editing past events would affect
        participation statistics. If you want to change a long-running practice series, end the previous one by
        setting the end date accordingly and create a new series.
      </span>
    {/if}

    <input
      autocomplete="off"
      bind:value={form.id}
      class="input"
      name="id"
      readonly
      type="hidden"
    />

    {#if isNewSeries}
      <label class="label field-wide">
        Series First Occurrence
        <Flatpickr
          bind:value={form.series_start}
          options={Object.assign(DateTimeUtility.datePickerOptions, {
                  static: true, // render the picker as a child element to the form to work in a sheet portal context
              })}
        />
      </label>

      <span class="field-wide hint hint-prominent">
      This is the first event in your new series. It will have the exact start time you enter here.<br>
      The next events are then created <em>RECURRING INTERVAL</em> days after the previous one and all will have
      the end time set to <em>DURATION</em> minutes after their start time.
    </span>
    {/if}

    <label class="label field-wide">
      Series End
      <Flatpickr
        bind:value={form.series_end}
        options={Object.assign(DateTimeUtility.datePickerOptionsNoTime, {
                  static: true,
              })}
      />
      <span class=" hint">No events will be created after this date.</span>
    </label>

    <label class="label field-wide mt-3 md:mt-0">
      Title
      <input
        bind:value={form.title}
        class="input"
        name="title"
        required
        type="text"
      />
    </label>

    {#if isNewSeries}
      <label class="label">
        Recurring Interval
        <input
          bind:value={form.interval}
          class="input"
          name="interval"
          type="number"
          step="7"
          min="7"
        >
        <span class="hint">Interval between each series occurrence (in days).</span>
      </label>

      <label class="label">
        Duration
        <input
          bind:value={form.duration}
          class="input"
          name="duration"
          type="number"
        >
        <span class="hint">Duration of each event to determine end time (in minutes).</span>
      </label>
    {/if}

    <label class="label">
      Description
      <textarea bind:value={form.desc} class="textarea" name="desc"
      ></textarea>
    </label>

    <label class="label">
      Location
      <select
        bind:value={form.location}
        class="select"
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

  <hr/>

  <div class="submit-container">
    <button class="btn preset-tonal border-surface-500" onclick={() => {if (showForm) showForm = false}}
            type="button">
      Cancel
    </button>

    <button class="btn preset-filled-primary-500" type="submit">
      Submit
    </button>
  </div>
</form>

<style>
  hr {
    margin-block: calc(var(--spacing) * 5);
  }

  .submit-container {
    display: flex;
    justify-content: center;
    gap: calc(var(--spacing) * 3);
  }

  .hint {
    font-size: var(--text-sm);
    font-weight: var(--font-weight-light);
    display: block;
  }

  .hint-prominent {
    padding: calc(var(--spacing) * 2);
    margin-block: calc(var(--spacing) * 1);
    border: 1px solid var(--color-primary-500);
    border-radius: var(--radius-container);
  }

  .edit-form-grid {
    gap: calc(var(--spacing) * 4);
  }
</style>
