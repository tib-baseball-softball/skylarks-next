<script lang="ts">
  import type {Field} from "bsm.js";

  interface Props {
    field: Field,
  }

  let {field}: Props = $props();

  function convertLineBreaksToHTML(input: string): string {
    return input.replace(/\n{2,}/g, '</p><p>').replace(/\n/g, '<br>');
  }
</script>

<article class="card p-3 variant-soft dark:border dark:border-surface-500 shadow-xl">
  <p class="space-y-3">
    {@html convertLineBreaksToHTML(field.description_html)}
  </p>
</article>