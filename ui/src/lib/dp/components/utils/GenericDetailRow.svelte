<script lang="ts">
  import type {Snippet} from "svelte";

  interface Props {
    icon: Snippet;
    categoryName: string;
    rowValue: string | number;
  }

  let {icon, categoryName, rowValue}: Props = $props();
</script>

<div class="container">
  {@render icon()}
  <dl class="row-content">
    <dt class="label">{categoryName}:</dt>
    <dd class="value">{rowValue}</dd>
  </dl>
</div>

<style>
  .container {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 4)
  }

  .row-content {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 4)
  }

  .label {
    font-weight: 300;
  }

  .value {
    font-weight: bold;
  }
</style>
