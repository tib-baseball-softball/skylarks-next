<script lang="ts">
  import type {EventseriesResponse} from "$lib/model/pb-types.ts";
  import {DateTimeUtility} from "$lib/service/DateTimeUtility.ts";

  interface Props {
    eventSeries: EventseriesResponse
  }

  let {eventSeries}: Props = $props()

  const startDate = $derived(new Date(eventSeries.series_start))
  const endDate = $derived(new Date(eventSeries.series_end))
</script>

<article class="p-4 preset-outline-surface rounded-base shadow-md">
  <div class="grid grid-cols-1 md:grid-cols-3 md:gap-2 lg:gap-4">

    <div class="flex flex-col justify-between mb-3 md:mb-0">
      <h4 class="h5 font-bold">{eventSeries.title}</h4>
      <div>every {eventSeries.interval} days</div>
    </div>

    <div class="grid grid-cols-2 gap-4 md:col-span-2">
      <!-- Start Date -->
      <div class="flex flex-col items-start">
        <span class="text-sm font-medium">Series First Occurrence:</span>
        <time class="text-sm md:text-base" datetime="{eventSeries.series_start}">
          {startDate.toLocaleDateString("de-DE", DateTimeUtility.eventSeriesDateTimeFormat)}
        </time>
      </div>
      <!-- End Date -->
      <div class="flex flex-col items-start">
        <span class="text-sm font-medium">Series End After:</span>
        <time class="text-sm md:text-base" datetime="{eventSeries.series_end}">
          {endDate.toLocaleDateString("de-DE", DateTimeUtility.eventSeriesDateFormat)}
        </time>
      </div>
      <!-- Start Time -->
      <div class="flex flex-col items-start">
        <span class="text-sm font-medium">Event Duration:</span>
        <time datetime="P{eventSeries.duration}M" class="text-sm md:text-base">{eventSeries.duration} Minutes</time>
      </div>
    </div>
  </div>
</article>
