<script lang="ts">
    import type {
        ParticipationsCreate,
        ParticipationsResponse,
        EventsResponse,
    } from "$lib/model/pb-types";
    import {
        CheckOutline,
        CloseOutline,
        QuestionCircleOutline,
    } from "flowbite-svelte-icons";
    import type { EventParticipationState } from "$lib/types/EventParticipationState";
    import { sendParticipationData } from "$lib/functions/sendParticipationData";
    import { invalidate } from "$app/navigation";
    import { authModel } from "$lib/pocketbase/Auth";

    interface props {
        event: EventsResponse;
    }

    const { event }: props = $props();

    let userParticipation: ParticipationsCreate = $state({
        id: "",
        user: $authModel?.id,
        event: event.id,
        state: "",
        comment: "",
    });

    let comment = $state("");

    let participations: ParticipationsResponse[] = $derived(
        event?.expand?.participations_via_event,
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

    async function updateParticipationStatus(
        state: EventParticipationState,
    ): Promise<void> {
        userParticipation.state = state;
        userParticipation.comment = comment;
        userParticipation = await sendParticipationData(userParticipation);
        await invalidate("event:list");
    }

    $effect(() => {
        participations?.forEach((participation) => {
            if (participation.user === $authModel?.id) {
                userParticipation = participation;
            }
        });
    });
</script>

<section class="flex justify-end items-end gap-2 flex-wrap">
    <button
        class="chip"
        class:variant-filled-success={userParticipation.state === "in"}
        class:variant-ringed-success={userParticipation.state !== "in"}
        onclick={() => updateParticipationStatus("in")}
    >
        <span><CheckOutline size="sm" /></span>
        <span>{participationsIn?.length ?? 0}</span>
    </button>

    <button
        class="chip variant-ringed-warning"
        class:variant-filled-warning={userParticipation.state === "maybe"}
        class:variant-ringed-warning={userParticipation.state !== "maybe"}
        onclick={() => updateParticipationStatus("maybe")}
    >
        <span><QuestionCircleOutline size="sm" /></span>
        <span>{participationsMaybe?.length ?? 0}</span>
    </button>

    <button
        class="chip"
        class:variant-filled-error={userParticipation.state === "out"}
        class:variant-ringed-error={userParticipation.state !== "out"}
        onclick={() => updateParticipationStatus("out")}
    >
        <span><CloseOutline size="sm" /></span>
        <span>{participationsOut?.length ?? 0}</span>
    </button>
</section>
