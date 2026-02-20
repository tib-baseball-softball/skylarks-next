<script lang="ts">
  import type {Snippet} from "svelte";
  import {invalidate} from "$app/navigation";
  import TabsRadioGroup from "$lib/dp/components/formElements/TabsRadioGroup.svelte";
  import Switch from "$lib/dp/components/formElements/Switch.svelte";
  //@ts-ignore
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import {client} from "$lib/dp/client.svelte.js";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {Extension} from "$lib/dp/types/ExpandedResponse.js";
  import type {ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";
  import type {LocationsResponse, UniformsetsResponse} from "$lib/dp/types/pb-types.ts";
  import Flatpickr from "$lib/dp/components/formElements/Flatpickr.svelte";
  import {Collection} from "$lib/dp/enum/Collection.ts";

  interface Props {
    event: ExpandedEvent | null;
    clubID: string;
    teamID: string;
    triggerContent: Snippet;
    triggerVariant?: "filled-primary" | "tonal-primary" | "tonal-secondary" | "tonal-tertiary" | "tonal-surface";
    triggerSize?: "default" | "sm";
    triggerIcon?: boolean;
    triggerSpaced?: boolean;
  }

  const {
    event,
    clubID,
    teamID,
    triggerContent,
    triggerVariant = "tonal-primary",
    triggerSize = "default",
    triggerIcon = false,
    triggerSpaced = false,
  }: Props = $props();

  let open = $state(false);

  function formFromProps(data: ExpandedEvent | null) {
    return data ?? {
      id: "",
      title: "",
      starttime: "",
      meetingtime: "",
      endtime: "",
      desc: "",
      location: "",
      type: "game",
      attire: "",
      cancelled: false,
      bsm_id: 0,
      team: teamID,
    };
  }

  let form: Extension<
    Partial<ExpandedEvent>,
    {
      starttime: string
      endtime: string
      meetingtime: string
      type: string
    }
  > = $derived.by(() => {
    const formData = $state(formFromProps(event));
    return formData;
  });

  const attireOptions = $derived(client.collection(Collection.UniformSets).getFullList<UniformsetsResponse>({
    filter: `club = "${clubID}"`,
    requestKey: `uniformsets-${clubID}`,
  }));

  const locationOptions = $derived(client.collection(Collection.Locations).getFullList<LocationsResponse>({
    filter: `club = "${clubID}"`,
    requestKey: `location-options-${clubID}`,
  }));

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ExpandedEvent | null = null;

    try {
      if (form.id) {
        result = await client.collection("events").update<ExpandedEvent>(form.id, form);
      } else {
        result = await client.collection("events").create<ExpandedEvent>(form);
      }
    } catch {
      toastController.triggerGenericFormErrorMessage("Event");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Event");
      open = false;
    }
    await invalidate("event:list");
  }
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger
    class={[
      "btn",
      "trigger-button",
      `trigger-variant-${triggerVariant}`,
      triggerSize === "sm" && "btn-sm",
      triggerIcon && "btn-icon",
      triggerSpaced && "trigger-spaced",
      triggerVariant === "filled-primary" && "preset-filled-primary-500",
      triggerVariant === "tonal-primary" && "preset-tonal-primary border-primary",
      triggerVariant === "tonal-secondary" && "preset-tonal-secondary border-secondary",
      triggerVariant === "tonal-tertiary" && "preset-tonal-tertiary border-tertiary",
      triggerVariant === "tonal-surface" && "preset-tonal-surface",
    ]}
  >
    {@render triggerContent()}
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header class="text-xl font-semibold">
      {#if form.id}
        <h2 class="h3">Edit Event "{form?.title}"</h2>
      {:else}
        <h2 class="h3">Create new Event</h2>
      {/if}
    </header>

    <form class="mt-4 space-y-3" onsubmit={submitForm}>
      <div class="edit-form-grid">
        <input
          autocomplete="off"
          bind:value={form.id}
          class="input"
          name="id"
          readonly
          type="hidden"
        />

        <label class="label">
          <span>Title</span>
          <input
            bind:value={form.title}
            class="input"
            name="title"
            required
            type="text"
          />
        </label>

        <label class="label">
          <span>BSM ID</span>
          <input
            bind:value={form.bsm_id}
            class="input"
            name="bsm_id"
            readonly
            type="text"
          />
        </label>

        <label class="label">
          <span>Start</span>
          <Flatpickr
            bind:value={form.starttime}
            options={Object.assign(DateTimeUtility.datePickerOptions, {
                      static: true, // render the picker as a child element to the form to work in a sheet portal context
                  })}
            required={true}
          />
        </label>

        <label class="label">
          <span>Meeting</span>
          <Flatpickr
            bind:value={form.meetingtime}
            options={Object.assign(DateTimeUtility.datePickerOptions, {
                      static: true,
                  })}
          />
        </label>

        <label class="label">
          <span>End</span>
          <Flatpickr
            bind:value={form.endtime}
            options={Object.assign(DateTimeUtility.datePickerOptions, {
                      static: true,
                  })}
          />
        </label>

        <span></span>

        <label class="label field-wide">
          Description
          <textarea bind:value={form.desc} class="textarea" data-testid="event-form-textarea-desc"
                    name="desc"
          ></textarea>
        </label>

        <label class="label field-wide">
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

        <TabsRadioGroup
          bind:value={form.type}
          label="Type"
          name="type"
          options={["game", "practice", "misc"]}
          required={true}
        />

        {#await attireOptions then options}
          <label class="label field-wide">
            Uniform Set
            <select class="select" bind:value={form.attire} data-testid="event-form-select-attire">
              {#each options as option}
                <option value={option.id}>{option.name}</option>
              {/each}
            </select>
          </label>
        {/await}

        <Switch
          bind:checked={form.cancelled}
          name="cancelled"
        >
          Cancelled
        </Switch>
      </div>

      <hr/>

      <div class="flex justify-center gap-3">
        <button class="mt-2 btn preset-tonal-primary border border-primary-500" data-testid="event-form-submit-button"
                type="submit">
          Submit
        </button>
      </div>
    </form>
  </Sheet.Content>
</Sheet.Root>

<style>
  .trigger-button {
    border-style: solid;
    border-width: 1px;
  }

  .trigger-variant-filled-primary {
    border-color: transparent;
  }

  .trigger-variant-tonal-primary {
    border-color: var(--color-primary-500);
  }

  .trigger-variant-tonal-secondary {
    border-color: var(--color-secondary-500);
  }

  .trigger-variant-tonal-tertiary {
    border-color: var(--color-tertiary-500);
  }

  .trigger-variant-tonal-surface {
    border-color: var(--color-surface-500);
  }

  .trigger-spaced {
    margin-block: calc(var(--spacing) * 3);
  }

  hr {
    margin-block: calc(var(--spacing) * 5);
  }
</style>
