<script lang="ts">
  import { authSettings } from "$lib/dp/client.svelte";
  import { toastController } from "$lib/dp/service/ToastController.svelte";
  import type { CustomAuthModel } from "$lib/dp/types/ExpandedResponse";

  const model = $derived(authSettings.record) as CustomAuthModel;

  async function copy() {
    await navigator.clipboard.writeText(model?.ical_link ?? "");
    toastController.trigger({
      message: "Copied to clipboard!",
      background: "preset-filled-success-500",
    });
  }
</script>

{#if model?.ical_link}
  <header class="card-header">
    <h2 class="h3">Calendar Import</h2>
  </header>

  <div class="card-body">
    <div>
      <h3 class="h4">Your personal calendar link</h3>
      <p class="cal-hint">
        This link includes all events for all teams that you are a member of,
        going back one year.
      </p>
    </div>
    <pre class="">{model?.ical_link}</pre>
    <div class="button-container">
      <button
        onclick={copy}
        class="btn btn-sm preset-tonal-primary border-primary"
        >Copy to clipboard</button
      >
      <a
        class="btn btn-sm preset-tonal-tertiary border-tertiary"
        href={model?.ical_link}
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
  
  .cal-hint {
    margin-block-start: var(--spacing);
  }
</style>
