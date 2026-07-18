<script lang="ts">
  import { Dialog as SheetPrimitive } from "bits-ui";

  let {
    ref = $bindable(null),
    class: className,
    ...restProps
  }: SheetPrimitive.OverlayProps = $props();

  export { className as class };
</script>

<SheetPrimitive.Overlay
  {...restProps}
  bind:ref
  class={["drawer-overlay", className]}
/>
