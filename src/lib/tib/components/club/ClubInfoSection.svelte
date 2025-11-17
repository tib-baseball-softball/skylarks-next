<script lang="ts">
import type { Club } from "bsm.js"
import { ClipboardList, Shield, Tag } from "lucide-svelte"

interface Props {
  club: Club
}

let { club }: Props = $props()
</script>

<section>
  <article class="flex justify-center">
    <img
            alt="Skylarks Logo"
            class="w-24 h-auto mb-1"
            src="/berlin_skylarks_logo.svg"
    />
  </article>

  <article class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl">
    <div class="container">
      <ClipboardList/>
      <span>{club.short_name}</span>
    </div>

    <hr class="my-2">

    <div class="container">
      <Tag/>
      <span>
        {club.acronym} / 0{club?.organization_id} {club.number}
      </span>
    </div>

    <hr class="my-2">

    <div class="container">
      <Shield/>
      <span>{club.main_club}</span>
    </div>
  </article>
</section>

<style>
    .container {
        display: flex;
        gap: 1rem;
    }

    article {
        margin-block: 1.5em;
    }
</style>
