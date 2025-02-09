<script lang="ts">
    import {
        CheckOutline,
        CloseOutline,
        FileLinesOutline,
        FlagOutline,
        TagOutline,
        UsersGroupOutline
    } from "flowbite-svelte-icons";
    import type {ClubTeam} from "bsm.js";

    interface Props {
    clubTeam: ClubTeam;
  }

  let {clubTeam}: Props = $props();
</script>

<div class="card variant-ghost-surface p-3">

  <div class="flex items-center gap-3 self-end">
    <FileLinesOutline/>
    <div>
      <p>{clubTeam.team.name}</p>
      <p class="text-sm font-light">Name</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <TagOutline/>
    <div>
      <p>{clubTeam.team.short_name}</p>
      <p class="text-sm font-light">Acronym</p>
    </div>
  </div>

  {#if clubTeam.team.league_entries.at(0)}

    <hr class="my-2">

    <div class="flex items-center gap-3">
      <FlagOutline/>
      <div>
        <p>{clubTeam.team.league_entries.at(0).league.sport}</p>
        <p class="text-sm font-light">Sport</p>
      </div>
    </div>

    <hr class="my-2">

    <div class="flex items-center gap-3">
      <UsersGroupOutline/>
      <div>
        <p>{clubTeam.team.league_entries.at(0).league.age_group}</p>
        <p class="text-sm font-light">Age Group</p>
      </div>
    </div>

    <hr class="my-2">

    <div class="flex items-center gap-3">
      {#if clubTeam.team.league_entries.at(0).not_competing}
        <CheckOutline/>
      {:else }
        <CloseOutline/>
      {/if}

      <p>Antritt außer Konkurrenz (aK)</p>
    </div>

  {/if}
</div>