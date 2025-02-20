<script lang="ts">
    import type {PageStore} from "$lib/pocketbase/PageStore";
    import {ChevronLeft, ChevronRight} from "lucide-svelte";

    const {
    store,
    showIfSinglePage = false,
  }: {
    store: PageStore;
    showIfSinglePage?: boolean;
  } = $props();
</script>

{#if showIfSinglePage || $store.totalPages > 1}
  <div
          class="paginator flex flex-col md:flex-row items-center space-y-4 md:space-y-0 md:space-x-4 justify-between"
  >
    <div class="paginator-controls btn-group variant-soft-surface">
      <button
              class="fill-current"
              type="button"
              onclick={() => store.prev()}
              disabled={$store.page <= 1}
      >
        <ChevronLeft/>
      </button>

      <button class="pointer-events-none !text-sm"
      >{$store.page}/{$store.totalPages}</button
      >

      <button
              class="fill-current"
              type="button"
              onclick={() => store.next()}
              disabled={$store.page >= $store.totalPages}
      >
        <ChevronRight/>
      </button>
    </div>
  </div>
{/if}
