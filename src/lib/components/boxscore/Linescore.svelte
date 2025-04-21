<script lang="ts">
  import type {Linescore} from "bsm.js";
  import {range} from "$lib/functions/range";

  interface Props {
    linescore: Linescore;
  }

  let {linescore}: Props = $props();

  const innings = range(1, linescore.played_innings);
</script>

<div class="table-container dark:border-2">
  <!-- Native Table Element -->
  <table class="linescore bg-surface-100-900">

    <thead>
    <tr>
      <th></th>
      {#each innings as inning}
        <th>{inning}</th>
      {/each}
      <th>R</th>
      <th>H</th>
      <th>E</th>
    </tr>
    </thead>

    <tbody>
    <tr>
      <td>{linescore.away.league_entry.team?.name}</td>
      {#each linescore.away.innings as awayInning}
        <td>{awayInning}</td>
      {/each}
      <td class="font-extrabold">{linescore.away.runs}</td>
      <td>{linescore.away.hits}</td>
      <td>{linescore.away.errors}</td>
    </tr>
    <tr>
      <td>{linescore.home.league_entry.team?.name}</td>
      {#each linescore.home.innings as homeInning}
        <td>{homeInning}</td>
      {/each}
      <td class="font-extrabold">{linescore.home.runs}</td>
      <td>{linescore.home.hits}</td>
      <td>{linescore.home.errors}</td>
    </tr>
    </tbody>

  </table>
</div>

<!-- svelte-ignore css_unused_selector -->
<style lang="postcss">

    /* adapted from Skeleton table style */

    .linescore {
        @apply w-full overflow-hidden;
        /* background */
        /* Theme: Rounded */
        @apply rounded-container;
    }

    /* === Hover Styles ==== */

    .linescore-hover tbody tr {
        @apply hover:bg-surface-500/20;
    }

    .linescore-interactive tbody tr {
        /*@apply hover:preset-tonal-primary even:hover:preset-tonal-primary cursor-pointer;*/
    }

    /* === linescore Head === */

    .linescore thead {
        @apply preset-tonal-primary border-b border-surface-500/20;
    }

    .linescore thead tr {
        @apply capitalize text-center;
    }

    .linescore thead th {
        @apply border-l border-surface-500/20 font-bold p-4;
    }

    .linescore thead th:first-child {
        @apply border-0;
    }

    /* === linescore Body === */

    .linescore tbody tr {
        @apply border-b border-surface-500/20 text-center;
    }

    .linescore tbody td {
        @apply border-l border-surface-500/20 px-3 py-4 align-top whitespace-nowrap md:whitespace-normal;
    }

    .linescore tbody td:first-child {
        @apply border-0 text-left;
    }

    .linescore-compact tbody td {
        @apply !py-3;
    }

    .linescore-comforlinescore tbody td {
        @apply !py-5;
    }
</style>