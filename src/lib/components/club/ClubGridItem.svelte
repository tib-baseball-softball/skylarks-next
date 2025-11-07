<script lang="ts">
import type { Snippet } from "svelte"

interface Props {
  href: string
  icon: Snippet
  label: string
  description: string
}

let { href, icon, label, description }: Props = $props()
</script>

<a
        href={href}
        class="card preset-tonal dark:border dark:border-surface-500 flex items-center p-4 shadow-xl"
>
  <article>
    {@render icon()}
    <h3 class="mt-2 h4 font-semibold">{label}</h3>
    <p class="font-light">{description}</p>
  </article>
</a>
