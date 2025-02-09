<script lang="ts">
    import type {LeagueGroup} from "bsm.js";
    import {FileLinesOutline, FlagOutline, TagOutline, UsersGroupOutline} from "flowbite-svelte-icons";

    interface Props {
    leagueGroup: LeagueGroup;
  }

  let {leagueGroup}: Props = $props();
</script>

<article class="card variant-soft-surface p-3">

  <div class="flex items-center gap-3 self-end">
    <FileLinesOutline/>
    <div>
      <p>{leagueGroup.name}</p>
      <p class="text-sm font-light">Name</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <TagOutline/>
    <div>
      <p>{leagueGroup.acronym}</p>
      <p class="text-sm font-light">Akronym</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <FlagOutline/>
    <div>
      <p>{leagueGroup.league.sport}</p>
      <p class="text-sm font-light">Sportart</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <UsersGroupOutline/>
    <div>
      <p>{leagueGroup.league.age_group}</p>
      <p class="text-sm font-light">Age Group</p>
    </div>
  </div>

</article>