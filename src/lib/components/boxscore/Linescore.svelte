<script lang="ts">
  import type {Linescore} from "bsm.js";
  import {range} from "$lib/functions/range";

  interface Props {
    linescore: Linescore;
  }

  let {linescore}: Props = $props();

  const innings = range(1, linescore.played_innings);
</script>

<div class="table-wrap ">
  <!-- Native Table Element -->
  <table class="linescore bg-surface-50 dark:bg-surface-800">

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