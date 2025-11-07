<script lang="ts">
import { preferences } from "$lib/globals.svelte.ts"
import type { ClubTeam } from "bsm.js"

interface Props {
  clubTeams?: ClubTeam[]
  onChange?: () => void
}

let { clubTeams = [], onChange = () => {} }: Props = $props()
</script>

<label class="flex items-center gap-2">
  <span class="text-nowrap font-light">Select Team:</span>
  <select id="team-picker" class="select my-4" bind:value={preferences.current.favoriteTeamID}
          onchange={() => onChange()}>
    <option value="{0}">None</option>
    {#each clubTeams as clubTeam}
      <option value="{clubTeam.team.id}">{clubTeam.team.name}
        ({clubTeam.team.league_entries.map(entry => entry.league.acronym)})
      </option>
    {/each}
  </select>
</label>
