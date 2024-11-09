<script lang="ts">
  import ClubDetailCard from "$lib/components/diamondplanner/club/ClubDetailCard.svelte";
  import UniformSetInfoCard from "$lib/components/diamondplanner/uniformset/UniformSetInfoCard.svelte";
  import {EnvelopeOutline, PlusOutline} from "flowbite-svelte-icons";
  import type {CustomAuthModel, ExpandedClub, ExpandedTeam, ExpandedUniformSet} from "$lib/model/ExpandedResponse";
  import {authModel} from "$lib/pocketbase/Auth.svelte";
  import {
    type DrawerSettings,
    getDrawerStore,
    getModalStore,
    type ModalComponent,
    type ModalSettings
  } from "@skeletonlabs/skeleton";
  import UniformSetForm from "$lib/components/forms/UniformSetForm.svelte";
  import TeamListTeaser from "$lib/components/diamondplanner/team/TeamListTeaser.svelte";

  let {data} = $props()

  const model = $authModel as CustomAuthModel;

  let club: ExpandedClub = $derived(data.club)
  let teams: ExpandedTeam[] = $derived(data.teams)
  let uniformSets: ExpandedUniformSet[] = $derived(data.uniformSets)

  const modalStore = getModalStore();
  const drawerStore = getDrawerStore();

  const teamSettings: DrawerSettings = $derived({
    id: "team-form",
    position: "right",
    width: "w-[100%] sm:w-[80%] lg:w-[70%] xl:w-[50%]",
    meta: {
      club: club,
      team: null,
    },
  });

  function triggerUniformModal() {
    const modalComponent: ModalComponent = {
      ref: UniformSetForm,
      props: {
        uniformSet: null,
        clubID: club.id,
      },
    };

    const modal: ModalSettings = {
      type: "component",
      component: modalComponent,
    };
    modalStore.trigger(modal);
  }
</script>

<svelte:head>
    <title>Details for {club.name}</title>
    <meta name="description" content="Club overview page for the {club.name} with info about teams and uniform sets."/>
</svelte:head>

<h1 class="h1">{club.name}</h1>

<section class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    <ClubDetailCard {club}/>
</section>

<section class="!mt-8">
    <header>
        <h2 class="h2 mb-3">Club Teams</h2>
    </header>

    <div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
        {#each teams as team}
            <TeamListTeaser {team} link={true}/>
        {/each}

        {#if teams.length === 0}
            <span>This club does not have any teams yet.</span>
        {/if}
    </div>

    {#if club?.admins.includes(model.id)}
        <button class="btn variant-ghost-primary" onclick={() => drawerStore.open(teamSettings)}>
            <PlusOutline/>
            <span>Create new</span>
        </button>
    {/if}
</section>

<section class="!mt-8">
    <header>
        <h2 class="h2 mb-3">Uniform Sets</h2>
    </header>

    <div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
        {#each uniformSets as uniformSet}
            <UniformSetInfoCard {uniformSet}/>
        {/each}
    </div>

    {#if club?.admins.includes(model.id)}
        <button class="btn variant-ghost-primary" onclick={triggerUniformModal}>
            <PlusOutline/>
            <span>Create new</span>
        </button>
    {/if}
</section>

{#if club?.admins.includes(model.id)}
    <hr class="my-2"/>

    <section>
        <header>
            <h2 class="h3 mb-3">Admin Section</h2>
        </header>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-2 lg:gap-3">
            <article class="card admin-card variant-ringed-surface">
                <header class="card-header">
                    <h3 class="h4 font-semibold">Club deletion</h3>
                </header>

                <section class="p-4 space-y-2">
                    <p>Deleting a club will delete all team, event and participation data. For safety reasons, a club
                        can therefore only be deleted by a superadmin.</p>
                    <p>Please contact your club's administration.</p>
                </section>

                <footer class="card-footer flex">
                    <a class="btn variant-ghost-secondary dark:variant-filled-secondary dark:border grow"
                       href="mailto:webmaster@tib-baseball.de">
                        <EnvelopeOutline/>
                        <span class="ms-2">Contact</span>
                    </a>
                </footer>
            </article>
        </div>
    </section>
{/if}

