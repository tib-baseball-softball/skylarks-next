<script lang="ts">
  import type {Snippet} from "svelte";

  interface Props {
    icon: Snippet,
    categoryName: string,
    rowValue: string,
  }

  let {icon, categoryName, rowValue}: Props = $props();
</script>

<div class="container">
  {@render icon()}
  <p>
    <strong class="me-2">{categoryName}:</strong>{rowValue}
  </p>
</div>

<style lang="postcss">
    .container {
        display: flex;
        align-items: center;
        @apply gap-4;
    }
</style>