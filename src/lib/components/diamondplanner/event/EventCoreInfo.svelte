<script lang="ts">
  import type {ExpandedEvent} from "$lib/model/ExpandedResponse";
  import {DateTimeUtility} from "$lib/service/DateTimeUtility";
  import {Calendar, Clock, MapPin} from "lucide-svelte";
  import TimeSection from "$lib/components/diamondplanner/event/TimeSection.svelte";
  import type {Snippet} from "svelte";

  interface props {
    event: ExpandedEvent;
    classes?: string;
    additionalTimeSection?: Snippet | undefined;
  }

  const {event, classes = "", additionalTimeSection = undefined}: props = $props();

  const startTime = $derived(new Date(event.starttime));
</script>

<section class={classes}>
  <div class="grid grid-cols-6 gap-4">
    <div class="flex items-center gap-2 col-span-6">
      <Calendar size="18"/>
      <time datetime="{event.starttime}" class="font-bold">
        {startTime.toLocaleDateString("de-DE", DateTimeUtility.eventDateFormat)}
      </time>
    </div>

    <TimeSection
            timeValue={event.meetingtime ? event.meetingtime : event.starttime}
            displayText="Meet:"
            classes={additionalTimeSection ? "col-span-2" : "col-span-3"}
    >
      {#snippet icon()}
        <Clock size="18"/>
      {/snippet}
    </TimeSection>

    <TimeSection
            timeValue={event.starttime}
            displayText="Start:"
            classes={additionalTimeSection ? "col-span-2" : "col-span-3"}
    >
      {#snippet icon()}
        <Clock size="18"/>
      {/snippet}
    </TimeSection>

    {@render additionalTimeSection?.()}

    <div class="flex items-center gap-2 col-span-6">
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