<script lang="ts">
  import {preferences} from "$lib/globals.svelte.ts";
  import type {LeagueGroup} from "bsm.js";

  interface Props {
    leagueGroups?: LeagueGroup[];
  }

  let {leagueGroups = []}: Props = $props();
</script>

<select class="select min-w-52" bind:value={preferences.current.leagueGroupID}>
  <option value="{0}" selected>Alle Ligen</option>

  {#each leagueGroups as leagueGroup}
    <option value="{leagueGroup.id}">{leagueGroup.name} ({leagueGroup.acronym})</option>
  {/each}
</select>