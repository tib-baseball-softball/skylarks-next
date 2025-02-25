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
  }

  const authRecord = authSettings.record as CustomAuthModel;
  const {event}: props = $props();

  let userParticipation: ParticipationsCreate = $derived(event.userParticipation ?? {
    id: "",
    user: authRecord?.id,
    event: event.id,
    state: "",
    comment: "",
  });

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
          class="chip"
          class:variant-filled-success={userParticipation.state === "in"}
          class:variant-ringed-success={userParticipation.state !== "in"}
          onclick={() => updateParticipationStatus("in")}
  >
    <span><Check size="14"/></span>
    <span>{event.participations.in.length ?? 0}</span>
  </button>

  <button
          class="chip variant-ringed-warning"
          class:variant-filled-warning={userParticipation.state === "maybe"}
          class:variant-ringed-warning={userParticipation.state !== "maybe"}
          onclick={() => updateParticipationStatus("maybe")}
  >
    <span><CircleHelp size="14"/></span>
    <span>{event.participations.maybe.length ?? 0}</span>
  </button>

  <button
          class="chip"
          class:variant-filled-error={userParticipation.state === "out"}
          class:variant-ringed-error={userParticipation.state !== "out"}
          onclick={() => updateParticipationStatus("out")}
  >
    <span><X size="14"/></span>
    <span>{event.participations.out.length ?? 0}</span>
  </button>
</section>
