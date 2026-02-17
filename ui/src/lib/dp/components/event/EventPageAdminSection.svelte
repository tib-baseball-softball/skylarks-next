<script lang="ts">
  import {CalendarArrowDown, CalendarPlus, Edit, Info} from "lucide-svelte";
  import {goto} from "$app/navigation";
  import EventForm from "$lib/dp/components/forms/EventForm.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import {client} from "$lib/dp/client.svelte.js";
  import type {ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";
  import type {EventsUpdate} from "$lib/dp/types/pb-types.ts";

  interface Props {
    event: ExpandedEvent;
  }

  const {event}: Props = $props();

  const guestPlayerForm = $state({
    name: "",
  });

  async function submitNewGuestPlayer(e: SubmitEvent) {
    e.preventDefault();

    await client.collection("events").update<EventsUpdate>(event.id, {
      guests:
          event.guests.length === 0 ? guestPlayerForm.name : event.guests + "," + guestPlayerForm.name,
    });
    guestPlayerForm.name = "";
  }

  async function deleteEvent(id: string) {
    await client.collection("events").delete(id);
    await goto(`/account/team/${event.team}`);
  }
</script>

<hr class="admin-divider"/>

<h2 class="admin-title">Admin Section</h2>

<div class="admin-grid">
  <div class="card admin-card preset-outlined-surface-500">
    <header class="card-header">
      <h3 class="card-title">Event Details</h3>
    </header>

    <section class="card-section info-section">
      <div class="info-row">
        <span class="icon"><Info/></span>
        <div>
          <p class="info-label">BSM event ID</p>
          <p>{event.bsm_id}</p>
        </div>
      </div>

      <div class="info-row">
        <span class="icon"><CalendarPlus/></span>
        <div>
          <p class="info-label">Created</p>
          <p>{new Date(event.created).toLocaleString()}</p>
        </div>
      </div>

      <div class="info-row">
        <span class="icon"><CalendarArrowDown/></span>
        <div>
          <p class="info-label">Last Updated</p>
          <p>{new Date(event.updated).toLocaleString()}</p>
        </div>
      </div>
    </section>
  </div>

  <div class="card admin-card preset-outlined-surface-500">
    <header class="card-header">
      <h3 class="card-title">Add Guest player</h3>
    </header>

    <section class="card-section">
      <form class="guest-form" onsubmit={submitNewGuestPlayer}>
        <label class="label">
          Name
          <input
                  bind:value={guestPlayerForm.name}
                  class="input name-input"
                  placeholder="Enter the guest player's name"
                  type="text"
          />
        </label>
        <button class="btn submit-btn" type="submit">Submit</button>
      </form>
    </section>
  </div>

  <div class="card admin-card preset-outlined-surface-500">
    <header class="card-header">
      <h3 class="card-title">Actions</h3>
    </header>

    <section class="card-section actions-section">
      <div class="actions-container">

        <EventForm
                buttonClasses="btn preset-tonal-surface border-surface"
                clubID={event?.expand?.team?.club ?? ""}
                event={event}
                teamID={event.team}
        >
          {#snippet triggerContent()}
            <Edit/>
            <span>Edit Event</span>
          {/snippet}
        </EventForm>

        <DeleteButton
                action={deleteEvent}
                buttonText="Delete Event"
                classes="preset-tonal-error border-error"
                id={event.id}
                modelName="Event"
        />
      </div>
    </section>
  </div>
</div>

<style lang="postcss">
    .admin-divider {
        margin-top: calc(var(--spacing) * 6);
    }

    .admin-title {
        font-size: var(--text-2xl);
        font-weight: bold;
        margin-block: calc(var(--spacing) * 4);
    }

    .admin-grid {
        display: grid;
        grid-template-columns: 1fr;
        gap: calc(var(--spacing) * 2);

        @media (min-width: 48rem) {
            grid-template-columns: repeat(2, minmax(0, 1fr));
        }

        @media (min-width: 64rem) {
            grid-template-columns: repeat(3, minmax(0, 1fr));
            gap: calc(var(--spacing) * 3);
        }
    }

    .card-title {
        font-size: var(--text-lg);
        font-weight: 600;
    }

    .card-section {
        padding: calc(var(--spacing) * 4);
        display: flex;
        flex-direction: column;
    }

    .info-section {
        gap: calc(var(--spacing) * 2);
    }

    .info-row {
        display: flex;
        align-items: center;
        gap: calc(var(--spacing) * 3);
    }

    .info-label {
        font-size: var(--text-sm);
        font-weight: 300;
    }

    .guest-form {
        display: flex;
        flex-direction: column;
        gap: calc(var(--spacing) * 2);

        .name-input {
            margin-top: calc(var(--spacing) * 1);
        }

        .submit-btn {
            margin-top: calc(var(--spacing) * 3);
            border: 1px solid var(--color-surface-500);
            background: var(--color-surface-200);
            
            &:hover {
                background: var(--color-surface-300);
            }
        }
    }

    .actions-section {
        gap: calc(var(--spacing) * 3);
    }

    .actions-container {
        display: flex;
        flex-direction: column;
        gap: calc(var(--spacing) * 2);

        @media (min-width: 64rem) {
            gap: calc(var(--spacing) * 3);
        }
    }

    :global(.border-surface) {
        border: 1px solid var(--color-surface-500) !important;
    }

    :global(.border-error) {
        border: 1px solid var(--color-error-500) !important;
    }
</style>