<script lang="ts">
    import {ClipboardListOutline, ClipboardOutline, FileLinesOutline, InfoCircleOutline} from "flowbite-svelte-icons";
    import type {ExpandedTeam} from "$lib/model/ExpandedResponse";

    /**
   * Used on Team Detail page.
   */

  interface props {
    team: ExpandedTeam;
    link: boolean;
  }

  let {team, link = false}: props = $props();
</script>

<article class="card block variant-soft-surface p-3" class:card-hover={link}>

  <div class="flex items-center gap-3 self-end">
    <FileLinesOutline/>
    <div>
      <p>{team.name}</p>
      <p class="text-sm font-light">Name</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <InfoCircleOutline/>
    <div>
      <p>{team.bsm_league_group}</p>
      <p class="text-sm font-light">BSM-Liga (für aktuelle Saison)</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <ClipboardOutline/>
    <div>
      <p>{team?.expand?.club.name}</p>
      <p class="text-sm font-light">Club</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <ClipboardListOutline/>
    <div>
      <p>{team.age_group}</p>
      <p class="text-sm font-light">Age Group</p>
    </div>
  </div>

</article>
