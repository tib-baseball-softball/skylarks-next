<script lang="ts">
  import type {Snippet} from "svelte";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";

  interface Props {
    timeValue: string;
    displayText: string;
    icon: Snippet;
    classes?: string;
  }

  const {timeValue, displayText, icon, classes = ""}: Props = $props();

  const displayedTime = $derived(new Date(timeValue));
</script>

<div class="section-container {classes}">
  {@render icon()}
  <p>
    {displayText}
    {#if timeValue}
      <time datetime="{timeValue}" class="time"
      >{displayedTime?.toLocaleTimeString(
        "de-DE",
        DateTimeUtility.eventTimeFormat,
      )}</time
      >
    {:else}
      <span class="fallback">---</span>
    {/if}
  </p>
</div>

<style>
  .section-container {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .time {
    font-weight: 700; /* font-bold */
  }

  .fallback {
    font-weight: 500; /* font-medium */
  }
</style>
