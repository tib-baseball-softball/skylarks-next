<script lang="ts">
import { preferences } from "$lib/globals.svelte.ts"
import type { LeagueGroup } from "bsm.js"

interface Props {
  leagueGroups?: LeagueGroup[]
}

let { leagueGroups = [] }: Props = $props()
</script>

<select bind:value={preferences.current.leagueGroupID} class="select min-w-52">
  <option selected value="{0}">All Leagues</option>

  {#each leagueGroups as leagueGroup}
    <option value="{leagueGroup.id}">{leagueGroup.name} ({leagueGroup.acronym})</option>
  {/each}
</select>