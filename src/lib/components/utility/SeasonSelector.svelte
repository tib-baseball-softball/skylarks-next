<script lang="ts">
import { range } from "$lib/functions/range"
import { preferences } from "$lib/globals.svelte.ts"

const seasonOptions = range(2021, new Date().getFullYear())

/**
 * Favorite Team always needs to align with the selected season.
 */
function resetFavoriteTeam() {
  preferences.current.favoriteTeamID = 0
}
</script>

<select class="select" bind:value={preferences.current.selectedSeason} onchange={resetFavoriteTeam}>
  {#each seasonOptions as option}
    <option value="{option}">{option}</option>
  {/each}
</select>