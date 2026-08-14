<script lang="ts">
  import { Check, CircleQuestionMark, X } from "@lucide/svelte";
  import { invalidate } from "$app/navigation";
  import { authSettings } from "$lib/dp/client.svelte.js";
  import type { EventParticipationState } from "$lib/dp/types/EventParticipationState.ts";
  import type {
    CustomAuthModel,
    ExpandedEvent,
  } from "$lib/dp/types/ExpandedResponse.ts";
  import type { ParticipationsCreate } from "$lib/dp/types/pb-types.ts";
  import { sendParticipationData } from "$lib/dp/utility/sendParticipationData.ts";

  interface props {
    event: ExpandedEvent;
    canParticipate: boolean;
    growChips?: boolean;
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const { event, canParticipate, growChips = false }: props = $props();

  const userParticipation: ParticipationsCreate = $derived(
    event.userParticipation ?? {
      id: "",
      user: authRecord?.id,
      event: event.id,
      state: "",
      comment: "",
    },
  );

  // JS: splitting an empty string by comma returns length `1`
  const guestCount = $derived(
    event.guests === "" ? 0 : event.guests.split(",").length,
  );

  let submitting = $state(false);

  async function updateParticipationStatus(
    state: EventParticipationState,
  ): Promise<void> {
    if (submitting) {
      return;
    }
    submitting = true;
    try {
      userParticipation.state = state;
      await sendParticipationData(userParticipation);
      await invalidate("event:list");
    } finally {
      submitting = false;
    }
  }
</script>

<section class="participation-section" data-grow={growChips}>
  <button
    aria-busy={submitting}
    class="chip preset-outlined-success-500 participation-button"
    class:preset-filled-success-500={userParticipation.state === "in"}
    disabled={submitting || !canParticipate}
    onclick={() => updateParticipationStatus("in")}
    title="set your participation to 'in' for this event"
  >
    <span aria-hidden="true"><Check size="14" /></span>
    <span class="sr-only"
      >set your participation to "in" for "{event.title}"</span
    >
    <span aria-hidden="true"
      >{(event.participations.in.length ?? 0) + guestCount}</span
    >
  </button>

  <button
    aria-busy={submitting}
    class="chip preset-outlined-warning-500 participation-button"
    class:preset-filled-warning-500={userParticipation.state === "maybe"}
    disabled={submitting || !canParticipate}
    onclick={() => updateParticipationStatus("maybe")}
    title="set your participation to 'maybe' for this event"
  >
    <span aria-hidden="true"><CircleQuestionMark size="14" /></span>
    <span class="sr-only"
      >set your participation to "maybe" for "{event.title}"</span
    >
    <span aria-hidden="true">{event.participations.maybe.length ?? 0}</span>
  </button>

  <button
    aria-busy={submitting}
    class="chip preset-outlined-error-500 participation-button"
    class:preset-filled-error-500={userParticipation.state === "out"}
    disabled={submitting || !canParticipate}
    onclick={() => updateParticipationStatus("out")}
    title="set your participation to 'out' for this event"
  >
    <span aria-hidden="true"><X size="14" /></span>
    <span class="sr-only"
      >set your participation to "out" for "{event.title}"</span
    >
    <span aria-hidden="true">{event.participations.out.length ?? 0} </span>
  </button>
</section>

<style>
  .participation-section {
    display: flex;
    justify-content: flex-end;
    align-items: flex-end;
    gap: calc(var(--spacing) * 2);
    flex-wrap: wrap;
  }

  [data-grow="true"] {
    .chip {
      flex-grow: 1;
    }
  }
</style>
