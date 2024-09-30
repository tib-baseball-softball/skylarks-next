<script lang="ts">
    import type { ExpandedEvent } from "$lib/model/ExpandedResponse";
    import { DateTimeUtility } from "$lib/service/DateTimeUtility";
    import { CalendarMonthOutline, ClockOutline, MapPinOutline } from "flowbite-svelte-icons";

    interface props {
        event: ExpandedEvent;
        classes?: string;
    }

    const { event, classes = "" }: props = $props();

    const startTime = $derived(new Date(event.starttime));
    const meetingTime = $derived(new Date(event.meetingtime));
</script>

<section class={classes}>
    <div class="grid grid-cols-2 gap-4">
        <div class="flex col-span-2 gap-2">
            <CalendarMonthOutline />
            <p class="font-bold">
                {startTime.toLocaleDateString(
                    "de-DE",
                    DateTimeUtility.eventDateFormat,
                )}
            </p>
        </div>

        <div class="flex gap-2">
            <ClockOutline />
            <p>
                Treffen:
                {#if event.meetingtime}
                    <span class="font-bold"
                        >{meetingTime?.toLocaleTimeString(
                            "de-DE",
                            DateTimeUtility.eventTimeFormat,
                        )}</span
                    >
                {:else}
                    <span class="font-medium">---</span>
                {/if}
            </p>
        </div>

        <div class="flex gap-2">
            <ClockOutline />
            <p>
                Start: <span class="font-bold"
                    >{startTime.toLocaleTimeString(
                        "de-DE",
                        DateTimeUtility.eventTimeFormat,
                    )}</span
                >
            </p>
        </div>

        <div class="flex col-span-2 gap-2">
            <MapPinOutline />
            <p>
                {event?.location ? event.location : "Kein Ort angegeben."}
            </p>
        </div>
    </div>
</section>
