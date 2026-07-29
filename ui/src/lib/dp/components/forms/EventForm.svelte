<script lang="ts">
  import { invalidate } from "$app/navigation";
  import { client } from "$lib/dp/client.svelte.js";
  import ISODatePicker from "$lib/dp/components/formElements/ISODatePicker.svelte";
  import MultiSelectCombobox from "$lib/dp/components/formElements/MultiSelectCombobox.svelte";
  import Switch from "$lib/dp/components/formElements/Switch.svelte";
  import TabsRadioGroup from "$lib/dp/components/formElements/TabsRadioGroup.svelte";
  import Sheet from "$lib/dp/components/modal/Sheet.svelte";
  import { Collection } from "$lib/dp/enum/Collection.ts";
  import { toastController } from "$lib/dp/service/ToastController.svelte.ts";
  import type { Extension } from "$lib/dp/types/ExpandedResponse.js";
  import type { ExpandedEvent } from "$lib/dp/types/ExpandedResponse.ts";
  import type {
    LocationsResponse,
    TeamsResponse,
    UniformsetsResponse,
  } from "$lib/dp/types/pb-types.ts";
  import clsx from "clsx";
  import type { Snippet } from "svelte";

  interface Props {
    event: ExpandedEvent | null;
    clubID: string;
    teamID: string;
    mode: "teamEvent" | "clubEvent";
    triggerContent: Snippet;
    triggerVariant?:
      | "filled-primary"
      | "filled-secondary"
      | "tonal-primary"
      | "tonal-secondary"
      | "tonal-tertiary"
      | "tonal-surface";
    triggerSize?: "default" | "sm";
    triggerIcon?: boolean;
    triggerSpaced?: boolean;
  }

  const {
    event,
    clubID,
    teamID,
    mode,
    triggerContent,
    triggerVariant = "tonal-primary",
    triggerSize = "default",
    triggerIcon = false,
    triggerSpaced = false,
  }: Props = $props();

  let open = $state(false);

  function formFromProps(data: ExpandedEvent | null) {
    const ret = data ?? {
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
      team: mode === "teamEvent" ? teamID : "",
      club: mode === "clubEvent" ? clubID : "",
      additional_teams: [],
      expand: {},
    };
    //@ts-expect-error - ugliest of workarounds: to bind the MultiSelectCombobox to the array,
    // it needs to have a value even if no expand is sent from the backend
    if (ret.expand?.additional_teams === undefined) {
      if (!ret.expand) {
        ret.expand = {};
      }
      // @ts-expect-error
      ret.expand.additional_teams = [];
    }
    return ret;
  }

  let form: Extension<
    Partial<ExpandedEvent>,
    {
      starttime: string;
      endtime: string;
      meetingtime: string;
      type: string;
    }
  > = $derived.by(() => {
    const formData = $state(formFromProps(event));
    return formData;
  });

  let additionalTeams = $derived(form?.expand?.additional_teams ?? []);

  const attireOptions = $derived(
    client.collection(Collection.UniformSets).getFullList<UniformsetsResponse>({
      filter: `club = "${clubID}"`,
      requestKey: `uniformsets-${clubID}`,
    }),
  );

  const locationOptions = $derived(
    client.collection(Collection.Locations).getFullList<LocationsResponse>({
      filter: `club = "${clubID}"`,
      requestKey: `location-options-${clubID}`,
    }),
  );

  const additionalTeamOptions = $derived(
    client.collection(Collection.Teams).getFullList<TeamsResponse>({
      filter: `club = "${clubID}" && id != "${teamID}"`,
      requestKey: `team-options-${clubID}`,
    }),
  );

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ExpandedEvent | null = null;

    form.additional_teams = additionalTeams.map((team) => team.id);

    try {
      if (form.id) {
        result = await client
          .collection(Collection.Events)
          .update<ExpandedEvent>(form.id, form);
      } else {
        result = await client
          .collection(Collection.Events)
          .create<ExpandedEvent>(form);
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

<Sheet
  side="right"
  bind:open
  triggerClasses={clsx([
    "btn",
    "trigger-button",
    `trigger-variant-${triggerVariant}`,
    triggerSize === "sm" && "btn-sm",
    triggerIcon && "btn-icon",
    triggerSpaced && "trigger-spaced",
    triggerVariant === "filled-primary" && "preset-filled-primary-500",
    triggerVariant === "filled-secondary" && "preset-filled-secondary-500",
    triggerVariant === "tonal-primary" && "preset-tonal-primary border-primary",
    triggerVariant === "tonal-secondary" &&
      "preset-tonal-secondary border-secondary",
    triggerVariant === "tonal-tertiary" &&
      "preset-tonal-tertiary border-tertiary",
  ])}
>
  {#snippet triggerContent()}
    {@render triggerContent()}
  {/snippet}

  {#snippet title()}
    <header class="text-xl font-semibold">
      {#if form.id}
        <h2 class="h3">Edit Event "{form?.title}"</h2>
      {:else}
        <h2 class="h3">Create new Event</h2>
      {/if}
    </header>
  {/snippet}

  <form class="edit-form" onsubmit={submitForm}>
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

      <label class="label field-wide">
        <span>Additional Teams</span><br />

        {#await additionalTeamOptions then options}
          <MultiSelectCombobox
            itemName="Team"
            bind:selectedItems={additionalTeams}
            allItems={options}
            labelFunc={(item) => item.name}
            allowDeletionOfLastItem={true}
          />
        {/await}
      </label>

      <label class="label">
        <span>Start</span>
        <ISODatePicker bind:value={form.starttime} required={true} />
      </label>

      <label class="label">
        <span>Meeting</span>
        <ISODatePicker bind:value={form.meetingtime} />
      </label>

      <label class="label">
        <span>End</span>
        <ISODatePicker bind:value={form.endtime} />
      </label>

      <span></span>

      <label class="label field-wide">
        Description
        <textarea
          bind:value={form.desc}
          class="textarea"
          data-testid="event-form-textarea-desc"
          name="desc"></textarea>
      </label>

      <label class="label field-wide">
        Location
        <select bind:value={form.location} class="select">
          {#await locationOptions then options}
            <option value="">None</option>
            {#each options as option}
              <option value={option.id}
                >{option?.address_addon
                  ? option.address_addon
                  : "No additional name"}
                ({option.name}), {option.street}, {option.postal_code}
                {option.city}, {option.country}</option
              >
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
          <select
            class="select"
            bind:value={form.attire}
            data-testid="event-form-select-attire"
          >
            {#each options as option}
              <option value={option.id}>{option.name}</option>
            {/each}
          </select>
        </label>
      {/await}

      <Switch bind:checked={form.cancelled} name="cancelled">Cancelled</Switch>
    </div>

    <hr />

    <div class="submit-container">
      <button
        class="btn preset-tonal-primary border border-primary-500"
        data-testid="event-form-submit-button"
        type="submit"
      >
        Submit
      </button>
    </div>
  </form>
</Sheet>

<style>
  .submit-container {
    margin-block: calc(var(--spacing) * 3);
    display: flex;
    justify-content: center;
    align-items: center;
  }

  hr {
    margin-block: calc(var(--spacing) * 5);
  }
</style>
