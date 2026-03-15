<script lang="ts">
  import { toastController } from "$lib/dp/service/ToastController.svelte";
  import type { Snippet } from "svelte";

  interface Props {
    link: string;
    header: Snippet;
    subheader: Snippet;
  }

  let { link, header, subheader }: Props = $props();

  async function copy() {
    await navigator.clipboard.writeText(link);
    toastController.trigger({
      message: "Copied to clipboard!",
      background: "preset-filled-success-500",
    });
  }
</script>

{#if link}
  <header class="card-header">
    <h2 class="h3">
      {@render header()}
    </h2>
  </header>

  <div class="card-body">
    <div>
      {@render subheader()}
    </div>
    <pre>{link}</pre>
    <div class="button-container">
      <button
        onclick={copy}
        class="btn btn-sm preset-tonal-primary border-primary"
        >Copy to clipboard</button
      >
      <a
        class="btn btn-sm preset-tonal-tertiary border-tertiary"
        href={link}
        target="_blank">Download .ics file</a
      >
    </div>
  </div>
{/if}

<style>
  .card-body {
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 6);
  }

  .button-container {
    display: flex;
    gap: calc(var(--spacing) * 3);
    flex-wrap: wrap;
  }

  pre {
    overflow-x: auto;
  }
</style>
