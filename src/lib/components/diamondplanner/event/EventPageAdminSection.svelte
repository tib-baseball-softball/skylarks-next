<script lang="ts">
  import type {ExpandedEvent} from "$lib/model/ExpandedResponse";
  import {type DrawerSettings, getDrawerStore} from "@skeletonlabs/skeleton";
  import {client} from "$lib/pocketbase/index.svelte";
  import type {EventsUpdate} from "$lib/model/pb-types";
  import {goto} from "$app/navigation";
  import DeleteButton from "$lib/components/utility/DeleteButton.svelte";
  import {CalendarArrowDown, CalendarPlus, Edit, Info} from "lucide-svelte";

  interface Props {
    event: ExpandedEvent;
  }

  let {event}: Props = $props();

  const drawerStore = getDrawerStore();

  const settings: DrawerSettings = $derived({
    id: "event-form",
    position: "right",
    width: "w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%]",
    meta: {
      event: event,
      club: event?.expand?.team?.club,
      team: event?.expand?.team?.id,
    },
  });

  let guestPlayerForm = $state({
    name: ""
  });

  async function submitNewGuestPlayer(e: SubmitEvent) {
    e.preventDefault();

    await client.collection("events").update<EventsUpdate>(event.id, {
      guests: event.guests + "," + guestPlayerForm.name
    });
    guestPlayerForm.name = "";
  }

  async function deleteEvent(id: string) {
    await client.collection("events").delete(id);
    await goto(`/account/team/${event.team}`);
  }
</script>

<hr class="mt-6"/>

<h2 class="h2 my-4">Admin Section</h2>

<div
        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2 lg:gap-3"
>
  <div class="card variant-ringed-surface">
    <header class="card-header">
      <h3 class="h4 font-semibold">Event Details</h3>
    </header>

    <section class="p-4 space-y-2">
      <div class="flex items-center gap-3">
        <Info/>
        <div>
          <p class="text-sm font-light">BSM event ID</p>
          <p>{event.bsm_id}</p>
        </div>
      </div>

      <div class="flex items-center gap-3">
        <CalendarPlus/>
        <div>
          <p class="text-sm font-light">Created</p>
          <p>{new Date(event.created).toLocaleString()}</p>
        </div>
      </div>

      <div class="flex items-center gap-3">
        <CalendarArrowDown/>
        <div>
          <p class="text-sm font-light">Last Updated</p>
          <p>{new Date(event.updated).toLocaleString()}</p>
        </div>
      </div>
    </section>
  </div>

  <div class="card variant-ringed-surface">
    <header class="card-header">
      <h3 class="h4 font-semibold">Add Guest player</h3>
    </header>

    <section class="p-4 space-y-2">
      <form
              onsubmit={submitNewGuestPlayer}
      >
        <label class="label">
          Name
          <input
                  type="text"
                  class="input mt-1"
                  placeholder="Enter the guest player's name"
                  bind:value={guestPlayerForm.name}
          />
        </label>
        <button type="submit" class="mt-3 btn variant-soft">Submit
        </button
        >
      </form>
    </section>
  </div>

  <div class="card variant-ringed-surface">
    <header class="card-header">
      <h3 class="h4 font-semibold">Actions</h3>
    </header>

    <section class="p-4 space-y-3">
      <div class="flex flex-col gap-2 lg:gap-3">
        <button
                class="btn variant-ghost-surface"
                onclick={() => drawerStore.open(settings)}
        >
          <Edit/>
          <span>Edit Event</span>
        </button>

        <DeleteButton
                id={event.id}
                modelName="Event"
                action={deleteEvent}
                classes="variant-ghost-error"
                buttonText="Delete Event"
        />
      </div>
    </section>
  </div>
</div>