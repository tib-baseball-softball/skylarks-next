<script lang="ts">
  import type {Snippet} from "svelte";

  interface Props {
    title: string
    icon?: Snippet
    children?: Snippet
  }

  let {title, icon, children}: Props = $props()
</script>

<article class="card variant-soft dark:border dark:border-surface-500 p-4 rounded-token shadow-xl mb-4">
  <div class="flex items-center gap-2 mb-2">
    {@render icon?.()}
    <h2 class="h2 text-lg font-semibold">{title}</h2>
  </div>
  {@render children?.()}
</article>
