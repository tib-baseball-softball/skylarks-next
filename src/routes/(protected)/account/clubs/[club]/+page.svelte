<script lang="ts">
  import ClubDetailCard from "$lib/components/diamondplanner/club/ClubDetailCard.svelte";
  import TeamListTeaser from "$lib/components/diamondplanner/team/TeamListTeaser.svelte";
  import UniformSetInfoCard from "$lib/components/diamondplanner/uniformset/UniformSetInfoCard.svelte";
  import UniformSetForm from "$lib/components/forms/UniformSetForm.svelte";
  import type {
    CustomAuthModel,
    ExpandedClub,
    ExpandedTeam,
    ExpandedUniformSet,
  } from "$lib/model/ExpandedResponse";
  import { authSettings, client } from "$lib/pocketbase/index.svelte";
  import { Mail, Plus, SquareArrowOutUpRight } from "lucide-svelte";
  import type { PageProps } from "./$types";
  import TeamForm from "$lib/components/forms/TeamForm.svelte";
  import Dialog from "$lib/components/utility/Dialog.svelte";
  import AnnouncementCard from "$lib/components/announcements/AnnouncementCard.svelte";

  let { data }: PageProps = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  let club: ExpandedClub = $derived(data.club);
  let teams: ExpandedTeam[] = $derived(data.teams);
  let uniformSets: ExpandedUniformSet[] = $derived(data.uniformSets);
  let announcementStore = $derived(data.announcementStore);
</script>

<svelte:head>
  <title>Details for {club.name}</title>
  <meta
    name="description"
    content="Club overview page for the {club.name} with info about teams and uniform sets."
  />
</svelte:head>

<h1 class="h1 mt-4! mb-6!">{club.name}</h1>

<section class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
  <ClubDetailCard {club} />
</section>

<section>
  <header>
    <h2 class="h2 mb-3">Announcements</h2>
  </header>

  <div class="">
    {#if $announcementStore.totalItems > 0}
      <ul class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-5 mb-3">
        {#each $announcementStore.items as announcement (announcement.id)}
          <li>
            <AnnouncementCard {announcement} />
          </li>
        {/each}
      </ul>
    {:else}
      <p>No announcements available.</p>
    {/if}
  </div>
</section>

<section class="mt-8!">
  <header>
    <h2 class="h2 mb-3">Club Teams</h2>
  </header>

  <div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-5 mb-3">
    {#each teams as team (team.id)}
      <TeamListTeaser {team} link={true} />
    {/each}

    {#if teams.length === 0}
      <span>This club does not have any teams yet.</span>
    {/if}
  </div>

  {#if club?.admins.includes(authRecord.id)}
    <TeamForm
      team={null}
      {club}
      buttonClasses="btn preset-filled-primary-500"
      showLabel={true}
    />
  {/if}
</section>

<section class="mt-8!">
  <header>
    <h2 class="h2 mb-3">Uniform Sets</h2>
  </header>

  <div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-5 mb-3">
    {#each uniformSets as uniformSet}
      <UniformSetInfoCard {uniformSet} />
    {/each}
  </div>

  {#if club?.admins.includes(authRecord.id)}
    <Dialog triggerClasses="btn preset-filled-primary-500">
      {#snippet triggerContent()}
        <Plus />
        <span>Create new</span>
      {/snippet}

      {#snippet title()}
        <header>
          <h2>Create new Uniform Set</h2>
        </header>
      {/snippet}

      <UniformSetForm uniformSet={null} clubID={club.id} />
    </Dialog>
  {/if}
</section>

<section class="mt-8!">
  <header>
    <h2 class="h2 mb-3">Team Locations</h2>
  </header>

  <a
    class="btn preset-filled-primary-500"
    href="/account/clubs/{club.id}/locations"
  >
    <span>Locations Page</span>
    <SquareArrowOutUpRight size="20" />
  </a>
</section>

{#if club?.admins.includes(authRecord.id)}
  <hr class="my-2" />

  <section>
    <header>
      <h2 class="h3 mb-3">Admin Section</h2>
    </header>

    <div
      class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-2 lg:gap-3"
    >
      <article class="card admin-card preset-outlined-surface-500">
        <header class="card-header">
          <h3 class="h4 font-semibold">Club deletion</h3>
        </header>

        <section class="p-4 space-y-2">
          <p>
            Deleting a club will delete all team, event and participation data.
            For safety reasons, a club can therefore only be deleted by a
            superadmin.
          </p>
          <p>Please contact your club's administration.</p>
        </section>

        <footer class="card-footer flex">
          <a
            class="btn preset-tonal-secondary border border-secondary-500 dark:preset-filled-secondary-500 dark:border grow"
            href="mailto:webmaster@tib-baseball.de"
          >
            <Mail />
            <span class="ms-2">Contact</span>
          </a>
        </footer>
      </article>
    </div>
  </section>
{/if}
