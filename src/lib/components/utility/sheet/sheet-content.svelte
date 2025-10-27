<script lang="ts" module>
  export const sheetVariants = {
    base: "bg-surface-50 dark:bg-surface-800 data-[state=open]:animate-in data-[state=closed]:animate-out fixed z-50 gap-4 p-6 shadow-lg transition ease-in-out data-[state=closed]:duration-300 data-[state=open]:duration-480",
    side: {
      top: "data-[state=closed]:slide-out-to-top data-[state=open]:slide-in-from-top inset-x-0 top-0 border-b",
      bottom:
          "data-[state=closed]:slide-out-to-bottom data-[state=open]:slide-in-from-bottom inset-x-0 bottom-0 border-t",
      left: "data-[state=closed]:slide-out-to-left data-[state=open]:slide-in-from-left inset-y-0 left-0 h-full border-r w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%] rounded-lg overflow-y-scroll",
      right:
          "data-[state=closed]:slide-out-to-right data-[state=open]:slide-in-from-right inset-y-0 right-0 h-full w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%] rounded-lg overflow-y-scroll",
    },
  };

  export type Side = "top" | "right" | "bottom" | "left";
</script>

<script lang="ts">
  import {
    Dialog as SheetPrimitive,
    type WithoutChildrenOrChild,
  } from "bits-ui";
  import X from "lucide-svelte/icons/x";
  import type {Snippet} from "svelte";
  import SheetOverlay from "./sheet-overlay.svelte";

  let {
    ref = $bindable(null),
    class: className,
    side = "right",
    portalProps,
    children,
    ...restProps
  }: WithoutChildrenOrChild<SheetPrimitive.ContentProps> & {
    portalProps?: SheetPrimitive.PortalProps;
    side?: Side;
    children: Snippet;
  } = $props();
</script>

<SheetPrimitive.Portal {...portalProps}>
  <SheetOverlay/>
  <SheetPrimitive.Content
          {...restProps}
          bind:ref
          class={[sheetVariants.side[side], className]}
  >
    <SheetPrimitive.Close
            class="btn preset-tonal-surface border border-surface-500 disabled:pointer-events-none sheet-close-button"
    >
      <X/>
      <span class="sr-only">Close</span>
    </SheetPrimitive.Close>

    {@render children?.()}
  </SheetPrimitive.Content>
</SheetPrimitive.Portal>
