<script lang="ts">
  import {CheckOutline, QuestionCircleOutline, TrashBinOutline} from "flowbite-svelte-icons";
  import CloseOutline from "flowbite-svelte-icons/CloseOutline.svelte";
  import type {CustomAuthModel, ExpandedEvent} from "$lib/model/ExpandedResponse";
  import {fade} from "svelte/transition";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import type {EventsUpdate} from "$lib/model/pb-types";

  interface Props {
    event: ExpandedEvent
  }

  let {event}: Props = $props()

  const authRecord = authSettings.record as CustomAuthModel

  // TODO: get this data via backend hook
  let participations = $derived(
      event?.expand?.participations_via_event ?? [],
  );

  let participationsIn = $derived(
      participations?.filter((participation) => {
        return participation.state === "in";
      }),
  );
  let participationsMaybe = $derived(
      participations?.filter((participation) => {
        return participation.state === "maybe";
      }),
  );
  let participationsOut = $derived(
      participations?.filter((participation) => {
        return participation.state === "out";
      }),
  );

  const displayedGuestPlayers = $derived(event.guests.split(","))
  const isAdmin = $derived(event.expand?.team?.admins.includes(authRecord.id) || event?.expand?.team?.expand?.club?.admins.includes(authRecord.id))

  async function removeGuestPlayer(playerToRemove: string) {
    const newGuestPlayerList = displayedGuestPlayers.filter((player) => player !== playerToRemove)
    await client.collection("events").update<EventsUpdate>(event.id, {
      guests: newGuestPlayerList.join()
    })
  }
</script>

<h2 class="h2">Participants</h2>
<section class="grid grid-cols-1 md:grid-cols-3 gap-2 lg:gap-3">
    <div class="card variant-ghost-success flex-grow">
        <header class="card-header flex items-center gap-2">
            <span><CheckOutline size="lg"/></span>
            <h3 class="h4">In</h3>
        </header>

        <section class="p-4 flex flex-wrap gap-1 lg:gap-2">
            {#key participations}
                {#each participationsIn as inResponse}
                    <div in:fade|global={{delay: 200}}>
                        <div class="chip variant-ghost-success">
                            {inResponse?.expand?.user?.first_name}
                        </div>
                    </div>
                {/each}

                {#each displayedGuestPlayers as guestPlayer}
                    {#if guestPlayer /* can be empty string */}
                        <div in:fade|global={{delay: 200}}>
                            <button
                                    aria-label="guest player name, click removes the player from the event"
                                    class="chip variant-soft"
                                    onclick={() => removeGuestPlayer(guestPlayer)}
                            >
                                {guestPlayer}
                                {#if isAdmin}
                                    <TrashBinOutline class="ms-0.5" size="xs"/>
                                {/if}
                            </button>
                        </div>
                    {/if}
                {/each}
            {/key}
        </section>
    </div>

    <div class="card variant-ghost-warning flex-grow">
        <header class="card-header flex items-center gap-2">
            <span><QuestionCircleOutline size="lg"/></span>
            <h3 class="h4">Maybe</h3>
        </header>
        <section class="p-4 flex flex-wrap gap-1 lg:gap-2">
            {#key participations}
                {#each participationsMaybe as maybeResponse}
                    <div in:fade|global={{delay: 200}}>
                        <div class="chip variant-ghost-warning">
                            {maybeResponse?.expand?.user?.first_name}
                        </div>
                    </div>
                {/each}
            {/key}
        </section>
    </div>

    <div class="card variant-ghost-error flex-grow">
        <header class="card-header flex items-center gap-2">
            <span><CloseOutline size="lg"/></span>
            <h3 class="h4">Out</h3>
        </header>
        <section class="p-4 flex flex-wrap gap-1 lg:gap-2">
            {#key participations}
                {#each participationsOut as outResponse}
                    <div in:fade|global={{delay: 200}}>
                        <div class="chip variant-ghost-error">
                            {outResponse?.expand?.user?.first_name}
                        </div>
                    </div>
                {/each}
            {/key}
        </section>
    </div>
</section>

{#if isAdmin}
    <button type="button" class="my-3 btn variant-ghost">
        Edit
    </button>
{/if}

<style lang="postcss">
    h2, h3 {
        @apply my-3
    }

    .chip {
        cursor: default;
    }

    .chip.variant-soft {
        cursor: pointer;
    }
</style>