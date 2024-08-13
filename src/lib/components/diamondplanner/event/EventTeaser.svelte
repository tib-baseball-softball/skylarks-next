<script lang="ts">
    import type {EventsResponse} from "$lib/model/pb-types";
    import EventTypeBadge from "$lib/components/diamondplanner/event/EventTypeBadge.svelte";
    import {DateTimeUtility} from "$lib/service/DateTimeUtility";
    import {
        CalendarMonthOutline,
        CheckOutline,
        ClockOutline,
        CloseOutline,
        MapPinAltOutline, QuestionCircleOutline
    } from "flowbite-svelte-icons";

    interface props {
        event: EventsResponse
        link: boolean
    }

    const {event, link}: props = $props()

    const startTime = new Date(event.starttime)
    const meetingTime = new Date(event.meetingtime)

    let selected = $state('')
</script>

<article class="card variant-ghost-surface text-sm h-full" class:card-hover={link}>

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

        <div class="flex justify-end gap-2 flex-wrap">

            <button
                class="chip variant-ringed-success"
                onclick={() => {  }}
            >
                <span><CheckOutline size="sm"/></span>
                <span>In</span>
            </button>

            <button
                class="chip variant-ringed-warning"
                onclick={() => {  }}
            >
                <span><QuestionCircleOutline size="sm"/></span>
                <span>Evtl.</span>
            </button>

            <button
                class="chip variant-ringed-error"
                onclick={() => {  }}
            >
                <span><CloseOutline size="sm"/></span>
                <span>Out</span>
            </button>

        </div>

    </footer>
</article>