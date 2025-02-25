<script lang="ts">
  import UserDataCard from "$lib/components/diamondplanner/user/UserDataCard.svelte";
  import PlayerProfileClubsSection from "$lib/components/diamondplanner/user/PlayerProfileClubsSection.svelte";
  import TeamListTeaser from "$lib/components/diamondplanner/team/TeamListTeaser.svelte";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";

  const authRecord = $derived(authSettings.record as CustomAuthModel);


  let {data} = $props();
</script>

<h1 class="h1 lg:mt-4">My Dashboard</h1>

<div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
  <UserDataCard model={authRecord}/>

  <PlayerProfileClubsSection clubs={data.clubs}/>
</div>

<h2 class="h2 mt-3">My Teams</h2>

<div class="grid grid-cols-1 md:grid-cols-2 gap-3 mb-3">
  {#each data.teams as team}
    <TeamListTeaser {team} link={true}/>
  {/each}

  {#if data.teams?.length === 0}
    <p>You are not a member of any teams yet.</p>
  {/if}
</div>
