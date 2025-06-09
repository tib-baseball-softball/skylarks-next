<script lang="ts">
  import {preferences} from "$lib/globals.svelte.ts";
  import type {ClubTeam} from "bsm.js";

  interface Props {
    clubTeams?: ClubTeam[];
  }

  let {clubTeams = []}: Props = $props();
</script>

<select class="select my-4" bind:value={preferences.current.favoriteTeamID}>
  {#each clubTeams as clubTeam}
    <option value="{clubTeam.id}">{clubTeam.team.name} ({clubTeam.team.league_entries.map(entry => entry.league.acronym)})</option>
  {/each}
</select>