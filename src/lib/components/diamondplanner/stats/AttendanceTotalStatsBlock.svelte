<script lang="ts">
  import type {AttendanceTotal} from "$lib/model/PersonalAttendanceStats";
  import StatsBlockContent from "$lib/components/utility/StatsBlockContent.svelte";
  import type {SingleStatElement} from "$lib/types/SingleStatElement";
  import {capitalize} from "$lib/functions/capitalize";

  interface Props {
    season: number | "All Time"
    attendance: AttendanceTotal
  }

  let {attendance, season}: Props = $props()

  const title = capitalize(attendance.type)

  const block: SingleStatElement = {
    title: title,
    value: `${attendance.attended} / ${attendance.total}`,
    desc: `All possible events of this type for ${season}`,
  };
</script>

<div class="card variant-glass-surface shadow-lg">
    <header class="card-header">
        <h2 class="h4 font-semibold">Totals - {title}</h2>
    </header>

    <section class="p-4 flex justify-center">
        <StatsBlockContent {block} classes="place-items-center gap-3"/>
    </section>
</div>