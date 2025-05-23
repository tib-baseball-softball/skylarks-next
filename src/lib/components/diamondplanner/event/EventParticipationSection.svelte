<script lang="ts">
  import type {ParticipationsCreate} from "$lib/model/pb-types";
  import type {EventParticipationState} from "$lib/types/EventParticipationState";
  import {sendParticipationData} from "$lib/functions/sendParticipationData";
  import type {CustomAuthModel, ExpandedEvent} from "$lib/model/ExpandedResponse";
  import {invalidate} from "$app/navigation";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import {Check, CircleHelp, X} from "lucide-svelte";

  interface props {
    event: ExpandedEvent;
    chipClasses?: string;
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const {event, chipClasses = ""}: props = $props();

  let userParticipation: ParticipationsCreate = $derived(event.userParticipation ?? {
    id: "",
    user: authRecord?.id,
    event: event.id,
    state: "",
    comment: "",
  });

  // JS: splitting an empty string by comma returns length `1`
  const guestCount = $derived(event.guests === "" ? 0 : event.guests.split(",").length)

  async function updateParticipationStatus(
      state: EventParticipationState,
  ): Promise<void> {
    userParticipation.state = state;
    await sendParticipationData(userParticipation);
    await invalidate("event:list");
  }
</script>

<section class="flex justify-end items-end gap-2 flex-wrap">
  <button
          class="chip {chipClasses} preset-outlined-success-500"
          class:preset-filled-success-500={userParticipation.state === "in"}
          onclick={() => updateParticipationStatus("in")}
  >
    <span><Check size="14"/></span>
    <span>{(event.participations.in.length ?? 0) + guestCount}</span>
  </button>

  <button
          class="chip {chipClasses} preset-outlined-warning-500"
          class:preset-filled-warning-500={userParticipation.state === "maybe"}
          onclick={() => updateParticipationStatus("maybe")}
  >
    <span><CircleHelp size="14"/></span>
    <span>{event.participations.maybe.length ?? 0}</span>
  </button>

  <button
          class="chip {chipClasses} preset-outlined-error-500"
          class:preset-filled-error-500={userParticipation.state === "out"}
          onclick={() => updateParticipationStatus("out")}
  >
    <span><X size="14"/></span>
    <span>{event.participations.out.length ?? 0}</span>
  </button>
</section>
