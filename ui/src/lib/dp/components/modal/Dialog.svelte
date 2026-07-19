<script lang="ts">
  import type { Snippet } from "svelte";
  import { X } from "lucide-svelte";
  import type { HTMLAttributes, HTMLButtonAttributes } from "svelte/elements";

  let dialog: HTMLDialogElement;

  type Props = {
    open?: boolean;
    children?: Snippet;
    triggerClasses?: string;
    title: Snippet;
    description?: Snippet;
    contentProps?: HTMLAttributes<HTMLDialogElement>;
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

  const uid = $props.id();
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

<dialog
  {...contentProps}
  class="modal-dialog card dialog-content shadow-2xl"
  bind:this={dialog}
  closedby="any"
  aria-labelledby="dialog-title-{uid}"
>
  <header class="header">
    <menu>
      <button
        onclick={() => dialog.close()}
        class="close-button btn preset-outlined-card {closeButtonClasses}"
        data-dialog-close
      >
        <X />
      </button>
    </menu>

    {#if title}
      <div id="dialog-title-{uid}" class="title">
        {@render title?.()}
      </div>
    {/if}
  </header>

  <article>
    {#if description}
      {@render description?.()}
    {/if}

    {@render children?.()}
  </article>
</dialog>
