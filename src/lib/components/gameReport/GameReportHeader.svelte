<script lang="ts">
  import type {GameReport} from "$lib/model/GameReport.ts";
  import {CalendarDays, ClipboardList, Pen} from "lucide-svelte";

  interface Props {
    report: GameReport;
    classes?: string;
  }

  let {report, classes = ""}: Props = $props();

  const date = $derived(new Date(report.date));
</script>

<section
        class="{classes} card preset-tonal dark:border dark:border-surface-500 rounded-base shadow-xl p-4 my-4">
  <div class="container">
    <CalendarDays/>
    <span class="font-bold">Veröffentlicht: </span>
    <span>{date.toLocaleDateString()}</span>
  </div>

  <hr class="my-2">

  <div class="container">
    <ClipboardList/>
    <span class="font-bold">Liga: </span>
    <span>{report.league.name} ({report.league.season})</span>
  </div>

  <hr class="my-2">

  <div class="container">
    <Pen/>
    <span class="font-bold">Autor: </span>
    <span>{report.author}</span>
  </div>

</section>

<style lang="postcss">
    .container {
        display: flex;
        @apply gap-4;
    }
</style>