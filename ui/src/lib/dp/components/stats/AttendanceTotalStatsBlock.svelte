<script lang="ts">
  import StatsBlockContent from "$lib/dp/components/stats/StatsBlockContent.svelte";
  import type {AttendanceTotal} from "$lib/dp/types/PersonalAttendanceStats.ts";
  import {capitalize} from "$lib/dp/utility/capitalize.ts";
  import type {SingleStatElement} from "$lib/dp/types/SingleStatElement.ts";

  interface Props {
    season: number | "All Time";
    attendance: AttendanceTotal;
  }

  const {attendance, season}: Props = $props();

  const title = $derived(capitalize(attendance.type));

  const block: SingleStatElement = $derived({
    title: title,
    value: `${attendance.attended} / ${attendance.total}`,
    desc: `All possible events of this type for ${season}`,
  });
</script>

<div class="card preset-tonal-surface shadow-lg">
  <header class="card-header">
    <h2 class="h4 title">Totals - {title}</h2>
  </header>

  <section class="content">
    <div class="block-wrapper">
      <StatsBlockContent {block}/>
    </div>
  </section>
</div>

<style>
  .title {
    font-weight: 600;
  }

  .content {
    padding: calc(var(--spacing) * 4);
    display: flex;
    justify-content: center;
  }

  .block-wrapper {
    place-items: center;
    gap: calc(var(--spacing) * 3);
  }
</style>
