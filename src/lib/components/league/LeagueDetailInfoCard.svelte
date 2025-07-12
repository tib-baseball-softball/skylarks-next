<script lang="ts">
  import type {LeagueGroup} from "bsm.js";
  import {Info, Tag, Users, Volleyball} from "lucide-svelte";

  interface Props {
    leagueGroup: LeagueGroup;
  }

  let {leagueGroup}: Props = $props();
</script>

<article class="card preset-tonal-surface p-4">

  <div class="container self-end">
    <Info/>
    <div>
      <p>{leagueGroup.name}</p>
      <p class="text-sm font-light">Name</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="container">
    <Tag/>
    <div>
      <p>{leagueGroup.acronym}</p>
      <p class="text-sm font-light">Acronym</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="container">
    <Volleyball/>
    <div>
      <p>{leagueGroup.league.sport}</p>
      <p class="text-sm font-light">Sport</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="container">
    <Users/>
    <dl>
      <dd>{leagueGroup.league.age_group}</dd>
      <dt class="text-sm font-light">Age Group</dt>
    </dl>
  </div>

</article>

<style>
    .container {
        display: flex;
        align-items: center;
        gap: 1.2rem;
    }
</style>