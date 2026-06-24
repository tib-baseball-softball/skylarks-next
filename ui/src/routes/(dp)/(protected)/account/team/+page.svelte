<script lang="ts">
  import TeamListTeaser from "$lib/dp/components/team/TeamListTeaser.svelte";
  import type {PageProps} from "./$types";
  import {authSettings} from "$lib/dp/client.svelte.ts";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const {data}: PageProps = $props();
  const userTeams = $derived(data.teams.filter((team) => authRecord.teams.includes(team.id)));
</script>

<svelte:head>
  <title>My Teams</title>
  <meta
    content="Team list page, displaying the user's teams and allowing them to manage them."
    name="description"
  />
</svelte:head>

<h1 class="h1">My Teams</h1>
<div class="teams-grid">
  {#each userTeams as team (team.id)}
    <TeamListTeaser {team} link={true}/>
  {/each}
</div>

<style>
  .teams-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 3);
    margin-block-end: calc(var(--spacing) * 6);

    @media (min-width: 48rem) {
      grid-template-columns: 1fr 1fr;
      gap: calc(var(--spacing) * 6);
    }
  }
</style>
