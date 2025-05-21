<script lang="ts">
  import type {EventseriesResponse} from "$lib/model/pb-types.ts";
  import {Weekday} from "$lib/types/Weekday.ts";
  import {DateTimeUtility} from "$lib/service/DateTimeUtility.ts";

  interface Props {
    eventSeries: EventseriesResponse
  }

  let {eventSeries}: Props = $props()

  const weekdayString = $derived(Weekday[eventSeries.weekday as Weekday])
  const startDate = $derived(new Date(eventSeries.series_start))
  const endDate = $derived(new Date(eventSeries.series_end))

  const startTime = DateTimeUtility.convertTimeFromUTC(eventSeries.starttime)
  const endTime = DateTimeUtility.convertTimeFromUTC(eventSeries.endtime)
</script>

<article class="p-4 preset-tonal-surface rounded-base shadow-md">
  <div class="grid grid-cols-1 md:grid-cols-3">

    <div class="flex flex-col justify-between mb-3 md:mb-0">
      <h4 class="h4 font-bold">{eventSeries.title}</h4>
      <div>{weekdayString}</div>
      <div>every {eventSeries.interval} days</div>
    </div>

    <div class="grid grid-cols-2 gap-4 md:col-span-2">
      <!-- Start Time -->
      <div class="flex flex-col items-start">
        <span class="text-sm font-medium">Start Time:</span>
        <time datetime="{startTime}" class="text-sm md:text-base">{startTime}</time>
      </div>
      <!-- Start Date -->
      <div class="flex flex-col items-start">
        <span class="text-sm font-medium">Series Start Date:</span>
        <time class="text-sm md:text-base" datetime="{eventSeries.series_start}">
          {startDate.toLocaleDateString("de-DE", DateTimeUtility.eventSeriesDateFormat)}
        </time>
      </div>
      <!-- End Time -->
      <div class="flex flex-col items-start">
        <span class="text-sm font-medium">End Time:</span>
        <time class="text-sm md:text-base" datetime="{endTime}">{endTime}</time>
      </div>
      <!-- End Date -->
      <div class="flex flex-col items-start">
        <span class="text-sm font-medium">Series End Date:</span>
        <time class="text-sm md:text-base" datetime="{eventSeries.series_end}">
          {endDate.toLocaleDateString("de-DE", DateTimeUtility.eventSeriesDateFormat)}
        </time>
      </div>
    </div>
  </div>
</article>
