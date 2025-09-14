<script lang="ts">
  import type {Player} from "$lib/model/Player.ts";
  import PlayerRow from "$lib/components/player/PlayerRow.svelte";

  interface Props {
    players: Player[];
  }

  let {players}: Props = $props();
</script>

<div class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl">
  {#each players as player (player.bsm_id)}
    <PlayerRow {player}/>

    <hr class="my-3">
  {/each}
</div>
