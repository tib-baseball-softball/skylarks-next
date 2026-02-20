<script lang="ts">
  import {Calendar, Clock, MapPin} from "lucide-svelte";
  import type {Snippet} from "svelte";
  import TimeSection from "$lib/dp/components/event/TimeSection.svelte";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";
  import type {ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";

  interface props {
    event: ExpandedEvent;
    additionalTimeSection?: Snippet | undefined;
  }

  const {event, additionalTimeSection = undefined}: props = $props();

  const startTime = $derived(new Date(event.starttime));
</script>

<section>
  <div class="info-grid">
    <div class="info-row">
      <Calendar size="18"/>
      <time class="event-time" datetime="{event.starttime}">
        {startTime.toLocaleDateString("de-DE", DateTimeUtility.eventDateFormat)}
      </time>
    </div>

    <TimeSection
      classes={additionalTimeSection ? "time-section-narrow" : "time-section-wide"}
      displayText="Meet:"
      timeValue={event.meetingtime ? event.meetingtime : event.starttime}
    >
      {#snippet icon()}
        <Clock size="18"/>
      {/snippet}
    </TimeSection>

    <TimeSection
      classes={additionalTimeSection ? "time-section-narrow" : "time-section-wide"}
      displayText="Start:"
      timeValue={event.starttime}
    >
      {#snippet icon()}
        <Clock size="18"/>
      {/snippet}
    </TimeSection>

    {@render additionalTimeSection?.()}

    <div class="info-row">
      <MapPin size="18"/>
      {#if event?.expand?.location}
        <p>
          {event?.expand?.location.address_addon}, {event?.expand?.location
          .city}
        </p>
      {:else}
        <p>No location provided.</p>
      {/if}
    </div>
  </div>
</section>

<style>
  .info-grid {
    display: grid;
    grid-template-columns: repeat(6, minmax(0, 1fr));
    gap: calc(var(--spacing) * 4);
  }

  .info-row {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 2);
    grid-column: span 6 / span 6;
  }

  .event-time {
    font-weight: var(--font-weight-bold);
  }

  :global(.time-section-narrow) {
    grid-column: span 2 / span 2;
  }

  :global(.time-section-wide) {
    grid-column: span 3 / span 3;
  }
</style>
