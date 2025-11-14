<script lang="ts">
import type { Field } from "bsm.js"

interface Props {
  fields: Field[]
}

let { fields }: Props = $props()
</script>

<div class="grid grid-cols-1 md:grid-cols-2 gap-3 md:gap-4">
  {#each fields as field}
    <a href="/club/ballpark/{field.id}">
      <article
              class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl flex items-center gap-4">
        <img class="rounded-base" width="80" height="80" loading="lazy" src="{field.photo_url}"
             alt="a baseball or softball field: {field.name}">
        <p>
          {field.name}
        </p>
      </article>
    </a>
  {:else}
    <p>No ballparks found.</p>
  {/each}
</div>
