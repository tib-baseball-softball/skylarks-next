<script lang="ts">
  import type {Team} from "bsm.js";
  import {Check, File, Tag, Users, Volleyball, X} from "lucide-svelte";

  interface Props {
    clubTeam: Team;
  }

  let {clubTeam}: Props = $props();
</script>

<div class="card preset-tonal dark:border dark:border-surface-500 shadow-xl p-3">

  <div class="flex items-center gap-3 self-end">
    <File/>
    <div>
      <p>{clubTeam.name}</p>
      <p class="text-sm font-light">Name</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <Tag/>
    <div>
      <p>{clubTeam.short_name}</p>
      <p class="text-sm font-light">Acronym</p>
    </div>
  </div>

  {#if clubTeam.league_entries.at(0)}

    <hr class="my-2">

    <div class="flex items-center gap-3">
      <Volleyball/>
      <div>
        <p>{clubTeam.league_entries.at(0)?.league.sport}</p>
        <p class="text-sm font-light">Sport</p>
      </div>
    </div>

    <hr class="my-2">

    <div class="flex items-center gap-3">
      <Users/>
      <div>
        <p>{clubTeam.league_entries.at(0)?.league.age_group}</p>
        <p class="text-sm font-light">Age Group</p>
      </div>
    </div>

    <hr class="my-2">

    <div class="flex items-center gap-3 mt-3">
      <p>Antritt außer Konkurrenz (aK):</p>

      {#if clubTeam.league_entries.at(0)?.not_competing}
        <div class="text-success-500">
          <Check/>
        </div>
      {:else }
        <div class="text-error-500">
          <X/>
        </div>
      {/if}

    </div>

  {/if}
</div>
