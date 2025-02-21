<script lang="ts">
  import {preferences} from "$lib/stores";
  import type {LeagueGroup} from "bsm.js";

  interface Props {
    leagueGroups?: LeagueGroup[];
  }

  let {leagueGroups = []}: Props = $props();
</script>

<select class="select min-w-52" bind:value={$preferences.leagueGroupID}>
  <option value="{0}" selected>Alle Ligen</option>

  {#each leagueGroups as leagueGroup}
    <option value="{leagueGroup.id}">{leagueGroup.name}</option>
  {/each}
</select>