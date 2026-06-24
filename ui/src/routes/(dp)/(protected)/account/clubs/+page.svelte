<script lang="ts">
  import type {PageProps} from "./$types";
  import PlayerProfileClubsSection from "$lib/dp/components/user/PlayerProfileClubsSection.svelte";

  const {data}: PageProps = $props();
</script>

<svelte:head>
	<title>My Clubs</title>
</svelte:head>

<h1 class="h1">My Clubs</h1>

<div class="grid">
  <PlayerProfileClubsSection clubs={data.clubs}/>
</div>

<style>
  .grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 3);
    margin-block-end: calc(var(--spacing) * 6);

    @media (min-width: 64rem) {
      gap: calc(var(--spacing) * 6);
      grid-template-columns: 1fr 1fr;
    }

    @media (min-width: 80rem) {
      grid-template-columns: 1fr 1fr 1fr;
    }
  }
</style>
