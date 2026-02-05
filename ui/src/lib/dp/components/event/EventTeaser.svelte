<script lang="ts">
  import {Ban} from "lucide-svelte";
  import EventTypeBadge from "$lib/dp/components/event/EventTypeBadge.svelte";
  import type {ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";
  import EventCoreInfo from "./EventCoreInfo.svelte";
  import EventParticipationSection from "./EventParticipationSection.svelte";

  interface props {
    event: ExpandedEvent;
    link: boolean;
  }

  const {event, link}: props = $props();
</script>

<article
  class="card preset-tonal-surface shadow-xl text-sm h-full"
  class:card-hover={link}
>
  <a class:line-through={event.cancelled} href="/account/event/{event.id}">
    <header class="card-header">
      <h2>
        <EventTypeBadge type={event.type}/>
        <span class="ms-1 font-bold h5">{event?.title}</span>
      </h2>
    </header>

    <EventCoreInfo classes={"p-4"} {event}/>
  </a>

  <footer class="card-footer">
    <hr class="my-2"/>
    {#if event.cancelled}
      <div class="flex justify-end">
                <span class="badge preset-filled-error-500 gap-1">
                    <Ban size="18"/>
                    Cancelled
                </span>
      </div>
    {:else}
      <EventParticipationSection {event}/>
    {/if}
  </footer>
</article>
