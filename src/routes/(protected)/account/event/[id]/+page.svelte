<script lang="ts">
  import EventCoreInfo from "$lib/components/diamondplanner/event/EventCoreInfo.svelte";
  import EventParticipationSection from "$lib/components/diamondplanner/event/EventParticipationSection.svelte";
  import EventTypeBadge from "$lib/components/diamondplanner/event/EventTypeBadge.svelte";
  import MatchTeaserCard from "$lib/components/match/MatchTeaserCard.svelte";
  import type { ParticipationsResponse } from "$lib/model/pb-types.js";
  import { CheckOutline, QuestionCircleOutline } from "flowbite-svelte-icons";
  import CloseOutline from "flowbite-svelte-icons/CloseOutline.svelte";

  let { data } = $props();

  const event = $derived(data.event);

  // TODO: get this data via backend hook

  let participations: ParticipationsResponse[] = $derived(
    $event?.expand?.participations_via_event,
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
</script>

<div class="event-container">
  <div class="flex items-center gap-3">
    <h1 class="h1">{$event.title}</h1>
    <div>
      <EventTypeBadge type={$event.type} />
    </div>
  </div>

  <article class="!mb-8">
    <section>
      <p>{$event.desc}</p>
    </section>
  </article>

  <EventCoreInfo event={$event} classes={"space-y-6"} />

  <hr class="!my-8" />

  <div class="flex justify-between items-center">
    <h2 class="h4">My Participation</h2>
    <EventParticipationSection event={$event} />
  </div>

  <hr class="my-8" />

  <h2 class="h2">Participants</h2>
  <section class="grid grid-cols-1 md:grid-cols-3 gap-2 lg:gap-3">
    <div class="card variant-soft-success flex-grow">
      <header class="card-header flex items-center gap-2">
        <span><CheckOutline size="lg" /></span>
        <h3 class="h4">In</h3>
      </header>

      <section class="p-4 flex flex-wrap gap-1 lg:gap-2">
        {#each participationsIn as inResponse}
          <div class="chip variant-ghost-success">
            {inResponse?.expand?.user.first_name}
          </div>
        {/each}
      </section>
    </div>

    <div class="card variant-soft-warning flex-grow">
      <header class="card-header flex items-center gap-2">
        <span><QuestionCircleOutline size="lg" /></span>
        <h3 class="h4">Maybe</h3>
      </header>
      <section class="p-4">
        {#each participationsMaybe as maybeResponse}
          <div class="chip variant-ghost-warning">
            {maybeResponse?.expand?.user.first_name}
          </div>
        {/each}
      </section>
    </div>

    <div class="card variant-soft-error flex-grow">
      <header class="card-header flex items-center gap-2">
        <span><CloseOutline size="lg" /></span>
        <h3 class="h4">Out</h3>
      </header>
      <section class="p-4">
        {#each participationsOut as outResponse}
          <div class="chip variant-ghost-error">
            {outResponse?.expand?.user.first_name}
          </div>
        {/each}
      </section>
    </div>
  </section>

  {#if $event.match_json}
    <section>
      <h2 class="h2 mb-3">Game Data</h2>
      <div class="grid grid-cols-1 md:grid-cols-2">
        <MatchTeaserCard match={$event.match_json} />
      </div>
    </section>
  {/if}

  <hr class="my-2" />

  <h2 class="h2">Admin Info</h2>
  <p>BSM-ID: {$event.bsm_id}</p>
  <p>Created: {$event.created}</p>
  <p>Updated: {$event.updated}</p>
</div>

<style lang="postcss">
  .event-container {
    @apply space-y-4 lg:space-y-6 xl:space-y-7;
  }
</style>
