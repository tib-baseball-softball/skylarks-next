<script lang="ts">
  import type {ClubTeam} from "bsm.js";
  import {preferences} from "$lib/tib/globals.svelte.ts";

  interface Props {
    clubTeams?: ClubTeam[];
    onChange?: () => void;
  }

  const {
    clubTeams = [], onChange = () => {
    }
  }: Props = $props();
</script>

<label class="flex items-center gap-2">
  <span class="text-nowrap font-light">Select Team:</span>
  <select bind:value={preferences.current.favoriteTeamID} class="select my-4" id="team-picker"
          onchange={() => onChange()}>
    <option value="{0}">None</option>
    {#each clubTeams as clubTeam}
      <option value="{clubTeam.team.id}">{clubTeam.team.name}
        ({clubTeam.team.league_entries.map(entry => entry.league.acronym)})
      </option>
    {/each}
  </select>
</label>
