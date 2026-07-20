<script lang="ts">
  import { X } from "lucide-svelte";
  import { onMount, type Snippet } from "svelte";

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

  export type Side = "right" | "left";

  const uid = $props.id();

  function show() {
    open = true;
  }

  function close() {
    open = false;
  }

  // manual data sync between dialog element and component props
  // enables closing dialog from parent elements by setting `open`
  $effect(() => {
    if (open === true) {
      sheet.showModal();
    } else {
      sheet.close();
    }
  });

  onMount(() => {
    sheet.addEventListener("close", () => {
      open = false;
    });
  });
</script>

<button class="sheet-trigger btn {triggerClasses}" onclick={show}>
  {@render triggerContent()}
</button>

<dialog
  class="modal-sheet"
  data-side={side}
  data-open={String(open)}
  aria-labelledby="sheet-title-{uid}"
  bind:this={sheet}
  closedby="any"
>
  <button
    class="btn preset-outlined-card sheet-close-button {closeButtonClasses}"
    onclick={close}
  >
    <X aria-hidden="true" />
    <span class="sr-only">Close</span>
  </button>

  <div class="sheet-title-{uid}">
    {@render title?.()}
  </div>

  {@render children?.()}
</dialog>
