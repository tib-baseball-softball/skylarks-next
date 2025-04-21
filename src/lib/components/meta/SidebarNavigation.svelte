<script lang="ts">
  import {Accordion} from "@skeletonlabs/skeleton-svelte";
  import {authSettings} from "$lib/pocketbase/index.svelte";
  import type {CustomAuthModel, ExpandedClub, ExpandedTeam} from "$lib/model/ExpandedResponse";
  import {
    ChartColumnStacked,
    ChartLine,
    CircleUserRound,
    IdCard,
    LockKeyhole,
    Shield,
    SquareUserRound,
    Users,
    UsersRound
  } from "lucide-svelte";

  interface Props {
    clubs: ExpandedClub[],
    teams: ExpandedTeam[],
    sheetOpen?: boolean,
  }

  let {clubs, teams, sheetOpen = $bindable()}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  let isUserAuthenticated = $derived(!!authSettings.record);

  function sheetClose(): void {
    if (sheetOpen) {
      sheetOpen = false;
    }
  }
</script>

<nav class="list-nav py-1 px-1 lg:px-4">
  {#if isUserAuthenticated}
    <Accordion>
      <Accordion.Item>

        {#snippet lead()}
          <SquareUserRound/>
        {/snippet}

        {#snippet control()}
          <span>My Account</span>
        {/snippet}

        {#snippet panel()}
          <a href="/account" onclick={sheetClose}>
            <CircleUserRound/>
            <span>Dashboard</span>
          </a>

          <a href="/stats/{authRecord?.id}" onclick={sheetClose}>
            <ChartLine/>
            <span>Personal Stats</span>
          </a>

          <a href="/account/playerprofile" onclick={sheetClose}>
            <IdCard/>
            <span>Player Profile</span>
          </a>
        {/snippet}

      </Accordion.Item>
    </Accordion>

    <hr class="my-2"/>

    <Accordion>
      <Accordion.Item>
        {#snippet lead()}
          <UsersRound/>
        {/snippet}

        {#snippet control()}
          <span>My Clubs & Teams</span>
        {/snippet}

        {#snippet panel()}
          {#each clubs as club (club.id)}
            <a href="/account/clubs/{club.id}" onclick={sheetClose}>
              <Shield/>
              <div>{club.name} ({club.acronym})</div>
            </a>
          {/each}

          {#if teams?.length > 0}
            <hr class="mx-8">

            {#each teams as team (team.id)}
              <a
                      href="/account/team/{team.id}"
                      onclick={sheetClose}
              >
                <Users/>
                <div
                >{team.name} ({team?.expand?.club
                    ?.acronym})
                </div
                >
              </a>
            {/each}
          {/if}
        {/snippet}
      </Accordion.Item>
    </Accordion>

    <hr class="my-2"/>

    <Accordion>
      <Accordion.Item>
        {#snippet lead()}
          <LockKeyhole/>
        {/snippet}

        {#snippet control()}
          Administration
        {/snippet}

        {#snippet panel()}
          <a href="/stats/admin" onclick={sheetClose}>
            <ChartColumnStacked/>
            <span>Admin Dashboard</span>
          </a>
        {/snippet}
      </Accordion.Item>
    </Accordion>
  {/if}
</nav>

<style lang="postcss">
    a {
        display: flex;
        align-items: center;
    }

    span {
        text-wrap: nowrap;
    }
</style>
