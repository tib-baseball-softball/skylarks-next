<script lang="ts">
    import EventTypeBadge from "$lib/components/diamondplanner/event/EventTypeBadge.svelte";
    import type { EventsResponse } from "$lib/model/pb-types";
    import EventCoreInfo from "./EventCoreInfo.svelte";
    import EventParticipationSection from "./EventParticipationSection.svelte";

    interface props {
        event: EventsResponse;
        link: boolean;
    }

    const { event, link }: props = $props();
</script>

<article
    class="card variant-soft-surface text-sm h-full"
    class:card-hover={link}
>
    <a href="/account/event/{event.id}">
        <header class="card-header">
            <h2>
                <EventTypeBadge type={event.type} />
                <span class="ms-1 font-medium">{event?.title}</span>
            </h2>
        </header>

        <EventCoreInfo {event} classes={"p-4"}/>
    </a>

    <footer class="card-footer">
        <hr class="my-2" />

        <EventParticipationSection {event} />
    </footer>
</article>
