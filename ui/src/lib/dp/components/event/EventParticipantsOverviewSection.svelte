<script lang="ts">
  import {Check, CircleQuestionMark, Ghost, Trash, X} from "lucide-svelte";
  import {fade} from "svelte/transition";
  import ExternalParticipationWrapper from "$lib/dp/components/event/ExternalParticipationWrapper.svelte";
  import IndividualParticipationEditButton from "$lib/dp/components/event/IndividualParticipationEditButton.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";
  import type {EventsUpdate} from "$lib/dp/types/pb-types.ts";

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
    await client.collection("events").update<EventsUpdate>(event.id, {
      guests: newGuestPlayerList.join(),
    });
  }
</script>

<h2 class="participants-title">Participants</h2>
<section class="participants-grid">
  <article class="card participant-card preset-tonal-success">
    <header class="participation-header">
      <span><Check/></span>
      <h3 class="h4">In</h3>
    </header>

    <section class="participation-content">
      {#key event.participations.in}
        {#each event.participations.in as inResponse}
          <div in:fade|global={{delay: 200}}>
            <IndividualParticipationEditButton
              participation={inResponse}
              {isAdmin}
              classes="chip preset-tonal-success border-success"
            />
          </div>
        {/each}

        {#each displayedGuestPlayers as guestPlayer}
          {#if guestPlayer /* can be an empty string */}
            <div in:fade|global={{delay: 200}}>
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
            </div>
          {/if}
        {/each}
      {/key}
    </section>
  </article>

  <article class="card participant-card preset-tonal-warning">
    <header class="participation-header">
      <span><CircleQuestionMark/></span>
      <h3 class="h4">Maybe</h3>
    </header>

    <section class="participation-content">
      {#key event.participations.maybe}
        {#each event.participations.maybe as maybeResponse}
          <div in:fade|global={{delay: 200}}>
            <IndividualParticipationEditButton
              participation={maybeResponse}
              {isAdmin}
              classes="chip preset-tonal-warning border-warning"
            />
          </div>
        {/each}
      {/key}
    </section>
  </article>

  <article class="card participant-card preset-tonal-error">
    <header class="participation-header">
      <span><X/></span>
      <h3 class="h4">Out</h3>
    </header>

    <section class="participation-content">
      {#key event.participations.out}
        {#each event.participations.out as outResponse}
          <div in:fade|global={{delay: 200}}>
            <IndividualParticipationEditButton
              participation={outResponse}
              {isAdmin}
              classes="chip preset-tonal-error border-error"
            />
          </div>
        {/each}
      {/key}
    </section>
  </article>
</section>

<section class="no-response-section">
  <article class="card participant-card preset-outlined">
    <header class="participation-header">
      <span><Ghost/></span>
      <h3 class="h4">No Response</h3>
    </header>

    <section class="participation-content">
      {#key event.participations.unspecified}
        {#each event.participations.unspecified as ghostResponse}
          <div in:fade|global={{delay: 200}}>
            <ExternalParticipationWrapper
              dto={ghostResponse}
              eventID={event.id}
              {isAdmin}
              classes="chip preset-outlined"
            />
          </div>
        {/each}
      {/key}
    </section>
  </article>
</section>

<style lang="postcss">
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

  .border-success {
    border: 1px solid var(--color-success-500);
  }

  .border-warning {
    border: 1px solid var(--color-warning-500);
  }

  .border-error {
    border: 1px solid var(--color-error-500);
  }

  .border-surface {
    border: 1px solid var(--color-surface-900);

    @media (prefers-color-scheme: dark) {
      border-color: var(--color-surface-100);
    }
  }
</style>
