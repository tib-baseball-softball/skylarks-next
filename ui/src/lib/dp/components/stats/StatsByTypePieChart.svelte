<script lang="ts">
  import ConicGradient from "$lib/dp/components/charts/ConicGradient.svelte";
  import type {ParticipationType} from "$lib/dp/types/ExpandedResponse.ts";
  import type {ParticipationTotal, PersonalAttendanceStatsItem,} from "$lib/dp/types/PersonalAttendanceStats.ts";
  import {capitalize} from "$lib/dp/utility/capitalize.ts";

  type ConicStop = {
    label: string;
    color: string;
    start: number;
    end: number;
  }

  interface Props {
    statsItem: PersonalAttendanceStatsItem;
  }

  let {statsItem}: Props = $props();

  const PARTICIPATION_COLORS: Record<ParticipationType, string> = {
    in: "var(--color-success-500)",
    maybe: "var(--color-warning-500)",
    out: "var(--color-error-500)",
    "": "",
  };

  function mapToConicStops(participationTotals: ParticipationTotal[]): ConicStop[] {
    let totalAmount = statsItem.totalPossibleEvents;
    if (!totalAmount) {
      totalAmount = 1; // don't divide by zero
    }

    let currentStart = 0;
    return participationTotals.map((p, i) => {
      const percentage = (p.amount / totalAmount) * 100;
      const conicStop: ConicStop = {
        label: capitalize(p.type),
        color: PARTICIPATION_COLORS[p.type],
        start: currentStart,
        end: i === participationTotals.length - 1 ? 100 : currentStart + percentage,
      };

      currentStart = conicStop.end;
      return conicStop;
    });
  }

  const conicStops = $derived(mapToConicStops(statsItem.participationTotals));
</script>

<div class="card preset-tonal-surface shadow-lg">
  <header class="card-header">
    <h2 class="h4 title">Attendance Stats</h2>
  </header>

  <section class="content">
    <ConicGradient legend stops={conicStops}></ConicGradient>
  </section>

  <footer class="card-footer footnote">
    Events where you did not select anything are counted as "out".
  </footer>
</div>

<style>
  .title {
    font-weight: 600;
  }

  .content {
    padding: calc(var(--spacing) * 4);
  }

  .footnote {
    margin-top: calc(var(--spacing) * 1);
    font-weight: var(--font-weight-light);
    font-size: var(--text-sm);
  }
</style>
