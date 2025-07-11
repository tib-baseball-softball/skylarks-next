<script lang="ts">
  import type {PageStore} from "$lib/pocketbase/PageStore.ts";
  import type {ExpandedAnnouncement} from "$lib/model/ExpandedResponse.ts";
  import AnnouncementCard from "$lib/components/announcements/AnnouncementCard.svelte";
  import Paginator from "$lib/pocketbase/Paginator.svelte";

  interface Props {
    store: PageStore<ExpandedAnnouncement>
  }

  let {store}: Props = $props();
</script>

<div>
  {#if $store.totalItems > 0}
    <ul class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-5 mb-6">
      {#each $store.items as announcement (announcement.id)}
        <li>
          <AnnouncementCard {announcement} />
        </li>
      {/each}
    </ul>

    <Paginator {store} showIfSinglePage={false} />
  {:else}
    <p>No announcements available.</p>
  {/if}
</div>