<script lang="ts">
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";
  import {EventSeriesUtility} from "$lib/dp/service/EventSeriesUtility.ts";
  import type {EventseriesResponse} from "$lib/dp/types/pb-types.ts";
  import {appLocale} from "$lib/dp/locale.svelte.ts";

  interface Props {
    eventSeries: EventseriesResponse;
  }

  const {eventSeries}: Props = $props();

  const startDate = $derived(new Date(eventSeries.series_start));
  const endDate = $derived(new Date(eventSeries.series_end));

  const options: { weekday: "long" | "short" | "narrow" | undefined } = {weekday: "long"};
  const seriesState = $derived(EventSeriesUtility.getSeriesState(startDate, endDate));
</script>

<article class="preset-outlined-surface-800-200 rounded-base shadow-md">
  <div class="outer-grid">

    <div class="title-wrapper">
      <h4 class="h5">{eventSeries.title}</h4>
      <div>every {eventSeries.interval} days</div>
    </div>

    <div class="inner-grid">
      <!-- Start Date -->
      <div class="date-wrapper">
        <span class="date-explanation-text">Series First Occurrence:</span>
        <time datetime="{eventSeries.series_start}">
          {startDate.toLocaleDateString("de-DE", DateTimeUtility.eventSeriesDateTimeFormat)}
        </time>
      </div>

      <!-- End Date -->
      <div class="date-wrapper">
        <span class="date-explanation-text">Series End After:</span>
        <time datetime="{eventSeries.series_end}">
          {endDate.toLocaleDateString("de-DE", DateTimeUtility.eventSeriesDateFormat)}
        </time>
      </div>

      <!-- Start Time -->
      <div class="date-wrapper">
        <span class="date-explanation-text">Event Duration:</span>
        <time datetime="P{eventSeries.duration}M">{eventSeries.duration} Minutes</time>
      </div>

      <div class="badges">
        <span class="badge preset-outlined">
          {new Intl.DateTimeFormat(appLocale.current, options).format(startDate)}
        </span>

        {#if seriesState === "ongoing"}
        <span class="badge preset-tonal-primary border border-primary-500">
          Ongoing
        </span>
        {:else if seriesState === "past"}
          <span class="badge preset-outlined">
          Past
        </span>
        {:else if seriesState === "future"}
          <span class="badge preset-tonal-tertiary border border-tertiary-500">
          Upcoming
        </span>
        {/if}
      </div>
    </div>
  </div>
</article>

<style>
  article {
    padding: calc(var(--spacing) * 4);
  }

  .outer-grid {
    display: grid;
    grid-template-columns: repeat(1, minmax(0, 1fr));

    @media (min-width: 40rem) {
      grid-template-columns: repeat(3, minmax(0, 1fr));
      gap: calc(var(--spacing) * 2);
    }
  }

  .inner-grid {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: calc(var(--spacing) * 4);

    @media (min-width: 40rem) {
      grid-column: span 2 / span 2;
    }
  }

  .date-wrapper {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }

  .date-explanation-text {
    font-size: var(--text-xs);
    line-height: var(--tw-leading, var(--text-xs--line-height));
    --tw-font-weight: var(--font-weight-medium);
    font-weight: var(--font-weight-medium);
  }

  .title-wrapper {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    margin-block-end: calc(var(--spacing) * 3);
  }

  time {
    font-size: var(--text-sm);
    line-height: var(--tw-leading, var(--text-sm--line-height));
    --tw-font-weight: var(--font-weight-medium);
    font-weight: var(--font-weight-medium);
  }

  h4 {
    font-weight: var(--font-weight-bold);
  }

  .badge {
    display: block;
  }

  .badges {
    justify-self: start;
    align-self: center;
    display: flex;
    gap: calc(var(--spacing) * 2);
    flex-wrap: wrap;
  }
</style>
