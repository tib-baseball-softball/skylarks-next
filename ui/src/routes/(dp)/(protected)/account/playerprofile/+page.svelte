<script lang="ts">
  import PlayerDataProfileSection from "$lib/dp/components/user/PlayerDataProfileSection.svelte";
</script>

<svelte:head>
  <title>My Player Data</title>
  <meta
    content="Player profile page, displaying player data and stats."
    name="description"
  />
</svelte:head>


<h1 class="h1">My Player Data</h1>

<div class="wrapper">
  <PlayerDataProfileSection/>
</div>

<style>
  .wrapper {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 5);
    margin-block-end: calc(var(--spacing) * 6);

    @media (min-width: 64rem) {
      grid-template-columns: 1fr 1fr;
    }

    @media (min-width: 80rem) {
      grid-template-columns: 1fr 1fr 1fr;
    }
  }
</style>
