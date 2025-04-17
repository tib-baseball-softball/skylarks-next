<script lang="ts">
  import type {Player} from "$lib/model/Player";
  import {getAgeFromTimestamp} from "$lib/functions/getAgeFromTimestamp.js";
  import {Calendar, Dna, Hand, Users, UsersRound} from "lucide-svelte";

  interface Props {
    player: Player;
    showTeams?: boolean;
  }

  let {player, showTeams = true}: Props = $props();
</script>

<div class="card variant-glass-surface shadow-xl p-3">
  {#if showTeams}
    <dl class="flex items-center gap-3">
      <Users aria-hidden="true"/>
      <div>
        <dd>{player.teams.map((team) => team.name)}</dd>
        <dt class="text-sm font-light">Teams</dt>
      </div>
    </dl>

    <hr class="my-2"/>
  {/if}

  <div class="flex items-center gap-3">
    <UsersRound aria-hidden="true"/>
    <dl>
      <!--do not calculate age for falsy values-->
      <dd>{player.birthday ? getAgeFromTimestamp(player.birthday) : "Not specified"}</dd>
      <dt class="text-sm font-light">Age</dt>
    </dl>
  </div>

  <hr class="my-2"/>

  <div class="flex items-center gap-3">
    <Calendar aria-hidden="true"/>
    <dl>
      <dd>{player.admission}</dd>
      <dt class="text-sm font-light">Member since</dt>
    </dl>
  </div>

  <hr class="my-2"/>

  <div class="flex items-center gap-3">
    <Dna aria-hidden="true"/>
    <dl>
      <dd>{player.positions.toLocaleString()}</dd>
      <dt class="text-sm font-light">Positions</dt>
    </dl>
  </div>

  <hr class="my-2"/>

  <div class="flex items-center gap-3">
    <Hand aria-hidden="true"/>
    <dl>
      <dd>{player.batting}</dd>
      <dt class="text-sm font-light">Bats</dt>
    </dl>
  </div>

  <hr class="my-2"/>

  <div class="flex items-center gap-3">
    <Hand aria-hidden="true"/>
    <dl>
      <dd>{player.throwing}</dd>
      <dt class="text-sm font-light">Throws</dt>
    </dl>
  </div>
</div>
