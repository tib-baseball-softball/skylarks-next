<script lang="ts">
  import type {CustomAuthModel, ExpandedEvent} from "$lib/model/ExpandedResponse";
  import {fade} from "svelte/transition";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import type {EventsUpdate} from "$lib/model/pb-types";
  import IndividualParticipationEditButton
    from "$lib/components/diamondplanner/event/IndividualParticipationEditButton.svelte";
  import {Check, CircleHelp, Trash, X} from "lucide-svelte";

  interface Props {
    event: ExpandedEvent;
  }

  let {event}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const displayedGuestPlayers = $derived(event.guests.split(","));
  const isAdmin = $derived((event.expand?.team?.admins.includes(authRecord.id) || event?.expand?.team?.expand?.club?.admins.includes(authRecord.id)) ?? false);

  async function removeGuestPlayer(playerToRemove: string) {
    const newGuestPlayerList = displayedGuestPlayers.filter((player) => player !== playerToRemove);
    await client.collection("events").update<EventsUpdate>(event.id, {
      guests: newGuestPlayerList.join()
    });
  }
</script>

<h2 class="h2">Participants</h2>
<section class="grid grid-cols-1 md:grid-cols-3 gap-2 lg:gap-3">
  <article class="card variant-soft-success grow">
    <header class="card-header flex items-center gap-2">
      <span><Check/></span>
      <h3 class="h4">In</h3>
    </header>

    <section class="p-4 flex flex-wrap gap-2">
      {#key event.participations.in}
        {#each event.participations.in as inResponse}
          <div in:fade|global={{delay: 200}}>
            <IndividualParticipationEditButton
                    participation={inResponse}
                    {isAdmin}
                    classes="chip variant-ghost-success"
            />
          </div>
        {/each}

        {#each displayedGuestPlayers as guestPlayer}
          {#if guestPlayer /* can be empty string */}
            <div in:fade|global={{delay: 200}}>
              <button
                      aria-label="guest player name, click removes the player from the event"
                      class="chip variant-soft dark:variant-soft-success"
                      onclick={() => removeGuestPlayer(guestPlayer)}
              >
                {guestPlayer}
                {#if isAdmin}
                  <Trash class="ms-1" size="10"/>
                {/if}
              </button>
            </div>
          {/if}
        {/each}
      {/key}
    </section>
  </article>

  <article class="card variant-soft-warning grow">
    <header class="card-header flex items-center gap-2">
      <span><CircleHelp/></span>
      <h3 class="h4">Maybe</h3>
    </header>

    <section class="p-4 flex flex-wrap gap-2">
      {#key event.participations.maybe}
        {#each event.participations.maybe as maybeResponse}
          <div in:fade|global={{delay: 200}}>
            <IndividualParticipationEditButton
                    participation={maybeResponse}
                    {isAdmin}
                    classes="chip variant-ghost-warning"
            />
          </div>
        {/each}
      {/key}
    </section>
  </article>

  <article class="card variant-soft-error grow">
    <header class="card-header flex items-center gap-2">
      <span><X/></span>
      <h3 class="h4">Out</h3>
    </header>

    <section class="p-4 flex flex-wrap gap-2">
      {#key event.participations.out}
        {#each event.participations.out as outResponse}
          <div in:fade|global={{delay: 200}}>
            <IndividualParticipationEditButton
                    participation={outResponse}
                    {isAdmin}
                    classes="chip variant-ghost-error"
            />
          </div>
        {/each}
      {/key}
    </section>
  </article>
</section>

<style lang="postcss">
    h2, h3 {
        @apply my-3;
    }

    .chip {
        cursor: default;
    }

    .chip.variant-soft {
        cursor: pointer;
    }
</style>