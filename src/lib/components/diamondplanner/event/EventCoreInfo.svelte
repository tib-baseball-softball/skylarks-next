<script lang="ts">
    import type {ExpandedEvent} from "$lib/model/ExpandedResponse";
    import {DateTimeUtility} from "$lib/service/DateTimeUtility";
    import {Calendar, Clock, MapPin} from "lucide-svelte";

    interface props {
    event: ExpandedEvent;
    classes?: string;
  }

  const {event, classes = ""}: props = $props();

  const startTime = $derived(new Date(event.starttime));
  const meetingTime = $derived(new Date(event.meetingtime));
</script>

<section class={classes}>
  <div class="grid grid-cols-2 gap-4">
    <div class="section-container col-span-2">
      <Calendar size="18"/>
      <p class="font-bold">
        {startTime.toLocaleDateString(
            "de-DE",
            DateTimeUtility.eventDateFormat,
        )}
      </p>
    </div>

    <div class="section-container">
      <Clock size="18"/>
      <p>
        Meet:
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

    <div class="section-container">
      <Clock size="18"/>
      <p>
        Start: <span class="font-bold"
      >{startTime.toLocaleTimeString(
          "de-DE",
          DateTimeUtility.eventTimeFormat,
      )}</span
      >
      </p>
    </div>

    <div class="section-container col-span-2">
      <MapPin size="18"/>
      <p>
        {event?.location ? event.location : "No location provided."}
      </p>
    </div>
  </div>
</section>

<style>
    .section-container {
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }
</style>
