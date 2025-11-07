<script lang="ts">
import type { Match } from "bsm.js"
import { Hourglass, Info, Link, Tag } from "lucide-svelte"

interface Props {
  match: Match
}

let { match }: Props = $props()
</script>

<div class="card preset-tonal-surface border border-surface-500 p-3">
  <div class="flex items-center gap-3">
    <Tag/>
    <div>
      <p>{match.league.name} ({match.league.acronym})</p>
      <p class="text-sm font-light">Liga</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <Info/>
    <div>
      <p>{match.match_id}</p>
      <p class="text-sm font-light">Spiel-ID</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <Hourglass/>
    <div>
      <p>{match.planned_innings}</p>
      <p class="text-sm font-light">geplante Innings</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <Link/>
    {#if match.scoresheet_url}
      <a class="anchor" target="_blank" href="{match.scoresheet_url}">Link zum Scoresheet</a>
    {:else}
      <span>Kein Scoresheet vorhanden.</span>
    {/if}
  </div>
</div>