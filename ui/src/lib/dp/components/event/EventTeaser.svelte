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
  class="card preset-tonal-surface shadow-xl teaser"
  class:card-hover={link}
>
  <a class:line-through={event.cancelled} href="/account/event/{event.id}">
    <header class="card-header">
      <h2 class="title">
        <EventTypeBadge type={event.type}/>
        <span class="title-text h5">{event?.title}</span>
      </h2>
    </header>

    <EventCoreInfo {event}/>
  </a>

  <footer class="card-footer">
    <hr class="divider"/>
    {#if event.cancelled}
      <div class="cancelled-container">
                <span class="badge preset-filled-error-500 cancelled-badge">
                    <Ban size="18"/>
                    Cancelled
                </span>
      </div>
    {:else}
      <EventParticipationSection {event}/>
    {/if}
  </footer>
</article>

<style>
  .teaser {
    font-size: var(--text-sm);
    height: 100%;
  }

  .title-text {
    margin-inline-start: calc(var(--spacing) * 1);
    font-weight: var(--font-weight-bold);
  }

  .divider {
    margin-top: calc(var(--spacing) * 2);
    margin-bottom: calc(var(--spacing) * 2);
  }

  .cancelled-container {
    display: flex;
    justify-content: end;
  }

  .cancelled-badge {
    gap: calc(var(--spacing) * 1);
  }
</style>
