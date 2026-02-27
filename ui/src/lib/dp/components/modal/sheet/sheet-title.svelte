<script lang="ts">
  import {Dialog as SheetPrimitive} from "bits-ui";

  let {ref = $bindable(null), class: className, ...restProps}: SheetPrimitive.TitleProps = $props();
</script>

<SheetPrimitive.Title
  {...restProps}
  bind:ref
  class={["text-foreground", "root", className]}
/>

<style>
  .root {
    font-size: 1.125rem; /* text-lg */
    font-weight: var(--font-weight-semibold, 600);
  }
</style>
