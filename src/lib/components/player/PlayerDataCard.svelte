<script lang="ts">
  import {
    CalendarMonthOutline,
    DnaOutline,
    DribbbleSolid,
    UserCircleOutline,
    UsersGroupOutline,
  } from "flowbite-svelte-icons";
  import type { Player } from "$lib/model/Player";
  import { getAgeFromTimestamp } from "$lib/functions/getAgeFromTimestamp.js";

  interface Props {
    player: Player;
    showTeams?: boolean;
  }

  let { player, showTeams = true }: Props = $props();
</script>

<div class="card variant-glass-surface shadow-xl p-3">
  {#if showTeams}
    <div class="flex items-center gap-3">
      <UsersGroupOutline />
      <div>
        <p>{player.teams.map((team) => team.name)}</p>
        <p class="text-sm font-light">Teams</p>
      </div>
    </div>

    <hr class="my-2" />
  {/if}

  <div class="flex items-center gap-3">
    <UserCircleOutline />
    <div>
      <p>{getAgeFromTimestamp(player.birthday)}</p>
      <p class="text-sm font-light">Alter</p>
    </div>
  </div>

  <hr class="my-2" />

  <div class="flex items-center gap-3">
    <CalendarMonthOutline />
    <div>
      <p>{player.admission}</p>
      <p class="text-sm font-light">Mitglied seit</p>
    </div>
  </div>

  <hr class="my-2" />

  <div class="flex items-center gap-3">
    <DnaOutline />
    <div>
      <p>{player.positions.toLocaleString()}</p>
      <p class="text-sm font-light">Positionen</p>
    </div>
  </div>

  <hr class="my-2" />

  <div class="flex items-center gap-3">
    <DribbbleSolid />
    <div>
      <p>{player.batting}</p>
      <p class="text-sm font-light">Schlaghand</p>
    </div>
  </div>

  <hr class="my-2" />

  <div class="flex items-center gap-3">
    <DribbbleSolid />
    <div>
      <p>{player.throwing}</p>
      <p class="text-sm font-light">Wurfhand</p>
    </div>
  </div>
</div>
