<script lang="ts">
  import {invalidate} from "$app/navigation";
  import type {ExpandedEvent} from "$lib/model/ExpandedResponse";
  import {type LocationsResponse, type UniformsetsResponse} from "$lib/model/pb-types";
  import {client} from "$lib/pocketbase/index.svelte";
  import {Segment, Switch} from "@skeletonlabs/skeleton-svelte";
  import Flatpickr from "../utility/Flatpickr.svelte";
  import {DateTimeUtility} from "$lib/service/DateTimeUtility.js";
  //@ts-ignore
  import * as Sheet from "$lib/components/utility/sheet/index.js";
  import type {Snippet} from "svelte";
  import type {Extension} from "$lib/model/ExpandedResponse.js";
  import {toastController} from "$lib/service/ToastController.svelte.ts";

  interface Props {
    event: ExpandedEvent | null;
    clubID: string;
    teamID: string;
    triggerContent: Snippet;
    buttonClasses?: string;
  }

  let {event, clubID, teamID, triggerContent, buttonClasses = ""}: Props = $props();

  let open = $state(false);

  const form: Extension<Partial<ExpandedEvent>, { starttime: string, endtime: string, meetingtime: string }> = $state(
      event ?? {
        id: "",
        title: "",
        starttime: "",
        meetingtime: "",
        endtime: "",
        desc: "",
        location: "",
        type: undefined,
        attire: "",
        cancelled: false,
        bsm_id: 0,
        team: teamID,
      },
  );

  const attireOptions = client
      .collection("uniformsets")
      .getFullList<UniformsetsResponse>({
        filter: `club = "${clubID}"`,
        requestKey: `uniformsets-${clubID}`
      });

  const locationOptions = client.collection("locations").getFullList<LocationsResponse>({
    filter: `club = "${clubID}"`,
    requestKey: `location-options-${clubID}`
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ExpandedEvent | null = null;

    try {
      if (form.id) {
        result = await client
            .collection("events")
            .update<ExpandedEvent>(form.id, form);
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
  <Sheet.Trigger class={buttonClasses}>
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

    <form onsubmit={submitForm} class="mt-4 space-y-3">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-2 lg:gap-3 xl:gap-4">
        <input
                name="id"
                autocomplete="off"
                class="input"
                type="hidden"
                readonly
                bind:value={form.id}
        />

        <label class="label">
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
          BSM ID
          <input
                  name="bsm_id"
                  class="input"
                  readonly
                  type="text"
                  bind:value={form.bsm_id}
          />
        </label>

        <label class="label">
          Start
          <Flatpickr
                  bind:value={form.starttime}
                  options={Object.assign(DateTimeUtility.datePickerOptions, {
                      static: true, // render the picker as a child element to the form to work in a sheet portal context
                  })}
          />
        </label>

        <label class="label">
          Meeting
          <Flatpickr
                  bind:value={form.meetingtime}
                  options={Object.assign(DateTimeUtility.datePickerOptions, {
                      static: true,
                  })}
          />
        </label>

        <label class="label">
          End
          <Flatpickr
                  bind:value={form.endtime}
                  options={Object.assign(DateTimeUtility.datePickerOptions, {
                      static: true,
                  })}
          />
        </label>

        <span></span>

        <label class="label md:col-span-2">
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

        <label class="label flex flex-col gap-1 md:col-span-2">
          Type
          <Segment name="type" value={form.type} onValueChange={(e) => (form.type = e.value ?? "game")}>
            <Segment.Item value={"game"} classes="flex-grow">
              Game
            </Segment.Item>
            <Segment.Item value={"practice"} classes="flex-grow">
              Practice
            </Segment.Item>
            <Segment.Item value={"misc"} classes="flex-grow">
              Other
            </Segment.Item>
          </Segment>
        </label>

        {#await attireOptions then options}
          <label class="label md:col-span-2">
            Uniform Set
            <select class="select" bind:value={form.attire}>
              {#each options as option}
                <option value={option.id}>{option.name}</option>
              {/each}
            </select>
          </label>
        {/await}

        <Switch
                name="cancelled"
                onCheckedChange={(e) => (form.cancelled = e.checked)}
        >
          Cancelled
        </Switch>
      </div>

      <hr class="my-5!"/>

      <div class="flex justify-center gap-3">
        <button type="submit" class="mt-2 btn preset-tonal-primary border border-primary-500">
          Submit
        </button>
      </div>
    </form>
  </Sheet.Content>
</Sheet.Root>
