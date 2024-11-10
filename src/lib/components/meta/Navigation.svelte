<script lang="ts">
  import {Accordion, AccordionItem, getDrawerStore,} from "@skeletonlabs/skeleton";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import {
    ChartLineUpOutline,
    ChartMixedOutline,
    FileShieldOutline,
    LockOutline,
    ProfileCardOutline,
    ShieldOutline,
    UserCircleOutline,
    UsersOutline,
    UsersSolid,
  } from "flowbite-svelte-icons";
  import type {CustomAuthModel, ExpandedClub, ExpandedTeam} from "$lib/model/ExpandedResponse";
  import StaticNavigationLinks from "$lib/components/navigation/StaticNavigationLinks.svelte";

  interface Props {
    clubs: ExpandedClub[],
    teams: ExpandedTeam[],
  }

  let {clubs, teams}: Props = $props()

  const drawerStore = getDrawerStore();
  const model = authSettings.record as CustomAuthModel;

  let isUserAuthenticated = $derived(!!authSettings.record)

  function drawerClose(): void {
    drawerStore.close();
  }


</script>

<nav class="list-nav py-1 px-1 lg:px-4">
    <ul class="subpixel-antialiased lg:hidden">
        <StaticNavigationLinks/>
    </ul>

    {#if isUserAuthenticated}
        <Accordion regionPanel="space-y-1">
            <AccordionItem open>
                <svelte:fragment slot="lead">
                    <FileShieldOutline size="lg"/>
                </svelte:fragment>

                <svelte:fragment slot="summary">
                    <span>My Account</span>
                </svelte:fragment>

                <svelte:fragment slot="content">
                    <a href="/account" onclick={drawerClose}>
                        <UserCircleOutline size="lg"/>
                        <span>Dashboard</span>
                    </a>

                    <a href="/stats/{model?.id}" onclick={drawerClose}>
                        <ChartLineUpOutline size="lg"/>
                        <span>Personal Stats</span>
                    </a>

                    <a href="/account/playerprofile" onclick={drawerClose}>
                        <ProfileCardOutline size="lg"/>
                        <span>Player Profile</span>
                    </a>
                </svelte:fragment>
            </AccordionItem>
        </Accordion>

        <hr class="my-2"/>

        <Accordion regionPanel="space-y-1">
            <AccordionItem open>
                <svelte:fragment slot="lead">
                    <UsersSolid size="lg"/>
                </svelte:fragment>

                <svelte:fragment slot="summary">
                    <span>My Clubs & Teams</span>
                </svelte:fragment>

                <svelte:fragment slot="content">
                    {#each clubs as club}
                        <a href="/account/clubs/{club.id}" onclick={drawerClose}>
                            <ShieldOutline size="lg"/>
                            <div>{club.name} ({club.acronym})</div>
                        </a>
                    {/each}

                    {#if teams?.length > 0}
                        <hr class="mx-8">

                        {#each teams as team}
                            <a
                                    href="/account/team/{team.id}"
                                    onclick={drawerClose}
                            >
                                <UsersOutline size="lg"/>
                                <div
                                >{team.name} ({team?.expand?.club
                                    ?.acronym})
                                </div
                                >
                            </a>
                        {/each}
                    {/if}
                </svelte:fragment>
            </AccordionItem>
        </Accordion>

        <hr class="my-2"/>

        <Accordion regionPanel="space-y-1">
            <AccordionItem open>
                <svelte:fragment slot="lead">
                    <LockOutline size="lg"/>
                </svelte:fragment>

                <svelte:fragment slot="summary">Administration</svelte:fragment>

                <svelte:fragment slot="content">
                    <a href="#" onclick={drawerClose}>
                        <ChartMixedOutline size="lg"/>
                        <span>Admin Dashboard</span>
                    </a>
                </svelte:fragment>
            </AccordionItem>
        </Accordion>
    {/if}
</nav>

<style lang="postcss">
    a {
        display: flex;
        @apply items-center;
    }

    span {
        @apply text-nowrap;
    }
</style>
