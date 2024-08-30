<script lang="ts">
    import type {
        EventsResponse,
        ParticipationsCreate,
        ParticipationsResponse,
    } from "$lib/model/pb-types";
    import EventTypeBadge from "$lib/components/diamondplanner/event/EventTypeBadge.svelte";
    import {DateTimeUtility} from "$lib/service/DateTimeUtility";
    import {
        CalendarMonthOutline,
        CheckOutline,
        ClockOutline,
        CloseOutline,
        MapPinAltOutline, QuestionCircleOutline
    } from "flowbite-svelte-icons";
    import {authModel} from "$lib/pocketbase";
    import type {EventParticipationState} from "$lib/types/EventParticipationState";
    import {sendParticipationData} from "$lib/functions/sendParticipationData";
    import {invalidateAll} from "$app/navigation";

    interface props {
        event: EventsResponse
        link: boolean
    }

    const {event, link}: props = $props()

    const startTime = new Date(event.starttime)
    const meetingTime = new Date(event.meetingtime)

    let participations: ParticipationsResponse[] = $derived(event?.expand?.participations_via_event)
    let participationsIn = $derived(participations?.filter((participation) => {
        return participation.state === "in"
    }))
    let participationsMaybe = $derived(participations?.filter((participation) => {
        return participation.state === "maybe"
    }))
    let participationsOut = $derived(participations?.filter((participation) => {
        return participation.state === "out"
    }))

    let userParticipation: ParticipationsCreate = $state({
        id: "",
        user: $authModel?.id,
        event: event.id,
        state: "",
    })

    async function updateParticipationStatus(state: EventParticipationState): Promise<void> {
        userParticipation = await sendParticipationData(state, userParticipation)
        await invalidateAll()
    }

    $effect(() => {
        participations?.forEach((participation) => {
            if (participation.user === $authModel?.id) {
                userParticipation = participation
            }
        })
    })

</script>

<article class="card variant-soft-surface text-sm h-full" class:card-hover={link}>

    <a href="/account/event/{event.id}">

        <header class="card-header">
            <h2>
                <EventTypeBadge type={event.type}/>
                <span class="ms-1 font-medium">{event?.title}</span>
            </h2>
        </header>

        <section class="p-4">
            <div class="grid grid-cols-2 gap-4">
                <div class="flex col-span-2 gap-2">
                    <CalendarMonthOutline/>
                    <p class="font-bold">

                        {startTime.toLocaleDateString("de-DE", DateTimeUtility.eventDateFormat)}
                    </p>
                </div>

                <div class="flex gap-2">
                    <ClockOutline/>
                    <p>Treffen:
                        {#if event.meetingtime}
                    <span
                        class="font-bold">{meetingTime?.toLocaleTimeString("de-DE", DateTimeUtility.eventTimeFormat)}</span>
                        {:else }
                            <span class="font-medium">---</span>
                        {/if}
                    </p>
                </div>

                <div class="flex gap-2">
                    <ClockOutline/>
                    <p>Start: <span
                        class="font-bold">{startTime.toLocaleTimeString("de-DE", DateTimeUtility.eventTimeFormat)}</span>
                    </p>
                </div>

                <div class="flex col-span-2 gap-2">
                    <MapPinAltOutline/>
                    <p>{event?.location ? event.location : "Kein Ort angegeben."}</p>
                </div>

            </div>
        </section>

    </a>

    <footer class="card-footer">

        <hr class="my-2">

        <div class="flex justify-end items-end gap-2 flex-wrap">

            <button
                class="chip"
                class:variant-filled-success={userParticipation.state === "in"}
                class:variant-ringed-success={userParticipation.state !== "in"}
                onclick={() => updateParticipationStatus("in")}
            >
                <span><CheckOutline size="sm"/></span>
                <span>{participationsIn?.length ?? 0}</span>
            </button>

            <button
                class="chip variant-ringed-warning"
                class:variant-filled-warning={userParticipation.state === "maybe"}
                class:variant-ringed-warning={userParticipation.state !== "maybe"}
                onclick={() => updateParticipationStatus("maybe")}
            >
                <span><QuestionCircleOutline size="sm"/></span>
                <span>{participationsMaybe?.length ?? 0}</span>
            </button>

            <button
                class="chip"
                class:variant-filled-error={userParticipation.state === "out"}
                class:variant-ringed-error={userParticipation.state !== "out"}
                onclick={() => updateParticipationStatus("out") }
            >
                <span><CloseOutline size="sm"/></span>
                <span>{participationsOut?.length ?? 0}</span>
            </button>

        </div>

    </footer>
</article>