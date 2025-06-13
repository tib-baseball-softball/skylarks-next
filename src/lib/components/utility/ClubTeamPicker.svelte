<script lang="ts">
  import {preferences} from "$lib/globals.svelte.ts";
  import type {ClubTeam} from "bsm.js";

  interface Props {
    clubTeams?: ClubTeam[];
    onChange?: () => void;
  }

  let {clubTeams = [], onChange = () => {}}: Props = $props();
</script>

<select class="select my-4" bind:value={preferences.current.favoriteTeamID} onchange={() => onChange()}>
  {#each clubTeams as clubTeam}
    <option value="{clubTeam.team.id}">{clubTeam.team.name} ({clubTeam.team.league_entries.map(entry => entry.league.acronym)})</option>
  {/each}
</select>