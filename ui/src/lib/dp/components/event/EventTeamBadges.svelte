<script lang="ts">
  import type { ExpandedEvent } from "$lib/dp/types/ExpandedResponse";

  interface Props {
    event: ExpandedEvent;
    isClubWide: boolean;
  }

  let { event, isClubWide }: Props = $props();
</script>

<ul class="badges-container">
  <li>
    <span class="badge preset-outlined-primary-500">
      {#if isClubWide}
        {event.expand?.club?.name}
      {:else}
        {event.expand?.team?.name}
      {/if}
    </span>
  </li>
  {#each event?.expand?.additional_teams as addTeam}
    <li>
      <span class="badge preset-outlined-secondary-800-200">
        {addTeam.name}
      </span>
    </li>
  {/each}
</ul>

<style>
  .badges-container {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    gap: calc(var(--spacing) * 2);
  }
</style>
