<script lang="ts">
  import { ChevronLeft, ChevronRight } from "lucide-svelte";
  import type { PageStore } from "$lib/dp/records/PageStore.ts";
  import { page } from "$app/state";

  interface Props {
    store: PageStore;
    showIfSinglePage?: boolean;
  }

  const {
    store,
    showIfSinglePage = false,
  }: Props = $props();

  type PagingAction = "increment" | "decrement";

  function pushNavState(action: PagingAction) {
    const newURL = page.url;
    const currentPage = $store.page
    switch (action) {
      case "increment":
        newURL.searchParams.set("page", String(currentPage + 1));
        break;
      case "decrement":
        newURL.searchParams.set("page", String(currentPage - 1));
        break;
    }
    history.pushState(null, "", newURL);
  }

  function prev() {
    store.prev();
    pushNavState("decrement")
  }

  function next() {
    store.next();
    pushNavState("increment")
  }
</script>

{#if showIfSinglePage || $store.totalPages > 1}
  <div class="paginator root">
    <div class="controls-wrapper rounded-container preset-outlined-card">
      <button
        class="nav-button rounded-container"
        type="button"
        onclick={prev}
        disabled={$store.page <= 1}
      >
        <ChevronLeft />
      </button>

      <span class="page-indicator">{$store.page}/{$store.totalPages}</span>

      <button
        class="nav-button rounded-container"
        type="button"
        onclick={next}
        disabled={$store.page >= $store.totalPages}
      >
        <ChevronRight />
      </button>
    </div>
  </div>
{/if}

<style>
  .root {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: calc(var(--spacing) * 4);
    justify-content: space-between;
  }

  @media (min-width: 48rem) {
    .root {
      flex-direction: row;
    }
  }

  .controls-wrapper {
    display: flex;
    padding: calc(var(--spacing) * 1);
  }

  .nav-button {
    fill: currentColor;
    padding: calc(var(--spacing) * 2);
  }

  .page-indicator {
    font-size: var(--text-sm) !important;
    padding: calc(var(--spacing) * 2);
  }

  button:hover:not(:disabled),
  button:focus:not(:disabled) {
    background-color: var(--color-primary-50-950);
    color: var(--color-primary-950-50);
  }
</style>
