<script lang="ts">
    import EventTypeBadge from "$lib/components/diamondplanner/event/EventTypeBadge.svelte";
    import type { ExpandedEvent } from "$lib/model/ExpandedResponse";
    import EventCoreInfo from "./EventCoreInfo.svelte";
    import EventParticipationSection from "./EventParticipationSection.svelte";
    import { CloseOutline } from "flowbite-svelte-icons";

    interface props {
        event: ExpandedEvent;
        link: boolean;
    }

    const { event, link }: props = $props();
</script>

<article
    class="card variant-soft-surface shadow-xl text-sm h-full"
    class:card-hover={link}
>
    <a href="/account/event/{event.id}" class:line-through={event.cancelled}>
        <header class="card-header">
            <h2>
                <EventTypeBadge type={event.type} />
                <span class="ms-1 font-medium">{event?.title}</span>
            </h2>
        </header>

        <EventCoreInfo {event} classes={"p-4"} />
    </a>

    <footer class="card-footer">
        <hr class="my-2" />
        {#if event.cancelled}
            <div class="flex justify-end">
                <span class="badge variant-filled-error">
                    <CloseOutline />
                    Cancelled
                </span>
            </div>
        {:else}
            <EventParticipationSection {event} />
        {/if}
    </footer>
</article>
