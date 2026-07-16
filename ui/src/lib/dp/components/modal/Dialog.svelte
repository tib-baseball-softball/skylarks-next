<script lang="ts">
  import type { Snippet } from "svelte";
  import { X } from "lucide-svelte";
  import type { HTMLAttributes, HTMLButtonAttributes } from "svelte/elements";

  let dialog: HTMLDialogElement;

  type Props = {
    open: boolean;
    children?: Snippet;
    triggerClasses?: string;
    title: Snippet;
    description?: Snippet;
    contentProps?: HTMLAttributes<HTMLDivElement>;
    triggerProps?: HTMLButtonAttributes;
    triggerContent: Snippet;
    closeButtonClasses?: string;
    disabled?: boolean;
  };

  let {
    open = $bindable(false),
    children,
    triggerClasses = "",
    contentProps,
    triggerProps,
    title,
    description,
    triggerContent,
    closeButtonClasses = "",
    disabled = false,
  }: Props = $props();
</script>

<button
  onclick={() => {
    dialog.showModal();
  }}
  {...triggerProps}
  class="{triggerClasses} trigger"
  {disabled}
  type="button"
>
  {@render triggerContent()}
</button>

<dialog class="modal-dialog" bind:this={dialog} closedby="any">
  <div {...contentProps} class="card dialog-content shadow-2xl">
    <div class="header">
      <button
        onclick={() => dialog.close()}
        class="close-button btn preset-outlined-card {closeButtonClasses}"
      >
        <X />
      </button>

      <div class="title">
        {@render title?.()}
      </div>
    </div>

    {#if description}
      {@render description?.()}
    {/if}

    {@render children?.()}
  </div>
</dialog>

<style>
  .trigger {
    display: flex;
    gap: var(--spacing);
  }

  .header {
    display: flex;
    gap: calc(var(--spacing) * 5);
    align-items: center;
    margin-bottom: calc(var(--spacing) * 2);
  }

  .title {
    font-size: var(--text-xl);
    font-weight: var(--font-weight-semibold);
  }

  .close-button {
    border: 1px solid;
  }
</style>
