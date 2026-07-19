<script lang="ts">
  import { X } from "lucide-svelte";
  import type { Snippet } from "svelte";

  interface Props {
    open?: Boolean;
    title?: Snippet;
    side: Side;
    triggerClasses?: String;
    triggerContent: Snippet;
    closeButtonClasses?: String;
    children?: Snippet;
  }

  let {
    open = $bindable(false),
    title,
    side,
    triggerClasses = "",
    triggerContent,
    closeButtonClasses = "",
    children,
  }: Props = $props();

  let sheet: HTMLDialogElement;

  export type Side = "bottom" | "left";

  const uid = $props.id();
</script>

<button
  class="sheet-trigger btn {triggerClasses}"
  onclick={() => sheet.showModal()}
>
  {@render triggerContent()}
</button>

<dialog
  class="modal-sheet"
  data-side={side}
  aria-labelledby="sheet-title-{uid}"
  bind:this={sheet}
  closedby="any"
>
  <button
    class="btn preset-outlined-card sheet-close-button {closeButtonClasses}"
    onclick={() => sheet.close()}
  >
    <X aria-hidden="true" />
    <span class="sr-only">Close</span>
  </button>

  <div class="sheet-title-{uid}">
    {@render title?.()}
  </div>

  {@render children?.()}
</dialog>
