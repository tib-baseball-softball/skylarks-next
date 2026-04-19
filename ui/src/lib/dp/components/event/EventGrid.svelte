<script lang="ts">
  import EventTeaser from "$lib/dp/components/event/EventTeaser.svelte";
  import type {ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    events: ExpandedEvent[];
  }

  const {events}: Props = $props();
</script>

<div class="events-grid">
  {#each events as event (event.id)}
    <div>
      <EventTeaser {event} link={true}/>
    </div>
  {:else}
    <p>No events available with the current filters.</p>
  {/each}
</div>

<style>
  .events-grid {
    margin-block: calc(var(--spacing) * 4);
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 4);

    @media (min-width: 48rem) {
      grid-template-columns: 1fr 1fr;
    }

    @media (min-width: 80rem) {
      grid-template-columns: 1fr 1fr 1fr;
    }
  }
</style>
