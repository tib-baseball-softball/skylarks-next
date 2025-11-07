<script lang="ts">
import type { Match } from "bsm.js"
import { Pencil, User } from "lucide-svelte"

interface Props {
  match: Match
}

let { match }: Props = $props()
</script>

<div class="card preset-tonal-surface border border-surface-500 p-3">
  <h4 class="h6">Umpire</h4>

  <div class="flex flex-col gap-2 mt-1">
    {#each match.umpire_assignments as umpire}
      <div class="flex items-center gap-3">
        <User/>
        <div>
          <p>{umpire.license.person.last_name}, {umpire.license.person.first_name}</p>
          <p class="text-sm font-light">{umpire.license.number}</p>
        </div>
      </div>
    {:else }
      <p class="font-light text-sm">No umpires assigned yet.</p>
    {/each}
  </div>

  <hr class="my-3">

  <h4 class="h6">Scorer</h4>

  <div class="flex flex-col gap-2 mt-1">
    {#each match.scorer_assignments as scorer}
      <div class="flex items-center gap-3">
        <Pencil/>
        <div>
          <p>{scorer.license.person.last_name}, {scorer.license.person.first_name}</p>
          <p class="text-sm font-light">{scorer.license.number}</p>
        </div>
      </div>
    {:else }
      <p class="font-light text-sm">No scorers assigned yet.</p>
    {/each}
  </div>
</div>