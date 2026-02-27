<script lang="ts">
  import {Plus, SquareArrowOutUpRight} from "lucide-svelte";
  import AnnouncementSectionContent from "$lib/dp/components/announcements/AnnouncementSectionContent.svelte";
  import ClubDetailCard from "$lib/dp/components/club/ClubDetailCard.svelte";
  import TeamListTeaser from "$lib/dp/components/team/TeamListTeaser.svelte";
  import UniformSetInfoCard from "$lib/dp/components/uniformset/UniformSetInfoCard.svelte";
  import AnnouncementForm from "$lib/dp/components/forms/AnnouncementForm.svelte";
  import TeamForm from "$lib/dp/components/forms/TeamForm.svelte";
  import UniformSetForm from "$lib/dp/components/forms/UniformSetForm.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {
    CustomAuthModel,
    ExpandedClub,
    ExpandedTeam,
    ExpandedUniformSet,
  } from "$lib/dp/types/ExpandedResponse.ts";
  import type {PageProps} from "./$types";

  const {data}: PageProps = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const club: ExpandedClub = $derived(data.club);
  const teams: ExpandedTeam[] = $derived(data.teams);
  const uniformSets: ExpandedUniformSet[] = $derived(data.uniformSets);
  const announcementStore = $derived(data.announcementStore);
</script>

<svelte:head>
  <title>Club Details</title>
  <meta
    content="Club overview page for the {club.name} with info about teams and uniform sets."
    name="description"
  />
</svelte:head>

<h1 class="h1 page-title">{club.name}</h1>

<section class="club-overview-section">
  <ClubDetailCard {club}/>
</section>

<section class="announcements-section section-spacing">
  <header>
    <h2 class="h2 section-title">Announcements</h2>
  </header>

  <div class="announcement-content">
    <AnnouncementSectionContent store={announcementStore}/>
  </div>

  {#if club?.admins.includes(authRecord.id)}
    <div class="form-wrapper">
      <AnnouncementForm
        announcement={null}
        team={null}
        {club}
        triggerVariant="filled-primary"
        showLabel={true}
      />
    </div>
  {/if}
</section>

<section class="teams-section section-spacing">
  <header>
    <h2 class="h2 section-title">Club Teams</h2>
  </header>

  <div class="club-page-grid">
    {#each teams as team (team.id)}
      <TeamListTeaser {team} link={true}/>
    {/each}

    {#if teams.length === 0}
      <span>This club does not have any teams yet.</span>
    {/if}
  </div>

  {#if club?.admins.includes(authRecord.id)}
    <div class="form-wrapper">
      <TeamForm
        team={null}
        {club}
        triggerVariant="filled-primary"
        showLabel={true}
      />
    </div>
  {/if}
</section>

<section class="uniform-sets-section section-spacing">
  <header>
    <h2 class="h2 section-title">Uniform Sets</h2>
  </header>

  <div class="club-page-grid">
    {#each uniformSets as uniformSet}
      <UniformSetInfoCard {uniformSet}/>
    {/each}
  </div>

  {#if club?.admins.includes(authRecord.id)}
    <div class="form-wrapper">
      <Dialog triggerClasses="btn preset-filled-primary-500">
        {#snippet triggerContent()}
          <Plus/>
          <span>Create new</span>
        {/snippet}

        {#snippet title()}
          <header>
            <h2>Create new Uniform Set</h2>
          </header>
        {/snippet}

        <UniformSetForm uniformSet={null} clubID={club.id}/>
      </Dialog>
    </div>
  {/if}
</section>

<section class="locations-section locations-section-spacing">
  <header>
    <h2 class="h2 section-title">Team Locations</h2>
  </header>

  <a
    class="btn preset-filled-primary-500"
    href="/account/clubs/{club.id}/locations"
  >
    <span>Locations Page</span>
    <SquareArrowOutUpRight size="20"/>
  </a>
</section>

{#if club?.admins.includes(authRecord.id)}
  <hr class="admin-divider"/>

  <section class="admin-module">
    <header>
      <h2 class="h3 admin-title">Admin Section</h2>
    </header>

    <div class="club-page-grid">
      <article class="card admin-card preset-outlined-surface-500">
        <header class="card-header">
          <h3 class="h4 admin-card-title">Club deletion</h3>
        </header>

        <section class="card-section">
          <p>
            Deleting a club will delete all team, event and participation data.
            For safety reasons, a club can therefore only be deleted by a
            superadmin.
          </p>
          <p>Please contact your club's administration.</p>
        </section>

        <!--        TODO: add email link-->
        <!--        <footer class="card-footer">-->
        <!--          <a-->
        <!--                  class="btn preset-tonal-secondary border border-secondary-500 dark:preset-filled-secondary-500 dark:border grow"-->
        <!--                  href=""-->
        <!--          >-->
        <!--            <Mail/>-->
        <!--            <span class="ms-2">Contact</span>-->
        <!--          </a>-->
        <!--        </footer>-->
      </article>
    </div>
  </section>
{/if}

<style>
  .club-overview-section {
    display: grid;
    grid-template-columns: 1fr;

    @media (min-width: 32rem) {
      grid-template-columns: 1fr 1fr;
    }

    @media (min-width: 64rem) {
      grid-template-columns: 1fr 1fr 1fr;
    }
  }

  .page-title {
    margin-bottom: calc(var(--spacing) * 4);
  }

  .section-spacing {
    margin-top: calc(var(--spacing) * 8);
  }

  .locations-section-spacing {
    margin-top: calc(var(--spacing) * 10);
  }

  .section-title {
    margin-bottom: calc(var(--spacing) * 3);
  }

  .announcement-content {
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 4);
  }

  .form-wrapper {
    margin-top: calc(var(--spacing) * 4);
  }

  .admin-divider {
    margin-top: calc(var(--spacing) * 10);
    margin-bottom: calc(var(--spacing) * 6);
  }

  .admin-title {
    margin-bottom: calc(var(--spacing) * 4);
  }

  .admin-card-title {
    font-weight: 600;
  }

  .card-section {
    padding: calc(var(--spacing) * 4);

    p:first-child {
      margin-block-end: calc(var(--spacing) * 2);
    }
  }

  .club-page-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 5);
    margin-block-end: calc(var(--spacing) * 4);

    @media (min-width: 64rem) {
      grid-template-columns: 1fr 1fr;
    }

    @media (min-width: 80rem) {
      grid-template-columns: 1fr 1fr 1fr;
    }
  }
</style>
