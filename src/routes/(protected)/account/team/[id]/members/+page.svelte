<script lang="ts">
  let { data } = $props();

  const players = $derived(data.players)
</script>

{#each $players.items as player}
  <p>{player.last_name}</p>
{/each}