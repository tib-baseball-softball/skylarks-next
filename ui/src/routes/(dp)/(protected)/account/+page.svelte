<script lang="ts">
  import TeamListTeaser from "$lib/dp/components/team/TeamListTeaser.svelte";
  import PlayerProfileClubsSection from "$lib/dp/components/user/PlayerProfileClubsSection.svelte";
  import UserDataCard from "$lib/dp/components/user/UserDataCard.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import AnnouncementSectionContent from "$lib/dp/components/announcements/AnnouncementSectionContent.svelte";
  import type {PageProps} from "./$types";
  import EventGrid from "$lib/dp/components/event/EventGrid.svelte";

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const {data}: PageProps = $props();
</script>

<h1 class="h1">My Dashboard</h1>

<div class="top-grid">
  <UserDataCard model={authRecord}/>

  <PlayerProfileClubsSection clubs={data.clubs}/>
</div>

<h2 class="h2">My Announcements</h2>

<AnnouncementSectionContent store={data.announcementStore}/>

<h2 class="h2">My Next Events</h2>

{#await data.events then events}
  <EventGrid events={events.items}/>
{/await}

<h2 class="h2">My Teams</h2>

<div class="teams-grid">
  {#each data.teams as team}
    <TeamListTeaser {team} link={true}/>
  {/each}

  {#if data.teams?.length === 0}
    <p>You are not a member of any teams yet.</p>
  {/if}
</div>

<style>
  h2 {
    margin-block-start: calc(var(--spacing) * 10);
    margin-block-end: calc(var(--spacing) * 5);
  }

  .top-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 3);
    margin-block-end: calc(var(--spacing) * 6);

    @media (min-width: 64rem) {
      grid-template-columns: 1fr 1fr;
    }

    @media (min-width: 80rem) {
      grid-template-columns: 1fr 1fr 1fr;
    }
  }

  .teams-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 3);
    margin-block-end: calc(var(--spacing) * 6);

    @media (min-width: 48rem) {
      grid-template-columns: 1fr 1fr;
    }
  }
</style>
