<script lang="ts">
  import type {ParticipationTotal, PersonalAttendanceStatsItem} from "$lib/model/PersonalAttendanceStats";
  import {type ConicStop} from "@skeletonlabs/skeleton-svelte";
  import type {ParticipationType} from "$lib/model/ExpandedResponse";
  import {capitalize} from "$lib/functions/capitalize";

  interface Props {
    statsItem: PersonalAttendanceStatsItem
  }

  let {statsItem}: Props = $props()

  const PARTICIPATION_COLORS: Record<ParticipationType, string> = {
    in: "rgb(var(--color-success-500))",
    maybe: "rgb(var(--color-warning-500))",
    out: "rgb(var(--color-error-500))",
    "": "",
  };

  function mapToConicStops(participationTotals: ParticipationTotal[]): ConicStop[] {
    let totalAmount = statsItem.totalPossibleEvents;
    if (!totalAmount) {
      totalAmount = 1 // don't divide by zero
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

  const conicStops = mapToConicStops(statsItem.participationTotals)
</script>

<div class="card preset-tonal-surface shadow-lg">
    <header class="card-header">
        <h2 class="h4 font-semibold">Attendance Stats</h2>
    </header>

    <section class="p-4">
        <ConicGradient legend stops={conicStops}></ConicGradient>
    </section>

    <footer class="mt-1 card-footer font-light text-sm">
        Events where you did not select anything are counted as "out".
    </footer>
</div>