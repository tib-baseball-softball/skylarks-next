<script lang="ts">
  import {Check, CircleQuestionMark, Ghost, Trash, X} from "lucide-svelte";
  import {fade} from "svelte/transition";
  import ExternalParticipationWrapper from "$lib/dp/components/event/ExternalParticipationWrapper.svelte";
  import IndividualParticipationEditButton from "$lib/dp/components/event/IndividualParticipationEditButton.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {
    CustomAuthModel,
    ExpandedEvent,
    ExpandedParticipation,
    ParticipationType
  } from "$lib/dp/types/ExpandedResponse.ts";
  import type {EventsUpdate} from "$lib/dp/types/pb-types.ts";
  import {Collection} from "$lib/dp/enum/Collection.ts";
  import {sendParticipationData} from "$lib/dp/utility/sendParticipationData.ts";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import {invalidate} from "$app/navigation";
  import {participationDTOToExpandedParticipation} from "$lib/dp/types/UserParticipationDTO.ts";

  interface Props {
    event: ExpandedEvent;
  }

  const {event}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const displayedGuestPlayers = $derived(event.guests.split(","));
  const isAdmin = $derived(
    (event.expand?.team?.admins.includes(authRecord.id) ||
      event?.expand?.team?.expand?.club?.admins.includes(authRecord.id)) ??
    false
  );

  async function removeGuestPlayer(playerToRemove: string) {
    const newGuestPlayerList = displayedGuestPlayers.filter((player) => player !== playerToRemove);
    await client.collection(Collection.Events).update<EventsUpdate>(event.id, {
      guests: newGuestPlayerList.join(),
    });
  }

  let currentDraggedParticipation: ExpandedParticipation | null = $state(null);

  function ondragover(event: DragEvent) {
    if (event.dataTransfer?.types?.includes("participation")) {
      event.preventDefault();
    }
  }

  function ondragstart(event: DragEvent, participation: ExpandedParticipation) {
    if (!event.dataTransfer) {
      return;
    }

    event.dataTransfer.effectAllowed = "move";
    event.dataTransfer.setData("participation", participation.id);
    currentDraggedParticipation = participation;
  }

  async function ondrop(event: DragEvent, type: ParticipationType) {
    event.preventDefault();
    console.log('Drop event triggered for participation:', event.dataTransfer?.getData("participation"));

    if (currentDraggedParticipation === null || currentDraggedParticipation.state === type) {
      return;
    }

    currentDraggedParticipation.state = type;

    try {
      await sendParticipationData(currentDraggedParticipation);
    } catch (e) {
      toastController.trigger({
        message: 'Failed to send participation data',
        background: 'preset-filled-error-500',
      });
    }
    currentDraggedParticipation = null;
    await invalidate("event:single");
  }
</script>

<h2 class="participants-title">Participants</h2>
<section class="event-drag-drop-section participants-grid">
  <article
    class="drop-target card participant-card preset-tonal-success"
    ondrop={(e) => ondrop(e, "in")}
    {ondragover}
  >
    <header class="participation-header">
      <span><Check/></span>
      <h3 class="h4">In</h3>
    </header>

    <section>
      <ul class="participation-content">
        {#key event.participations.in}
          {#each event.participations.in as inResponse (inResponse.id)}
            <li draggable="true" in:fade|global={{delay: 200}} ondragstart={(event) => ondragstart(event, inResponse)}>
              <IndividualParticipationEditButton
                participation={inResponse}
                {isAdmin}
                classes="chip preset-tonal-success border-success"
              />
            </li>
          {/each}

          {#each displayedGuestPlayers as guestPlayer}
            {#if guestPlayer /* can be an empty string */}
              <li in:fade|global={{delay: 200}}>
                <button
                  aria-label="guest player name, click removes the player from the event"
                  class="chip guest-chip preset-tonal border-surface"
                  onclick={() => removeGuestPlayer(guestPlayer)}
                >
                  {guestPlayer}
                  {#if isAdmin}
                    <Trash size="10"/>
                  {/if}
                </button>
              </li>
            {/if}
          {/each}
        {/key}
      </ul>
    </section>
  </article>

  <article
    class="drop-target card participant-card preset-tonal-warning"
    ondrop={(e) => ondrop(e, "maybe")}
    {ondragover}
  >
    <header class="participation-header">
      <span><CircleQuestionMark/></span>
      <h3 class="h4">Maybe</h3>
    </header>

    <section>
      <ul class="participation-content">
        {#key event.participations.maybe}
          {#each event.participations.maybe as maybeResponse (maybeResponse.id)}
            <li draggable="true" in:fade|global={{delay: 200}}
                ondragstart={(event) => ondragstart(event, maybeResponse)}>
              <IndividualParticipationEditButton
                participation={maybeResponse}
                {isAdmin}
                classes="chip preset-tonal-warning border-warning"
              />
            </li>
          {/each}
        {/key}
      </ul>
    </section>
  </article>

  <article
    class="drop-target card participant-card preset-tonal-error"
    ondrop={(e) => ondrop(e, "out")}
    {ondragover}
  >
    <header class="participation-header">
      <span><X/></span>
      <h3 class="h4">Out</h3>
    </header>

    <section>
      <ul class="participation-content">
        {#key event.participations.out}
          {#each event.participations.out as outResponse (outResponse.id)}
            <li draggable="true" in:fade|global={{delay: 200}} ondragstart={(event) => ondragstart(event, outResponse)}>
              <IndividualParticipationEditButton
                participation={outResponse}
                {isAdmin}
                classes="chip preset-tonal-error border-error"
              />
            </li>
          {/each}
        {/key}
      </ul>
    </section>
  </article>
</section>

<section class="no-response-section">
  <article class="card participant-card preset-outlined-card">
    <header class="participation-header">
      <span><Ghost/></span>
      <h3 class="h4">No Response</h3>
    </header>

    <section>
      <ul class="participation-content">
        {#key event.participations.unspecified}
          {#each event.participations.unspecified as ghostResponse}
            <li draggable="true" in:fade|global={{delay: 200}}
                ondragstart={(e) => ondragstart(e, participationDTOToExpandedParticipation(ghostResponse, event.id))}>
              <ExternalParticipationWrapper
                dto={ghostResponse}
                eventID={event.id}
                {isAdmin}
                classes="chip preset-outlined"
              />
            </li>
          {/each}
        {/key}
      </ul>
    </section>
  </article>
</section>

<style>
  .participants-title {
    margin-block: calc(var(--spacing) * 3);
    font-size: var(--text-2xl);
    font-weight: bold;
  }

  h3 {
    margin-block: calc(var(--spacing) * 3);
  }

  .participants-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 2);

    @media (min-width: 48rem) {
      grid-template-columns: repeat(3, minmax(0, 1fr));
    }

    @media (min-width: 64rem) {
      gap: calc(var(--spacing) * 3);
    }
  }

  .participant-card {
    flex-grow: 1;
  }

  .participation-header {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 2);
    padding-inline-start: calc(var(--spacing) * 4);
  }

  .participation-content {
    padding: calc(var(--spacing) * 4);
    display: flex;
    flex-wrap: wrap;
    gap: calc(var(--spacing) * 2);
  }

  .no-response-section {
    margin-top: calc(var(--spacing) * 6);
  }

  .chip {
    cursor: default;
  }

  .chip.guest-chip {
    cursor: pointer;
    gap: calc(var(--spacing) * 1);
  }

  .border-surface {
    border: 1px solid var(--color-surface-900);

    @media (prefers-color-scheme: dark) {
      border-color: var(--color-surface-100);
    }
  }
</style>
